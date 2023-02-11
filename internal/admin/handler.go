package admin

import (
	"awesomeProject/internal/handlers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	adminsURL string = "/admin"
	adminURL         = "/admin/:uuid"
)

type handler struct {
}

func NewAdminHandler() handlers.Handler {
	return &handler{}
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
