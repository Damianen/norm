package handler

import (
	"net/http"

	"github.com/Damianen/appie-app/view/aisle"
)

type AisleHandler struct {}

func (a AisleHandler) HandleAisleShow(w http.ResponseWriter, r *http.Request) {
    aisle.Show().Render(r.Context(), w)
}
