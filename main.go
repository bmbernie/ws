package main

import (
	"github.com/bmbernie/httprouter"
	"github.com/bmbernie/ws/middleware"
	"github.com/bmbernie/ws/routes"
	"log"
	"net/http"
)

func main() {
	// Setup Routes
	router := httprouter.New()
	router.GET("/", routes.indexHandler)

	// Setup Middleware
	mh := make(middleware.MiddlewareHandler)
	mh["bpf.io"] = router

	log.Fatal(http.ListenAndServe(":80", mh))
}
