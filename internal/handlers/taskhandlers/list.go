package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) list(c *gin.Context) {
	tasks, err := h.srv.List(c.Request.Context())
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, tasks)
}
