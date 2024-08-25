package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/planning"
)

type PlanningHandler struct {}

func (p PlanningHandler) HandlePlanningShow(w http.ResponseWriter, r *http.Request) {
    planning.Show().Render(r.Context(), w)
}
