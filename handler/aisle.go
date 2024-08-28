package handler

import (
	"database/sql"
	"net/http"

	"github.com/Damianen/norm/model"
	"github.com/Damianen/norm/view/aisle"
	"github.com/Damianen/norm/view/popup"
)

type AisleHandler struct {
    DB *sql.DB
}

func (h AisleHandler) HandleAisleInsertShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }

    aisle.NewAisle().Render(r.Context(), w)
}

func (h AisleHandler) HandleAisle(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }

    switch r.Method {
    case "GET":
        h.HandleAisleShow(w, r)
        return
    case "POST":
        h.HandleAisleInsert(w, r)
        return
}
}

func (h AisleHandler) HandleAisleShow(w http.ResponseWriter, r *http.Request) {
    aisles, err := model.GetAisles(h.DB)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
    }
    aisle.Show(aisles, nil, false).Render(r.Context(), w)
}

func (h AisleHandler) HandleAisleInsert(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    aisle := model.Aisle{}
    aisle.Name = r.Form.Get("name")

    err := model.InsertAilse(h.DB, aisle)
    if err != nil {
        h.HandleAisleInsertShow(w, r)
    }

    h.HandleAisleShow(w, r)
}
