package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sakho13/backend/maps"
	"github.com/sakho13/backend/types"
)

// ErrResponse
//
// エラー時のレスポンス処理
func ErrResponse(ctx *gin.Context, error_no uint) {

	errorMsg := maps.ErrorMsgMap[error_no]

	response := types.CommonResponseType[interface{}]{
		ResultFlg: 0,
		Response:  nil,
	}

	if errorMsg != "" {
		response.ErrorNo = error_no
		response.Message = errorMsg
	} else {
		response.ErrorNo = 9999
		response.Message = "不明なエラー"
	}

	log.Fatalln(response.Message)

	ctx.JSON(http.StatusOK, response)

}

// BasicResponse
//
// 共通レスポンス処理
func BasicResponse[T interface{}](ctx *gin.Context, res T) {

	response := types.CommonResponseType[T]{
		ResultFlg: 1,
		ErrorNo:   0,
		Message:   "",
		Response:  res,
	}
	ctx.JSON(http.StatusOK, response)

}
