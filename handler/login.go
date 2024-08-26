package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Damianen/norm/model"
	"github.com/Damianen/norm/view/base"
	"github.com/Damianen/norm/view/login"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	DB *sql.DB
}

func (h LoginHandler) HandleLoginShow(w http.ResponseWriter, r *http.Request) {
	management, err := getCookieData(r)

	if err != nil {
		ServeLogin(w, r)
		return
	}

	base.Show(base.Sidebar(nil, management.Name)).Render(r.Context(), w)
}

func (h LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pnl := r.Form.Get("pnl")
	password := r.Form.Get("password")

	management, err := model.GetManagementWithPnl(h.DB, pnl)

	if err != nil {
		content := login.Get(false)
		base.Show(content).Render(r.Context(), w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(management.Password), []byte(password))

	if err != nil {
		content := login.Get(true)
		base.Show(content).Render(r.Context(), w)
		return
	}

	token, err := createToken(management)

	if err != nil {
		h.HandleLogin(w, r)
		return
	}

	cookie := &http.Cookie{
		Name:     "JWT-token",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	base.Sidebar(nil, management.Name).Render(r.Context(), w)
}

func (h LoginHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:     "JWT-token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	ServeLogin(w, r)
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	content := login.Get(false)
	base.Show(content).Render(r.Context(), w)
}
