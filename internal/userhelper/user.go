package userhelper

import (
	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/models"
)

func GetFromContext(c *gin.Context) (*models.User, error) {
	val, ok := c.Get(models.UserContextKey)
	if !ok {
		return nil, models.ErrInvalidUser
	}
	user, ok := val.(*models.User)
	if !ok {
		return nil, models.ErrInvalidUser
	}
	return user, nil
}

func MustFromContext(c *gin.Context) *models.User {
	user, err := GetFromContext(c)
	if err != nil {
		panic(err)
	}
	return user
}
