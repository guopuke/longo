package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/guopuke/longo/handler"
	"github.com/guopuke/longo/pkg/errno"
	"github.com/guopuke/longo/service"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
