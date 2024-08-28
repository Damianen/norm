package handler

import (
	"database/sql"
	"net/http"
	"strconv"

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

func (h AisleHandler) HandleAisleUpdateShow(w http.ResponseWriter, r *http.Request) {
    aisleId := r.URL.Query()["id"][0]
    id, err := strconv.Atoi(aisleId)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
        return
    }

    aisleModel, err := model.GetAisle(h.DB, id)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
        return
    }

    aisle.AisleInfo(aisleModel).Render(r.Context(), w)
}

func (h AisleHandler) HandleAisle(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }

    if r.URL.Query().Has("id") {
        h.HandleAisleUpdateShow(w, r)
        return
    }

    switch r.Method {
    case "GET":
        h.HandleAisleGet(w, r)
        return
    case "POST":
        h.HandleAislePost(w, r)
        return
    case "DELETE":
        h.HandleAisleDelete(w, r)
        return
    case "UPDATE":
        h.HandleAisleUpdateShow(w, r)
        return
}
}

func (h AisleHandler) HandleAisleGet(w http.ResponseWriter, r *http.Request) {
    aisles, err := model.GetAisles(h.DB)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
    }
    aisle.Show(aisles, nil, false).Render(r.Context(), w)
}

func (h AisleHandler) HandleAislePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    aisle := model.Aisle{}
    aisle.Name = r.Form.Get("name")

    err := model.InsertAilse(h.DB, aisle)
    if err != nil {
        h.HandleAisleInsertShow(w, r)
    }

    h.HandleAisleGet(w, r)
}

func (h AisleHandler) HandleAisleDelete(w http.ResponseWriter, r *http.Request) {
    aisleId := r.URL.Query()["delete"][0]
    id, err := strconv.Atoi(aisleId)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
    }

    err = model.DeleteAisle(h.DB, id)
    if err != nil {
        aisle.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
    }
    h.HandleAisleGet(w, r)
}
