### Patron de Concurrencia Double Checked Locking (DCL)

Es una tecnica de programacion que se utiliza para
reducir el costo de obtener un bloque al inicializar
objectos compartidos en un entorno multithreading.

Garantiza que el bloque se adquiera solo cuando sea
estrictamente necesario permitiendo que se pueda acceder
simultaneamente a recursos de solo lectura sin neceidad
de bloquearse mutuamente.

Caracteristicas
- Primer chequeo sin bloqueo
- Segundo chequeo con bloqueo
