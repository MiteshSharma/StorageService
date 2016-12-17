package main

import (
	"os"
	"os/signal"
	"syscall"
	"flag"
	"github.com/MiteshSharma/StorageService/utils"
	"github.com/MiteshSharma/StorageService/api"
	log "github.com/Sirupsen/logrus"
)

var configFileName string

func main() {
	parseCmdParams()

	utils.LoadConfig(configFileName)

	api.InitServer()
	api.StartServer()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
}

func parseCmdParams()  {
	flag.StringVar(&configFileName, "config", "config.json", "")
	flag.Parse()
}