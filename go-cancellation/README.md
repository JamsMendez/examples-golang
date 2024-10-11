## Patron de Concurrencia Cancellation (Cancelacion)

Es un patron que se utiliza para cancelar o abortar
operaciones en sistemas concurrentes. Proporciona
un mecanimos para notificar a las goroutines u otras
unidades de trabajo concurrente que deben de detenerse
o terminar su ejecucion.

### Ventajas

- Control de goroutines: Permiten cancelar operaciones si ya no es necesaria o
  si ha tardado demasiado.

- Recursos eficientes: Evita que se sigan ejecutando operaciones que consumen
  recursos cuando ya no son utiles.
