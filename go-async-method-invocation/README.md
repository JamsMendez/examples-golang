Patrón de concurrencia Async Method Invocation

Es un patrón donde una tarea o método se ejecuta de manera asíncrona,
permitiendo que el flujo principal de un programa continúe sin esperar
a que el método finalice.
Este patrón permite que las tareas se ejecuten en paralelo y se puedan
recuperar sus resultados más tarde cuando sea necesario.

En Go, este patrón se puede implementar mediante goroutines y canales
para manejar las tareas asincrónicas y recuperar los resultados
cuando estén listos.


Client:        El invocador de la función asincrónica.
AsyncFunction: La función que se ejecutará de forma asincrónica.
Future:        El resultado de la operación y que se devuelve al Client.

