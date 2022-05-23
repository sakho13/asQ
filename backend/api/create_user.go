package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/models"
	"github.com/sakho13/backend/types"
	"github.com/sakho13/backend/utils"
)

func CreateUser(c *gin.Context) {
	var input types.CreateUserInput
	c.Bind(&input)

	userUtil := utils.UserUtil{}

	// validate
	err := userUtil.CheckCreateUserInput(&input)
	if err != nil {
		ErrResponse(c, err)
		return
	}

	user := models.User{
		FireBaseUID: input.FireBaseUID,
		NickName:    input.NickName,
	}

	BasicResponse(c, user.NickName)
}
