package main

import (
	"clean-architecture/datastore"
	"clean-architecture/registry"
	"clean-architecture/router"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// config.ReadConfig
	config := struct {
		Port string
	}{}

	db := datastore.NewDB()
	// db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:", config.Port)

	if err := e.Start(":" + config.Port); err != nil {
		log.Fatalln(err)
	}
}
