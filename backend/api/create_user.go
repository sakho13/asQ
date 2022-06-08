package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/types"
)

func CreateUser(c *gin.Context) {
	var input types.CreateUserInput
	c.Bind(&input)

	log.Println(input.FirebaseJWT)

	// user := models.User{
	// 	FireBaseUID: input.FireBaseUID,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// 	Permission:  0,
	// 	Status:      0,
	// 	Initialized: false,
	// }

	// existUsers := []models.User{}
	// result := DB.Where(&models.User{FireBaseUID: input.FireBaseUID}).Find(&existUsers)

	// if result.RowsAffected > 0 {
	// 	// 既に存在している
	// } else {
	// 	// 完全に新規
	// 	result := DB.Create(&user)
	// 	if result.Error != nil {
	// 		log.Fatalln(result.Error.Error())
	// 		ErrResponse(c, "登録処理に失敗しました。")
	// 		return
	// 	}
	// }

	response := types.CreateUserOutput{
		Through:     true,
		Initialized: false,
	}
	BasicResponse(c, response)
}
