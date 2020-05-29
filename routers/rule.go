package routers

import (
	"github.com/pangxieke/simple/controllers"
)

func Rule() *Router {
	r := New()

	r.AddPath("/hello", Action((*controllers.HelloController).Hello))
	return r
}
