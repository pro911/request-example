package main

import (
	"fmt"
	"github.com/pro911/request-example/pkg/setting"
	"github.com/pro911/request-example/routers"
	"net/http"
)

func main() {
	r := routers.InitRouter()
	
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
