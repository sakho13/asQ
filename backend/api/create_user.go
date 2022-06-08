package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/models"
	"github.com/sakho13/backend/types"
	"github.com/sakho13/backend/utilities"
)

func CreateUser(c *gin.Context) {
	var input types.CreateUserInput
	c.Bind(&input)

	userUtil := utilities.UserUtil{}

	decodedToken, err := userUtil.DecodeFirebaseToken(input.FirebaseJWT)
	if err != nil {
		log.Fatalln(err.Error())
		ErrResponse(c, "トークンデコードエラー")
		return
	}

	if decodedToken.UID != input.FirebaseUID {
		ErrResponse(c, "ID非マッチ")
		return
	}

	log.Println(decodedToken)

	user := models.User{
		FireBaseUID: decodedToken.UID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Permission:  0,
		Status:      0,
		Initialized: false,
	}

	existUsers := []models.User{}
	result := DB.Where(&models.User{FireBaseUID: decodedToken.UID}).Find(&existUsers)

	if result.RowsAffected > 0 {
		// 既に存在している
		if result.RowsAffected == 1 {
			// 1人だけである
			response := types.CreateUserOutput{
				Through:     true,
				Initialized: existUsers[0].Initialized,
			}
			BasicResponse(c, response)
			return
		} else {
			// 複数人いる
			log.Fatalln("同一UIDのユーザが複数人います")
			ErrResponse(c, "ユーザ数マッチ")
			return
		}
	} else {
		// 完全に新規
		result := DB.Create(&user)
		if result.Error != nil {
			log.Fatalln(result.Error.Error())
			ErrResponse(c, "登録処理に失敗。")
			return
		}

		response := types.CreateUserOutput{
			Through:     true,
			Initialized: false,
		}
		BasicResponse(c, response)
		return
	}
}
