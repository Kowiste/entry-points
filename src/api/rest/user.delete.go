package rest

import (
	"entry/src/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a API) deleteUser(c *gin.Context) {
	user := new(user.User)
	user.ID = c.Param("id")
	err := user.Create(c.Request.Context())
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, user)
}
