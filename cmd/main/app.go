package main

import (
	"awesomeProject/internal/admin"
	"awesomeProject/internal/user"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// create router
	log.Println("Create Router")
	router := httprouter.New()

	// create user handlers
	log.Println("Create User Handlers")
	userHandler := user.NewHandler()
	userHandler.Register(router)

	// create admin handlers
	log.Println("Create Admin Handlers")
	adminHandler := admin.NewAdminHandler()
	adminHandler.Register(router)

	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	log.Println("~ Start server on port: http://localhost:8080/ ~")
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
