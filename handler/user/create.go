package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/qingeekk/longo/handler"
	"github.com/qingeekk/longo/model"
	"github.com/qingeekk/longo/pkg/errno"
	"github.com/qingeekk/longo/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		log.Infof("Database error: %s", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is : [%s], password is [%s]", r.Username, r.Password)

	rsp := CreateResponse{Username: r.Username}

	// Show the user information
	SendResponse(c, nil, rsp)
}
