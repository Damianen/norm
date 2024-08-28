package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/dashboard"
)

type DashboardHandler struct {}

func (d DashboardHandler) HandleDashboardShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }
    dashboard.Show().Render(r.Context(), w)
}
