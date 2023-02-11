package main

import (
	"awesomeProject/internal/tasks"
	"awesomeProject/internal/user"
	"database/sql"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect to database
	database := GetDatabase()

	// create router
	log.Println(green + "Create Router" + reset)
	router := httprouter.New()

	// create user handlers
	log.Println(green + "Create User Handlers" + reset)
	userHandler := user.NewHandler(database)
	userHandler.Register(router)

	// create tasks handlers
	log.Println(green + "Create Admin Handlers" + reset)
	adminHandler := tasks.NewAdminHandler(database)
	adminHandler.Register(router)

	log.Println(blue + "~ Start server on port: http://localhost:8080/ ~" + reset)
	start(router)

}

func GetDatabase() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// check tables or create
	var tableExists bool
	err = db.QueryRow("SELECT * FROM users").Scan(&tableExists)
	if err != nil && err != sql.ErrNoRows {
		log.Println(yellow + "The table does not exist" + reset)
		log.Println(cyan + "Creating tables for database..." + reset)

		for _, value := range Scheme {
			db.MustExec(value)
		}

		log.Printf(green + "Creating Tables!" + reset)
	} else {
		log.Println(green + "Database is ok" + reset)
	}
	log.Println(green + "Check the Database Done!" + reset)

	return db
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":"+Port)
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
