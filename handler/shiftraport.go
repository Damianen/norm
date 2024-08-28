package handler

import (
	"net/http"

	"github.com/Damianen/norm/view/shiftraport"
)

type ShiftraportHandler struct {}

func (s ShiftraportHandler) HandleShiftraportShow(w http.ResponseWriter, r *http.Request) {
    err := verifyCookie(r)
    if err != nil {
        ServeLogin(w, r)
        return
    }
    shiftraport.Show().Render(r.Context(), w)
}
