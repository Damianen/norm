package handler

import (
	"net/http"

	"github.com/Damianen/appie-app/view/team"
)

type TeamHandler struct {}

func (t TeamHandler) HandleTeamShow(w http.ResponseWriter, r *http.Request) {
    team.Show().Render(r.Context(), w)
}
