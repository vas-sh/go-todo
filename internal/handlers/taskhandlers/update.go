package taskhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) update(c *gin.Context) {
	var body models.Task
	err := c.ShouldBindJSON(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user := userhelper.MustFromContext(c)
	userID := user.ID

	err = h.srv.Update(c, body, userID, int64(taskID))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}
