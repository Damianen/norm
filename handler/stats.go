package handler

import (
	"net/http"

	"github.com/Damianen/appie-app/view/stats"
)

type StatsHandler struct {}

func (s StatsHandler) HandleStatsShow(w http.ResponseWriter, r *http.Request) {
    stats.Show().Render(r.Context(), w)
}
