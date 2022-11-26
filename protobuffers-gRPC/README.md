https://developers.google.com/protocol-buffers/docs/reference/go-generated

sudo dnf isnstall protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@lastest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/student.proto

gRPC reduce las latencias de las interaciones entre los microservicios
    Las aplicaciones era 100% monoliticas, lo que provocaba que 1 solo
    fallo derrumbara la aplicacion.
    Arquitectura orientada a servicios: aplicacion particionada segun
    sector/area, todas las partes conectadas a 1 solo servicio
    centralizaba todo.
    Microservicios: modulos de codigo con tareas bastantes especificas,
    por lo general, contaban con su propio almacenamiento.

== RPC y gRPC ==
Remote Procedure Call (RPC): Protocolo que oculta la implementacion en el backend
de la peticion que hizo un cliente, aunque el cliente sepa como hacer la peticion
y pueda invocarla como si fuese suya.
gRPC: framework creado por Google para trabajar RPC con mas eficiencia y alto
rendimiento.
    El transporte de datos funciona con HTTP2.
        Permite crear multiplexacion a la hora de enviar mensajes: mas mensajes en
        la conexion TCP de manera simultanea.
        Permite serializar datos.
    Usar los protobuffers como estructura para intercambio de datos.

== Metodos de gPRC ==
Unary: Similar a como funciona REST. Envia una peticion al servidor y el servidor
la responde.
Streaming: Permite constante envio de data en un canal
    De lado de cliente: el cliente envia muchas peticiones y el servidor responde
    una sola vez.
    De lado del servidor: el cliente realiza una sola peticion y el servidor responde
    enviando la data en partes.
    Bidireccional: el cliente y el servidor deciden ambos comunicarse por streaming
    de datos.

== Protobuffers vs JSON ==
La serializacion y deserializacion  de ambos formatos
siempre ocurre. Los protobuffers tiene mucha menos latencia
que los JSON al hacerlo.
JSON: formato de mensajes eficiente para JS
    Pares clave-valor
    Es mas facil de leer
    Es costoso en rendimiento si se quiere trabajar con otros lenguajes distintos de JS
Protobuffers: formato de mensaje agnostico a cualquier lenguaje de programcion
    Un compilador se encarga de convertir la sintaxis de protobuffer al lenguaje correspondiente.
    Esta compilacion solo ocurre en tiempo de creacion o modificacion, no en timepo de ejecucion.
    Se puede llmar archivos .proto desde otro archivos.proto

JSON se debe usar cuando la aplicacion require que la data sea mas flexible.
Protobuffers cuando la aplicacion necesita correr muy rapido, cuando los procesos
de serializacion y deserializacion debe ocurrir rapidamente.


Construir la base de datos
    docker build . -t app-grpc // app-grpc es el tag
    docker run -p 54321:5432 service-grpc-db
