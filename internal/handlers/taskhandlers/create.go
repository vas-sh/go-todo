package taskhandlers

import (
	"net/http"
)

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	name, err := h.name(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Create(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *handler) createForm(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "html/add-task.html", nil)
}
