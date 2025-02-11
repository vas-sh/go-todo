package taskhandlers

import (
	"net/http"
)

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	body, err := h.body(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Create(r.Context(), body.title, body.description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, h.homePath, http.StatusSeeOther)
}

func (h *handler) createForm(w http.ResponseWriter, _ *http.Request) {
	err := h.createFormTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
