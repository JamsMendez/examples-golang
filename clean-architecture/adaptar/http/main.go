package main

import (
	"clean-architecture/adaptar/postgres"
	"clean-architecture/di"
	"context"
	"net/http"
)

func init() {
	// reading cofiguration file

	/*
		 	viper.SetConfigFile(`config.json`)
			err := viper.ReadInConfig()
			if err != nil {
				panic(err)
			}
	*/
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)

	postgres.RunMigration()
	productService := di.ConfigProductDI(conn)

	http.HandleFunc("/product", productService.Create)
	http.HandleFunc("/products", productService.Fetch)

	err := http.ListenAndServe(":PORT", nil)
	if err != nil {
		return
	}
}
