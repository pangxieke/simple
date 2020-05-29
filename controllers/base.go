package controllers

import (
	//"github.com/pangxieke/simple/routers"

	//"github.com/julienschmidt/httprouter"
	//"github.com/pangxieke/simple/routers"
	"encoding/json"
	"net/http"
)

type Controller interface {
	Init(http.ResponseWriter, *http.Request) error
	Destroy()
	Error(err error)
}

type BaseController struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (this *BaseController) Init(w http.ResponseWriter, r *http.Request) error {
	this.Request = r
	this.Response = w
	return nil
}

func (this *BaseController) Destroy() {
}

func (this *BaseController) Error(err error) {
	http.Error(this.Response, err.Error(), http.StatusInternalServerError)
}

func (this *BaseController) SuccessJson(data interface{}) (err error) {
	this.Response.WriteHeader(http.StatusOK)
	this.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	r, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = this.Response.Write(r)
	return
}
