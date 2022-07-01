package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sakho13/backend/models"
	"github.com/sakho13/backend/types"
)

func GetUser(c *gin.Context) {
	log.Println("> GetUser")

	var input types.GetUserInput
	c.Bind(&input)

	authJwt := c.GetHeader("Authorization")
	if authJwt == "" {
		ErrResponse(c, 1003)
		return
	}

	decodedToken, err := DecodeFirebaseToken(authJwt)
	if err != nil {
		ErrResponse(c, 1000)
		return
	}

	uuid, err := uuid.Parse(input.UserID)
	if err != nil {
		ErrResponse(c, 1002)
	}

	userInfo := models.User{}
	result :=
		DB.Where(&models.User{FireBaseUID: decodedToken.UID, ID: uuid}).Find(&userInfo)

	if result.Error != nil {
		ErrResponse(c, 1)
		return
	}
	response := types.GetUserOutput{
		UserId:    userInfo.NickName,
		TwitterID: userInfo.TwitterID,
	}
	BasicResponse(c, response)
}
