package controller

import "net/http"

type Controller interface {
	ListAll(rw http.ResponseWriter, r *http.Request)
	GetByID(rw http.ResponseWriter, r *http.Request)
	Post(rw http.ResponseWriter, r *http.Request)
	Put(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
}
