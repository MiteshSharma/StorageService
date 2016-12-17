package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/MiteshSharma/StorageService/service/data"
	"github.com/MiteshSharma/StorageService/utils"
)

func InitBucket(router *httprouter.Router) {
	router.GET("/bucket", getBuckets)
	router.GET("/bucket/:bucketName", getBucket)
	router.POST("/bucket/:bucketName", createBucket)
	router.DELETE("/bucket/:bucketName", deleteBucket)
}

func getBuckets(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var buckets []*data.Bucket
	var err error
	if buckets,err = ServerObj.StorageService.GetBuckets(); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error fetching bucket.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(buckets)))
}

func getBucket(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	var bucket *data.Bucket
	var err error
	if bucket,err = ServerObj.StorageService.GetBucket(bucketName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error fetching bucket with given name.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(bucket)))
}

func createBucket(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	var bucket *data.Bucket
	var err error
	if bucket,err = ServerObj.StorageService.CreateBucket(bucketName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error creating bucket with given name.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(bucket)))
}

func deleteBucket(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bucketName := ps.ByName("bucketName")
	if bucketName == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("No bucket name provided.")))
	}
	var err error
	if err = ServerObj.StorageService.DestroyBucket(bucketName); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson("Error removing bucket with given name.")))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("{}"))
}