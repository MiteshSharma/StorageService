package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Init all routes here
func InitApi() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", get)
	InitBucket(router)
	InitFile(router)
	return router
}

func get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("{'Hello': 'World'}"))
}