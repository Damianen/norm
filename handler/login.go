package handler

import (
	"net/http"
    "time"

	"github.com/Damianen/norm/view/base"
	"github.com/Damianen/norm/view/login"
)

type LoginHandler struct {}

func (h LoginHandler) HandleLoginShow(w http.ResponseWriter, r *http.Request) {
    pnl, err := getCookieData(r)

    if err != nil {
        ServeLogin(w, r)
        return
    }

    base.Show(base.Sidebar(nil, pnl)).Render(r.Context(), w)
}

func (h LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()
    pnl := r.Form.Get("pnl")
    password := r.Form.Get("password")

    // will add db later
    if !(pnl == "pnl12345" && password == "password") {
        content := login.Get(true)
        base.Show(content).Render(r.Context(), w)
        return
    }

    token, err := createToken(pnl)

    if err != nil {
        h.HandleLogin(w, r)
        return
    }

    cookie := &http.Cookie{
        Name: "JWT-token",
        Value: token,
        HttpOnly: true,
    }

    http.SetCookie(w, cookie)

    base.Sidebar(nil, pnl).Render(r.Context(), w)
}

func (h LoginHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {

    pnl, err := getCookieData(r)

    if err != nil && pnl == "" {
        ServeLogin(w, r)
        return
    }

    cookie := &http.Cookie{
        Name: "JWT-token",
        Value: "",
        Expires: time.Unix(0, 0),
        HttpOnly: true,
    }

    http.SetCookie(w, cookie)

    ServeLogin(w, r)
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
        content := login.Get(false)
        base.Show(content).Render(r.Context(), w)
}
