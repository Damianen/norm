package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/shiftraport"
)

type ShiftraportHandler struct {}

func (s ShiftraportHandler) HandleShiftraportShow(w http.ResponseWriter, r *http.Request) {
    shiftraport.Show().Render(r.Context(), w)
}
