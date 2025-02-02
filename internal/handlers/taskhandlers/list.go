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
	h.renderTemplate(w, "html/home.html", tasks)
}
