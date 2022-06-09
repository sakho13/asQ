package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/utils"
)

// ヘッダーでトークンを受け取る
func CheckUser(c *gin.Context) {
	headerToken := c.GetHeader("Authorization")

	userUtil := utils.UserUtil{}
	token, err := userUtil.DecodeFirebaseToken(headerToken)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(token.UID)
	// BasicResponse[]()
}
