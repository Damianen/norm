package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Damianen/norm/model"
	"github.com/Damianen/norm/view/management"
	"github.com/Damianen/norm/view/popup"
)

type ManagementHandler struct {
    DB *sql.DB
}

func (h ManagementHandler) HandleManagement(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }

    switch r.Method {
    case "GET":
        h.handleManagementShow(w, r)
        return
    case "POST":
        h.handleManagementInsert(w, r)
        return
}
}

func (h ManagementHandler) HandleInsertManagementShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }
    management.NewManagement().Render(r.Context(), w)
}

func (h ManagementHandler) handleManagementShow(w http.ResponseWriter, r *http.Request) {
    managementWorkers, err := model.GetManagement(h.DB)
    if err != nil {
        management.Show(nil, popup.Err(err.Error()), true).Render(r.Context(), w)
        return
    }

    management.Show(managementWorkers, nil, false).Render(r.Context(), w)
}

func (h ManagementHandler) handleManagementInsert(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    management := model.Management{}
    management.Pnl = r.Form.Get("pnl")
    management.Password = r.Form.Get("password")
    management.Name = r.Form.Get("name")
    management.Email = r.Form.Get("email")
    management.Function = r.Form.Get("function")

    err := model.InsertManagement(h.DB, management)
    if err != nil {
        fmt.Println(err)
        return
    }
    h.handleManagementShow(w, r)
}
