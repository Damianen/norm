package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/team"
)

type TeamHandler struct {}

func (t TeamHandler) HandleTeamShow(w http.ResponseWriter, r *http.Request) {
    team.Show().Render(r.Context(), w)
}
