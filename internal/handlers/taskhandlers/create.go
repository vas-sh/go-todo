package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskBody struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

func (h *handler) create(c *gin.Context) {
	var body taskBody
	err := c.Bind(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := h.srv.Create(c.Request.Context(), body.Title, body.Description)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, task)
}
