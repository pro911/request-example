package config

import "time"

type appConf struct {
	RunMode   string `json:"run_mode" default:"debug"`
	PageSize  int    `json:"page_size" default:"10"`
	JwtSecret string `json:"jwt_secret" default:"23347$040412"`
}

type dbConf struct {
	Name            string `json:"name" default:"./demo.db"`
	TablePrefix     string `json:"table_prefix" default:"d_"`
	SetMaxIdleConns int    `json:"set_max_idle_conns" default:"10"`
	SetMaxOpenConns int    `json:"set_max_open_conns" default:"10"`
}

type httpServer struct {
	Host         string        `json:"host" default:"0.0.0.0"`
	Port         int           `json:"port" default:"30008"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

var DbConf *dbConf

func InitDbConf() {
	DbConf = &dbConf{
		Name:            "./demo.db",
		TablePrefix:     "d_",
		SetMaxIdleConns: 10,
		SetMaxOpenConns: 10,
	}
	return
}

var AppConf appConf

func InitAppConf() {
	AppConf = appConf{
		RunMode:   "debug",
		PageSize:  10,
		JwtSecret: "23347$040412",
	}
	return
}

var HttpServer *httpServer

func InitHttpServer() {
	HttpServer = &httpServer{
		Host:         "0.0.0.0",
		Port:         30008,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	return
}

var JwtSecret string

func InitJwt() {
	JwtSecret = "23347$040412"
	return
}
