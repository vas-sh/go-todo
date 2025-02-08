package taskhandlers

import (
	"net/http"
)

func (h *handler) remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	title, err := h.title(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Remove(r.Context(), title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, h.homePath, http.StatusSeeOther)
}
