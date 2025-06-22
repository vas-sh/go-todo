package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) reportCompletions(c *gin.Context) {
	user := userhelper.MustFromContext(c)
	completions, err := h.srv.ReportCompletions(c, user.ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, completions)
}
