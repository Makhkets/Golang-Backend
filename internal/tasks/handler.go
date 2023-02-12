package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"awesomeProject/cmd/main/config"
	"awesomeProject/internal/handlers"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
)

const (
	tasksURL string = "/tasks"
	taskURL  string = "/tasks/:uuid"
)

type handler struct {
	db       *sqlx.DB
	validate *validator.Validate
}

func NewAdminHandler(db *sqlx.DB, validate *validator.Validate) handlers.Handler {
	return &handler{
		db:       db,
		validate: validate,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(tasksURL, h.GetTasks)    // [~] Get Tasks
	router.GET(taskURL, h.GetTask)      // Get Task
	router.POST(tasksURL, h.CreateTask) // [+] Create Task
	router.PUT(taskURL, h.GetTask)      // Update Task
	router.DELETE(taskURL, h.GetTask)   // Delete Task
}

func (h *handler) GetTasks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tasks, err := GetTasks(h.db)
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(config.Red, err, config.Reset)
		return
	}

	res, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(config.Red, err, config.Reset)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("uuid")
	uuid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid UUID format. UUID must be an integer.", http.StatusBadRequest)
		return
	}

	task, err := GetTask(uuid, h.db)
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(config.Red, err, config.Reset)
		return
	}

	log.Println(config.Blue, "Success found task with your id", config.Reset)

	res, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Fprint(w, err)
		log.Println(config.Red, err, config.Reset)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	task := Task{
		Title:       r.PostFormValue("title"),
		Description: r.PostFormValue("description"),
		Author:      r.PostFormValue("author"),
	}

	log.Printf(config.Green + fmt.Sprintf("%v", task) + config.Reset)
	log.Println(config.Cyan, r.PostForm, config.Reset)

	if err := h.validate.Struct(task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}

	err := AddTaskDatabase([]Task{task}, h.db)
	if err != nil {
		log.Println(config.Red, "Error! -> Not added tasks to db", "\n", err, config.Reset)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) UpdateTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Update Task
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Delete Task
}
