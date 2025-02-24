package taskhandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) remove(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Remove(c.Request.Context(), int64(id))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(c.Writer, c.Request, h.homePath, http.StatusSeeOther)
}

func (h *handler) apiRemove(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.srv.Remove(c.Request.Context(), int64(id))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, id)
}
