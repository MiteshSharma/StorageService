package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/MiteshSharma/StorageService/utils"
	"github.com/MiteshSharma/StorageService/middleware"
	"github.com/MiteshSharma/StorageService/service"
)

type Server struct  {
	Router *httprouter.Router
	StorageService *service.StorageService
}

var ServerObj *Server

func InitServer()  {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
	ServerObj.StorageService = service.NewStorageService(utils.ConfigParam.StorageConfig.ProviderType)
}

func StartServer()  {
	go func() {
		negroni := negroni.Classic()
		negroni.Use(middleware.NewRequest())
		negroni.UseHandler(ServerObj.Router)
		negroni.Run(utils.ConfigParam.ServerConfig.Port)
	}()
}

func StopServer()  {
	// Cleanup on server close
}