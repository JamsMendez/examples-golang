## Active Object

El patrón de concurrencia Active Object desacopla la ejecución de métodos de su invocación para mejorar la concurrencia y controlar la ejecución de la aplicación.

Es decir separa la invocación de los métodos de su ejecución para asegurarse de que el invador(caller) no quede bloqueado mientras espera que la ejecución se complete. Esto permite que múltiples solicitudes de métodos sean procesadas de manera asíncrona.

### Componentes del patrón Active Object:

- Client: 
    Realiza las solicitudes.

- Active Object:
    Encapsula el objeto y delega la ejecución de los métodos a un hilo separado.

- Scheduler:
    Gestiona la cola de solicitudes y asegura que se ejecuten en el hilo adecuado.

- Result:
    Representa el valor devuelto por una solicitud cuando se complete su ejecución.

### Casos de uso
Manejo simultáneo de múltiples consultas a bases de datos.
Procesamiento de tareas en segundo plano sin bloquear la ejecución de la aplicación principal.
Aplicaciones de procesamiento de datos en tiempo real y streaming.
