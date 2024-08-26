package handler

import (
	"database/sql"
	"net/http"

	"github.com/Damianen/norm/view/popup"
	"github.com/Damianen/norm/view/team"
)

type TeamHandler struct {
    DB *sql.DB
}

func (h TeamHandler) HandleTeam(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        h.HandleTeamShow(w, r)
        return
    case "POST":
        h.HandleTeamInsert(w, r)
        return
}
}

func (h TeamHandler) HandleTeamShow(w http.ResponseWriter, r *http.Request) {
    team.Show(nil, popup.Err("test"), true).Render(r.Context(), w)
}

func (h TeamHandler) HandleTeamInsert(w http.ResponseWriter, r *http.Request) {

}
