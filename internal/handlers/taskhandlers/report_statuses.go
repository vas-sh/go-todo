package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) reportStatuses(c *gin.Context) {
	user := userhelper.MustFromContext(c)
	statuses, err := h.srv.ReportStatuses(c, user.ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, statuses)
}
