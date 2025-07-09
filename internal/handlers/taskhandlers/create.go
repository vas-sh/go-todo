package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) create(c *gin.Context) {
	var body models.Task
	err := c.ShouldBindJSON(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user := userhelper.MustFromContext(c)
	body.UserID = user.ID
	task, err := h.srv.Create(c.Request.Context(), body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	go h.event.CreateTask(user.ID)
	c.JSON(http.StatusOK, task)
}
