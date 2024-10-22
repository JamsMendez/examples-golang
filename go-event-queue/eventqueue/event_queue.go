package eventqueue

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrUnableToEnqueueEvent         = errors.New("unable to enqueue event")
	ErrEventHasAlreadyEnqueueCalled = errors.New("event has already had enqueue called")
)

// EventQueue es una estructura que se utiliza para manejar Events
// en un orden FIFO. Un EventQueue puede cerrarse, en cuyo caso todos
// los eventos que estan en la cola pero aun no han sido procesados
// seran cancelados (es decir, no se ejecutaran). Se garantiza que no
// programaran mas eventos en la EventQueue despues de que haya sido
// cerrada, sera canselado inmediatamente. Para que cualquier evento
// sea procesado por la EventQueue, debe implementar la interface
// EventHandler. Esto permite que diferentes tipos de eventos sea
// procesados por cualquier entidad que elija utilizar en una EventQueue
type EventQueue struct {
	// name es usando para diferenciar esta EventQueue de otras.
	name   string

	// queue event representa la cola de eventos.
	queue chan *Event

	// close es cerrada una vez que la EventQueue es cerrada.
	close  chan struct{}

	// drain se cierra cuando la EventQueue se detiene. Cualquier
	// Event que se Enqueue despues de que este canal este cerrado
	// sera cancelado / no sera procesado por la cola. Si un Event
	// ha sido Enqueue pero no ha sido procesado antes de que este
	// canal se cierre, tambien sera cancelado y no sera procesado.
	drain chan struct{}

	// eventQueueOnce se utiliza para asegurarse que el metodo run
	// solo se ejecute una sola vez
	eventQueueOnce sync.Once

	// closeOne se utiliza para asegurar que la EventQueue solo pueda
	// cerrarse una vez.
	closeOne       sync.Once

	eventsMu sync.RWMutex

	//  eventsClosed un canal que se cierra cuando el ciclo de eventos
	// del metodo run termine
	eventsClosed chan struct{}
}

func NewEventQueue(name string, bufferSize int) *EventQueue {
	return &EventQueue{
		name:         name,
		queue:       make(chan *Event, bufferSize),
		close:        make(chan struct{}),
		drain:        make(chan struct{}),
		eventsClosed: make(chan struct{}),
	}
}

// Enqueue inserta el evento en la EventQueue. Si la cola ha sido
// detenida, el Event no sera encolado y su canal de cancelacion
// sera cerrado, indicando que el Event no fue ejecutado. Esta 
// funcion puede bloquearse si la cola esta en su capacidad maxima
// de eventos. Si se llama a Enqueue para un unico Event varias
// veces de forma asincrona, no hay garantia de cual de ellas devolvera
// el canal que pasa los resultados de vuelta al solicitante.
// Depende del solicitante verificar si el canal devuelto es nulo, ya que
// se esperar recibir en dicho canal bloqueara indefinidamente.
// Devuelve un error si el Event ha sido encolado previamente, si el
// Evento es nulo, o si la EventQueue no esta inicializada correctamente.
func (eq *EventQueue) Enqueue(event *Event) (<-chan any, error) {
	if eq.notSafeToAccess() || event == nil {
		return nil, ErrUnableToEnqueueEvent
	}

	if !event.enqueued.CompareAndSwap(false, true) {
		// event has already had enqueue called on it
		return nil, ErrEventHasAlreadyEnqueueCalled
	}

	// Multiples llamadas a Enqueue pueden ocurrir al mismo tiempo.
	// Asegurate de que el canal de eventos no este cerrado mientras
	// estamos encolado eventos.
	eq.eventsMu.RLock()
	defer eq.eventsMu.RUnlock()

	select {
	// El Event debe ser drenado de la cola(es decir, no debe ser procesado)
	case <-eq.drain:
		// El canal eventResults cerrado indica cancelacion
		close(event.cancelled)
		close(event.eventResults)

		return event.eventResults, nil
	default:
		// El canal de Events puede cerrarse incluso si un Event ha sido
		// enviado al canal de Events, ya que los eventos se consumen del
		// canal de Events de manera asincrona. Si la EventQueue se cierra
		// antes de que este evento sea procesado, entonces sera cancelado.
		eq.queue <- event
		return event.eventResults, nil
	}
}

func (eq *EventQueue) Run() {
	if eq.notSafeToAccess() {
		return
	}

	go eq.run()
}

// Stop detiene el procesamiento de cualquier Event futuro por parte
// de la EventQueue.
// Cualquier Event que este siendo procesado actualmente por la EventQueue
// seguira ejecutandose. Todos los demas eventos que estan esperando por
// ser procesados, y todos los eventos que puedan ser encolados, no seran
// procesados por la EventQueue; seran cancelados. Si la cola ya ha sido
// detenida, esta operacion no tendra ningun efecto.
func (eq *EventQueue) Stop() {
	if eq.notSafeToAccess() {
		return
	}

	eq.closeOne.Do(func() {
		close(eq.drain)
		close(eq.close)

		eq.eventsMu.Lock()
		close(eq.queue)
		eq.eventsMu.Unlock()
	})
}

// WaitToBeDrained devuelve el canal que espera a que se detenga
// EventQueue. Esto permite que los usuarios de la cola se aseguren
// de que todos los eventos de la cola se hayan procedo o cancelado.
// Si la cola es nula, regresa inmediatamente.
func (eq *EventQueue) WaitToBeDrained() {
	if eq == nil {
		return
	}

	<-eq.close

	// Si la cola se esta ejecutando, es posible que los eventos en
	// curso aun lo sigan. Espera a que se completen para que 
	// la cola se vacie por completo. Si la cola no se esta ejecutando,
	// debemos ejecutarla a la fuerza por que nadamas lo hara para que
	// puede vaciarse.
	go eq.run()

	<-eq.eventsClosed
}

func (eq *EventQueue) run() {
	eq.eventQueueOnce.Do(func() {
		defer close(eq.eventsClosed)

		for event := range eq.queue {
			select {
			case <-eq.drain:
				close(event.cancelled)
				close(event.eventResults)
			default:
				event.Metadata.Handle(event.eventResults)
				close(event.eventResults)
			}
		}
	})
}

func (eq *EventQueue) notSafeToAccess() bool {
	return eq == nil || eq.close == nil || eq.drain == nil || eq.queue == nil
}

type Event struct {
	Metadata EventHandler

	eventResults chan any
	cancelled    chan struct{}

	enqueued atomic.Bool
}

func NewEvent(meta EventHandler) *Event {
	return &Event{
		Metadata: meta,
		// Te permite 1 escritura sin que bloquear a que alguien lo lea.
		// Buffered puedes escribir n veces (donde n es la capacidad)
		// sin esperar a que alquien lo lea, o haya espacio en la capacidad.
		eventResults: make(chan any, 1),
		cancelled:    make(chan struct{}),
	}
}

func (event *Event) WasCancelled() bool {
	select {
	case <-event.cancelled:
		return true
	default:
		return false
	}
}

// EventHandler es la interface que permite que EventQueue gestione
// eventos de forma generica. Para que EventQueue los procese, todos
// los tipos de eventos debe implementar cualquier funcion especificada
// en la interface.
type EventHandler interface {
	Handle(chan any)
}
