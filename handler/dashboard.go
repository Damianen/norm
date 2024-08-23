package handler

import (
	"net/http"

	"github.com/Damianen/appie-app/view/dashboard"
)

type DashboardHandler struct {}

func (d DashboardHandler) HandleDashboardShow(w http.ResponseWriter, r *http.Request) {
    dashboard.Show().Render(r.Context(), w)
}
