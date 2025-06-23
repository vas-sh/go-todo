package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

func (h *handler) login(c *gin.Context) {
	var body models.LoginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.srv.Login(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	if !user.Activated {
		http.Error(c.Writer, "not activated", http.StatusForbidden)
		return
	}
	h.returnJWTToken(c, user)
}

func (h *handler) returnJWTToken(c *gin.Context, user models.User) {
	tokenStr, err := h.userFetcher.CreateJWT(user)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"token": tokenStr,
		"type":  "JWT",
		"user":  user,
	})
}
