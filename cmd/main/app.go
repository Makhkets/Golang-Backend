package main

import (
	"awesomeProject/cmd/main/config"
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
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	// connect to database
	database := GetDatabase()

	// validator
	config.Validate = validator.New()
	// create router
	log.Println(config.Green + "Create Router" + config.Reset)
	router := httprouter.New()

	// create user handlers
	log.Println(config.Green + "Create User Handlers" + config.Reset)
	userHandler := user.NewHandler(database)
	userHandler.Register(router)

	// create tasks handlers
	log.Println(config.Green + "Create Tasks Handlers" + config.Reset)
	adminHandler := tasks.NewAdminHandler(database, config.Validate)
	adminHandler.Register(router)

	log.Println(config.Blue + "~ Start server on port: http://localhost:8000/ ~" + config.Reset)
	start(router)

}

func GetDatabase() *sqlx.DB {
	db, err1 := sqlx.Open("sqlite3", "file:database.db?cache=shared&mode=rwc")
	if err1 != nil {
		log.Fatalln(err1)
	}
	err1 = db.Ping()
	if err1 != nil {
		log.Fatalln(err1)
	}

	var err error
	var tableExists bool
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='tasks'").Scan(&tableExists)
	if err == sql.ErrNoRows {
		log.Println(config.Yellow + "The table does not exist" + config.Reset)
		log.Println(config.Cyan + "Creating tables for database..." + config.Reset)

		for _, value := range config.Scheme {
			_, err := db.Exec(value)
			if err != nil {
				return db
			}
		}

		log.Printf(config.Green + "Creating Tables!" + config.Reset)
	} else if err != nil {
		return db
	} else {
		log.Println(config.Green + "Table exists" + config.Reset)
	}
	log.Println(config.Green + "Checking tables in database Done!" + config.Reset)

	return db
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":"+config.Port)
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
