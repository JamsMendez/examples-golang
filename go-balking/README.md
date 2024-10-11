Patron de Concurrencia Balking

Es un patron utilizado para evitar que se realicen operaciones
innecesarias o redundantes en un entorno concurrente cuando
no se cumplen ciertas condiciones.
Es decir se asegura que una operacion solo se ejecute si el
sistema se encuentra en un estado especifico, de lo contrario
lo abandona y no realiza la operacion.

Estado especifico: Se basa en verificar el estado de un recurso
                   o entidad compartida.
Rechazo de accion: Si el recurso o entidad no esta en el estado
                   adecuado, la operacion se rechaza.
Sin bloqueo
innecesario:       El patron simplemente rechaza la operacion, sin
                   bloquear o esperar. Lo que puede hacer la operacion
                   mas eficiente.
