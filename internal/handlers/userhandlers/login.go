package userhandlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	now := time.Now()
	claims := jwt.MapClaims{
		"user": user,
		"iat":  now.Unix(),
		"hbf":  now.Unix(),
		"exp":  now.Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(h.secretJWT))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"token": tokenStr,
		"type":  "JWT",
	})
}
