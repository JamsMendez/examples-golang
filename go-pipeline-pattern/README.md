### Patrón Pipeline

El Patrón Pipeline se destaca como un diseño robusto que facilita el procesamiento
eficiente de datos al organizar las tareas en una serie de etapas.
Cada etapa se encarga de una operación específica, y los datos fluyen a través de estas etapas,
permitiendo una ejecución concurrente y un mejor rendimiento. 
Este patrón es particularmente adecuado para Go, dado su soporte nativo para goroutines y canales,
que proporcionan los bloques de construcción necesarios para implementar pipelines.
Permite descomponer tareas de procesamiento complejas en etapas más pequeñas y manejables,
cada una ejecutándose de manera concurrente. Esto no solo mejora el rendimiento, sino que también
facilita la legibilidad y el mantenimiento del código al fomentar la separación de responsabilidades.

Caracteristicas

Etapas:         Cada Etapa se ejecuta en su propia "goroutine", lo que permite
                la ejecucion concurrente. Los "canales" se utilizan para pasar
                los datos entre las etapas.
                Deben ser reutilizables y componibles, es decir debe realizar
                una unica tarea bien definida y ser facilmente combinables con
                las otras etapas para formar pipelines complejos.
                
Errores:        Propagacion de errores. La utilizacion de un canal de errores para
                pasar errores junto con los datos. Cada etapa puede enviar errores
                a este canal con una goroutine dedicada a gestionarlos.
                Manejo de errores integrados. La integracion de errores en el canal
                de datos mediante el uso de un struct que incluya los errores asociados.

Cancelacion:    Proporcionar el mecanismo para cancelar y gestionar los tiempos de
                espera de manera eficaz.
