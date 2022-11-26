package main

import (
	"clean-architecture-three/api"
	"clean-architecture-three/core/domain/entity"
	"clean-architecture-three/repository/mysql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	repository := chooseRepository()
	service := entity.NewUserService(repository)
	handler := api.NewHanlder(service)

	http.HandleFunc("/user", handler.Get)
	http.HandleFunc("/user/add", handler.Post)

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)

		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return fmt.Sprintf(":%s", port)
}

func chooseRepository() entity.UserRepository {
	switch os.Getenv("URL_DB") {
	case "mysql":
		mysqlRepository, err := mysql.NewMySQL("connection_url")
		if err != nil {
			log.Fatalln(err)
		}

		return mysqlRepository

	case "postgresql":
		return nil
	case "api":
		return nil
	case "service":
		return nil
	}

	return nil
}
