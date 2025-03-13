package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) list(c *gin.Context) {
	user := userhelper.MustFromContext(c)
	tasks, err := h.srv.List(c.Request.Context(), user.ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, tasks)
}
