package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/qingeekk/longo/handler"
	"github.com/qingeekk/longo/model"
	"github.com/qingeekk/longo/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
