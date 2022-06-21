package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sakho13/backend/models"
	"github.com/sakho13/backend/types"
)

func CreateUser(c *gin.Context) {
	log.Println("> CreateUser")

	var input types.CreateUserInput
	c.Bind(&input)

	decodedToken, err := DecodeFirebaseToken(input.FirebaseJWT)
	if err != nil {
		log.Fatalln(err.Error())
		ErrResponse(c, "トークンデコードエラー")
		return
	}

	userInfo, err := GetUserInfo(input.FirebaseJWT)
	if err != nil {
		log.Fatalln(err.Error())
		ErrResponse(c, "ユーザ情報取得エラー")
		return
	}

	log.Printf("UserDisplayName: %v\n", userInfo.UserInfo.DisplayName)

	user := models.User{
		FireBaseUID:   decodedToken.UID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		NickName:      "",
		SelfIntroduce: "",
		IconImg:       userInfo.PhotoURL,
		Sex:           100,
		Permission:    0,
		Status:        0,
		Initialized:   false,
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
		newUUID, err := uuid.NewRandom()
		if err != nil {
			log.Fatalln(err.Error())
			ErrResponse(c, "UUID生成失敗")
			return
		}

		user.ID = newUUID

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
