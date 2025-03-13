package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/todo/internal/userhelper"
)

func (h *handler) remove(c *gin.Context) {
	user := userhelper.MustFromContext(c)
	err := h.srv.Remove(c.Request.Context(), user.ID)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
