package taskhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user := userhelper.MustFromContext(c)
	err = h.srv.Remove(c.Request.Context(), int64(id), user.ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	go h.event.DeleteTask(user.ID)
	c.JSON(http.StatusNoContent, nil)
}
