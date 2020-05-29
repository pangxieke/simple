package controllers

import (
	"github.com/pangxieke/simple/log"
)

type HelloController struct {
	BaseController
}

func (this *HelloController) Hello() (err error) {
	log.Info("hello controller, method Hello")

	return this.SuccessJson("success")
}
