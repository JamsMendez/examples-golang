### Patron de Concurrencia Producer-Consumer

Es un patron que involucra a dos tipos de entidades
Productores y Consumidores.

- Productores:  Generan los datos o elementos y los colocan en 
                una cola o buffer.
- Consumidores: Toman esos datos de la cola o buffer y los procesan

En Go se puede implementar facilmente utilizando goroutines
y canales.

#### Caracteristicas

Sincronización:  Sincroniza el acceso al buffer/recurso compartido.
Desacoplamiento: El productor no depende del consumidor y viceversa.
Buffering:       Gestiona el almacenamiento temporal para los datos intermedios.
