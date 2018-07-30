package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/guopuke/longo/handler"
	"github.com/guopuke/longo/model"
	"github.com/guopuke/longo/pkg/errno"
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
