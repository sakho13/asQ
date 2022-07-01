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
		ErrResponse(c, 1000)
		return
	}

	userInfo, err := GetUserInfo(input.FirebaseJWT)
	if err != nil {
		log.Fatalln(err.Error())
		ErrResponse(c, 1)
		return
	}

	log.Printf("UserDisplayName: %v\n", userInfo.UserInfo.DisplayName)

	user := models.User{
		FireBaseUID: decodedToken.UID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		NickName:    userInfo.DisplayName,
		IconImg:     userInfo.PhotoURL,
		TwitterID:   "",
	}

	existUsers := []models.User{}
	result := DB.Where(&models.User{FireBaseUID: decodedToken.UID}).Find(&existUsers)

	if result.RowsAffected > 0 {
		// 既に存在している
		if result.RowsAffected == 1 {
			// 1人だけである
			user := existUsers[0]
			response := types.CreateUserOutput{
				UserId: user.ID.String(),
			}
			BasicResponse(c, response)
			return
		} else {
			// 複数人いる
			ErrResponse(c, 2000)
			return
		}
	} else {
		// 完全に新規
		newUUID, err := uuid.NewRandom()
		if err != nil {
			log.Fatalln(err.Error())
			ErrResponse(c, 1001)
			return
		}

		user.ID = newUUID

		if result := DB.Create(&user); result.Error != nil {
			log.Fatalln(result.Error.Error())
			ErrResponse(c, 2000)
			return
		}

		response := types.CreateUserOutput{
			UserId: user.ID.String(),
		}
		BasicResponse(c, response)
		return
	}
}
