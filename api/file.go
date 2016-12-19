package api

import (
	"net/http"
	"github.com/MiteshSharma/StorageService/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/StorageService/service/data"
)

func InitFile(router *httprouter.Router) {
	router.GET("/bucket/:bucketName/file", getFiles)
	router.GET("/bucket/:bucketName/file/:fileName", getFile)
	router.POST("/bucket/:bucketName/file", uploadFile)
	router.DELETE("/bucket/:bucketName/file/:fileName", deleteFile)
}

func getFiles(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	var files []*data.File
	var err error
	if files,err = ServerObj.StorageService.GetFiles(bucketName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error fetching bucket.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(files)))
}

func getFile(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	fileName := ps.ByName("fileName")
	if fileName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No file name provided.")))
	}
	var file *data.File
	var err error
	if file,err = ServerObj.StorageService.GetFile(bucketName, fileName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error fetching file with given name.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(file)))
}

func deleteFile(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	fileName := ps.ByName("fileName")
	if fileName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No file name provided.")))
	}
	var err error
	if err = ServerObj.StorageService.RemoveFile(bucketName, fileName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error removing bucket with given name.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("{}"))
}

func uploadFile(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
		return
	}

	var file []*data.File
	var err error
	if file, err = ServerObj.StorageService.UploadFile(bucketName, r); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error uploading file in bucket.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(file)))
}