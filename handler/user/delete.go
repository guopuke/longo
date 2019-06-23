package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/qingeekk/longo/handler"
	"github.com/qingeekk/longo/model"
	"github.com/qingeekk/longo/pkg/errno"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
