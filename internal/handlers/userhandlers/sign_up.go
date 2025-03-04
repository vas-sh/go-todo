package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

func (h *handler) signUp(c *gin.Context) {
	var body models.CreateUserBody
	err := c.BindJSON(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.srv.SignUp(c.Request.Context(), body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, user)
}
