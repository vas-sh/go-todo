package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) home(c *gin.Context) {
	tasks, err := h.srv.List(c.Request.Context())
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.homeTemplete.Execute(c.Writer, tasks)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}
