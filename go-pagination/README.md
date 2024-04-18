### Paginación por desplazamiento

La paginación por desplazamiento aprovecha los comandos 'OFFSET' y 'LIMIT' en SQL para paginar datos.

#### Ventajas

- Permite obtener el número total de páginas.
- Permite saltar a una página específica pasando el número de página.

#### Contras

- Si un elemento en una página anterior es eliminado, los datos se desplazarán hacia adelante, causando que algunos
  resultados sean omitidos.
- Si un elemento en una página anterior es añadido, los datos se desplazarán hacia atrás, causando que algunos
  resultados sean duplicados.
- Ineficiencia en el desplazamiento: No escala bien con conjuntos de datos grandes

### Paginación por Cursor

La paginación por cursor utiliza un puntero que hace referencia a un registro específico en la base de datos.

#### Ventajas

- Como se obtiene desde un punto de referencia estable, la adición o eliminación de registros no afectará la paginación.
- Escalabilidad con conjuntos de datos grandes, porque cursor es único e indexado y la base de datos salta directamente
  al registro sin iterar a través de los datos no deseados.

#### Contras

- La paginación por cursor no permite a los clientes saltar a una página específica.
- El cursor debe provenir de una columna única y secuencial (id, timestamp, etc), de lo contrario, algunos datos
  pueden ser omitidos.
- Su clasificación limitadas. Si el requisito es ordenar basado en una columna no única, será difícil de implementar usando la paginación por cursor. La concatenación de múltiples columnas para obtener una clave única conduce a una complejidad temporal más lenta.