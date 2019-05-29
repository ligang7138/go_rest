package main

import (
	"encoding/json"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/gomodule/redigo/redis"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go_rest/config"
	"go_rest/controllers"
	"go_rest/models"
	_ "go_rest/models"
	router2 "go_rest/router"
	"gopkg.in/tylerb/graceful.v1"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	var (
		MaxWorker = os.Getenv("MAX_WORKERS")
		MaxQueue  = os.Getenv("MAX_QUEUE")
		//MaxQueue  = os.Getenv("GOPATH")
	)
	for i:=0;i<5;i++{
		fmt.Println(MaxWorker)
		fmt.Println(MaxQueue)
		time.Sleep(1*time.Second)
	}

	//fmt.Println(time.Now().UnixNano())
	os.Exit(0)
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

	models.Init()

	server := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":8080",
			Handler: api.MakeHandler(),
		},
	}

	log.Fatal(server.ListenAndServe())


	os.Exit(3)

	// gin框架实现restful
	g := gin.Default()

	//g = router.Load(g,[]gin.HandlerFunc{}...)
	dispatcher := router2.NewDispatcher(viper.GetInt("max_workers"))
	dispatcher.Run()

	g.Run()
	//http.ListenAndServe(":8000",g)
}