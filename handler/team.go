package handler

import (
	"database/sql"
	"net/http"

	"github.com/Damianen/norm/model"
	"github.com/Damianen/norm/view/popup"
	"github.com/Damianen/norm/view/team"
)

type TeamHandler struct {
    DB *sql.DB
}

func (h TeamHandler) HandleInsertStockerShow(w http.ResponseWriter, r*http.Request) {
    team.NewStocker().Render(r.Context(), w)
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
    stockers, err := model.GetStockers(h.DB)
    if err != nil {
        team.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
        return
    }
    team.Show(stockers, nil, false).Render(r.Context(), w)
}

func (h TeamHandler) HandleTeamInsert(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    stocker := model.Stocker{}
    stocker.Name = r.Form.Get("name")

    err := model.InsertStocker(h.DB, stocker)
    if err != nil {
        h.HandleInsertStockerShow(w, r)
    }

    h.HandleTeamShow(w, r)
}
