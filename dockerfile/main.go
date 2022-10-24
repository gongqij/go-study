package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"k8s-study/server"
)

var (
	help   bool
	config string
)

func init() {
	flag.StringVar(&config, "c", "config/app.json", "configuration json file like config/app.json")

}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
	}
	setupConfig()
	engine := gin.New()
	server.SetEngineHandle(engine)
	srv := server.NewServer(engine)
	srv.Run()
}
