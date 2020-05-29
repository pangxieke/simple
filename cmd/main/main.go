package main

import (
	"github.com/pangxieke/simple/log"
	"github.com/pangxieke/simple/middles"
	"github.com/pangxieke/simple/routers"
	"net/http"
)

var (
	port = "localhost:8081"
)

func main() {
	log.Info("Service start:", port)

	//log.Fatal(http.ListenAndServe(":"+port, http.Handler(routers.Rule())))

	//use middleware
	log.Fatal(http.ListenAndServe(port, middles.Handler(routers.Rule())))
}
