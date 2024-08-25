package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/dashboard"
)

type DashboardHandler struct {}

func (d DashboardHandler) HandleDashboardShow(w http.ResponseWriter, r *http.Request) {
    dashboard.Show().Render(r.Context(), w)
}
