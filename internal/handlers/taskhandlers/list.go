package taskhandlers

import (
	"net/http"
)

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.srv.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.homeTemplete.Execute(w, tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
