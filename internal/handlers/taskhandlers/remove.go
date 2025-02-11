package taskhandlers

import (
	"net/http"
	"strconv"
)

func (h *handler) remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Remove(r.Context(), int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, h.homePath, http.StatusSeeOther)
}
