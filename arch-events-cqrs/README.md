Arquitectura en Eventos
    Se comunican atraves de una cola de eventos,
    publican sus mensajes y se suscriben a los mensajes.

CQRS
    Command Query Responsibility Segregation
    RS = Debe tener habilidades unicas para C y Q (servicios).
    Una sola forma de acceder a los datos, tanto para escrbir(C) y leer(Q) a la base de datos.
    Estos permite escalar C y Q, servicios para leer y escribir.

Como se va acoplar? 
    [C, C, C, ...] -> MB, MQ <- [Q, Q ...]

NATS servicio que te permite publicar y suscribirse a eventos.
