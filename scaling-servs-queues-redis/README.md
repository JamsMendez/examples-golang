```
	public static void producer() {
		Jedis jedis = new Jedis( REDIS_HOST );

		while (true) {
		    String newElement=UUID.randomUUID().toString();
		    jedis.rpush( REDIS_LIST, newElement );

			/* limpiamos pantalla */
			System.out.print("\033[H\033[2J");
		    System.out.flush();

		    System.out.println("REDIS Productor");
		    System.out.println("===============");
		    System.out.println("");
		    System.out.println("Total de elementos: "+jedis.llen( REDIS_LIST ));
		    System.out.println("Elemento: "+jedis.lrange( REDIS_LIST , 0, -1));

		    try {
		    	Thread.sleep(1000);
		    } catch (Exception ex) { ex.printStackTrace(); }

			}

		}

	public static void consumer() {
		Jedis jedis = new Jedis( REDIS_HOST );

		/* limpiamos pantalla */
		System.out.print("\033[H\033[2J");
	    System.out.flush();

	    System.out.println("REDIS Consumidor");
	    System.out.println("================");

		while (true) {
			List<String> element=jedis.blpop( 0, REDIS_LIST );

		    System.out.print("Elemento leí­do: "+element);
		    System.out.println(" ("+jedis.llen( REDIS_LIST )+" Pendientes)");

		    try {
		    	Thread.sleep(1500);
		    } catch (Exception ex) { ex.printStackTrace(); }

			}
		}


	public static void printUsage() {
		System.out.println("");
		System.out.println("USAGE: java -jar Redis.jar <producer|consumer>");
		System.out.println("");
	}

}

```