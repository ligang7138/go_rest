package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"go_rest/config"
	"go_rest/controllers"
	"go_rest/models"
	"log"
	"net/http"
	"time"
)

func main_rest()  {
	// go-json-rest框架实现restful
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/message/:id", controllers.New().GetUserName),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	models.DB.Init()

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: api.MakeHandler(),
		},
	}

	log.Fatal(server.ListenAndServe())

}
