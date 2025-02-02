package taskhandlers

import (
	"net/http"
)

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	name, err := h.name(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Delete(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
