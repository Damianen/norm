package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/planning"
)

type PlanningHandler struct {}

func (p PlanningHandler) HandlePlanningShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }
    planning.Show().Render(r.Context(), w)
}
