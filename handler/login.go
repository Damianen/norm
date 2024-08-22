package handler

import (
    "net/http"
    "github.com/Damianen/appie-app/view/login"
    "github.com/Damianen/appie-app/view/base"
    "github.com/Damianen/appie-app/view/dashboard"
)

type LoginHandler struct {}

func (h LoginHandler) HandleLoginShow(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case "GET":
        content := login.Get()
        base.Show(content, "Login").Render(r.Context(), w)
    case "POST":
        content := dashboard.Get()
        base.Show(content, "Dashboard").Render(r.Context(), w)
    default:
        return;

}
}
