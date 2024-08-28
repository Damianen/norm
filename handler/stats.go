package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/stats"
)

type StatsHandler struct {}

func (s StatsHandler) HandleStatsShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }
    stats.Show().Render(r.Context(), w)
}
