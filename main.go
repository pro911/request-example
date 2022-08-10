package main

import (
	"fmt"
	"net/http"
	"request-example/config"
	"request-example/models"
	"request-example/routers"
)

func init() {
	config.InitDbConf()
	config.InitAppConf()
	config.InitHttpServer()
	config.InitJwt()
	models.InitDb()
}

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port),
		Handler:        r,
		ReadTimeout:    config.HttpServer.ReadTimeout,
		WriteTimeout:   config.HttpServer.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

}
