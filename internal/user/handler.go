package user

import (
	"awesomeProject/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Check Implement Interface
var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("GET LIST"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("CreateUser"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("GetUserByUUID"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("UpdateUser"))
}

func (h *handler) PartialUpdateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("PartialUpdateUser"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Write([]byte("DeleteUser"))
}
