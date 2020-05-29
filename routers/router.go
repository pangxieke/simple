package routers

import (
	"fmt"
	"github.com/pangxieke/simple/controllers"
	"net/http"
	"reflect"
)

// 编译时检查Router 是否实现ServeHTTP 方法
var _ http.Handler = new(Router)

func New() *Router {
	router := new(Router)
	return router
}

type Handle func(w http.ResponseWriter, r *http.Request)

type Router struct {
	paths map[string]Handle
}

//register handle
func (r *Router) AddPath(path string, controller Handle) {
	if r.paths == nil {
		r.paths = map[string]Handle{}
	}
	paths := r.paths
	paths[path] = controller
	r.paths = paths
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handle, ok := r.paths[req.URL.Path]; ok {
		handle(w, req)
	} else {
		//method not register

		w.WriteHeader(http.StatusInternalServerError)
		err := fmt.Errorf("method %s is not found", req.URL.Path)
		w.Write([]byte(err.Error()))

	}
}

// 反射执行方法调用
func Action(action interface{}) Handle {
	val := reflect.ValueOf(action)
	if val.Kind() != reflect.Func {
		panic("action not func")
	}
	t := val.Type()

	//确保方法返回参数为error
	out := t.Out(0)
	if !out.Implements(interfaceOf((*error)(nil))) {
		fmt.Println("Action return not error")
	}

	t = t.In(0)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// 确保
	if !reflect.PtrTo(t).Implements(interfaceOf((*controllers.Controller)(nil))) {
		panic("not implement controller")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		v := reflect.New(t)

		//反射获取controller
		c := v.Interface().(controllers.Controller)
		//init
		err := c.Init(w, r)
		defer c.Destroy()
		if err != nil {
			c.Error(err)
			return
		}
		//反射调用方法
		ret := val.Call([]reflect.Value{v})[0].Interface()
		//fmt.Println(reflect.TypeOf(ret))
		if ret != nil {
			c.Error(ret.(error))
			return
		}
	}
}

func interfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
