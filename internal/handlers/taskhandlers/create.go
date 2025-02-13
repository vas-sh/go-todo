package taskhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskBody struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func (h *handler) create(c *gin.Context) {
	var body taskBody
	err := c.Bind(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Create(c.Request.Context(), body.Title, body.Description)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(c.Writer, c.Request, h.homePath, http.StatusSeeOther)
}

func (h *handler) addTask(c *gin.Context) {
	err := h.createFormTemplate.Execute(c.Writer, nil)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}
