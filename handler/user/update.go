package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/guopuke/longo/handler"
	"github.com/guopuke/longo/model"
	"github.com/guopuke/longo/pkg/errno"
	"github.com/guopuke/longo/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, err, nil)
		return
	}

	u.Id = uint64(userId)

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, userId)
}
