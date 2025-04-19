package taskhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) statuses(c *gin.Context) {
	user := userhelper.MustFromContext(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	statuses, err := h.srv.Statuses(c, user.ID, int64(id))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, statuses)
}
