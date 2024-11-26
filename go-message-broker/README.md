### Patron de Concurrencia Publish-Subscribe (Pub/Sub) 

Permite la comunicacion de mensajes que desacopla el 
envio y recepcion de mensajes entre multiples partes.

#### Publisher
Envian mensajes a un canal comun llamado
Topic y Subject, mientras ..

#### Subscribers
Se suscriben a esos Topics para reciber
los mensajes que les interesan.

Los publisher no necesitan saben quien esta escuchando
y los subscribers no necesitan saber quien esta enviado
los mensajes.


### Message Broker

Es un componente de software que facilita la comunicacion
entre diferentes aplicaciones o servicios al enviar y recibir
mensajes de manera asincrona.

Actua como intermediario, permitiendo que los mensajes se publiquen
en un canal o cola y sea consumido por otros servicios, desacoplando
asi los componentes del sistema y mejorando la escalabiidad y
resiliencia de la aplicaciones.

#### Caracteristicas

Enrutamiento de mensajes: Permite enviar mensajes desde un productor a uno
                         o varios consumidores a traves de colas.

Desacoplamiento:          Facilita la comunicacion entres componentes sin que
                          estos necesiten estar directamente conectados o ser
                          conscientes de la ubicacion o estado del otro.

Persistencia de Mensajes: Almacena mensajes de forma persistente para asegurar
                          su entrega, incluso si el consumidor no esta disponible
                          en el momento en que el mensaje es enviado.

Escalabidad:              Permite distribuir la carga de mensajes entre multiples
                          consumidores o instancias para manejar grandes volumenes
                          de trafico.
