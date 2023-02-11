package tasks

import (
	"awesomeProject/internal/handlers"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	adminsURL string = "/tasks"
	adminURL         = "/tasks/:uuid"
)

type handler struct {
	db *sqlx.DB
}

func NewAdminHandler(db *sqlx.DB) handlers.Handler {
	return &handler{
		db: db,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(adminsURL, h.GetAdmins)
	router.GET(adminURL, h.GetAdmin)
}

func (h *handler) GetAdmins(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "GetAdmins page")
}

func (h *handler) GetAdmin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Get Admin Page")
}
