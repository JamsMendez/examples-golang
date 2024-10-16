#### Patron de Concurrencia Event Queue

Permite manejar y procesar eventos de manera asincrona mediante
una cola (Queue) que los ordena y almacena temporalmente hasta
que puedan ser procesados. Es util cuando se requiere evitar el 
bloque de un programa o servicio mientras se manejan eventos (notificaciones,
mensajes, tareas, etc.). Ya que las tareas se agregan en la cola
y son consumidas de forma asincrona.

### Caracteristicas
Aislamiento de eventos: Los eventos se almacena en una cola antes de ser
                        procesados, lo que permite diferir la ejecucion 
                        de una tarea hasta que sea posible.

No bloqueante:          El productor de eventos no se bloquea esperando que
                        el evento se procese.

Procesamiento
asincrono:              Los consumidores pueden manejar los eventos de la
                        cola de menera asincrona y procesarlos conforme van
                        llegando.

Control de flujo:       Permite gestionar y controlar la velocidad a la que los
                        eventos son procesados, evitando la sobrecarga.


Orden:                  Los eventos generalmente se procesan en el mismo orden en
                        que se van agregando.

Tolerante a fallas:     Los eventos se pueden reintentar o registrar cuando el
                        procesamiento falla.
