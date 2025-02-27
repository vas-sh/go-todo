package taskhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Remove(c.Request.Context(), int64(id))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
