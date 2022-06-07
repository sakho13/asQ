package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sakho13/backend/types"
)

// ErrResponse
//
// エラー時のレスポンス処理
func ErrResponse(ctx *gin.Context, err string) {

	response := types.CommonResponseType[interface{}]{
		ResultFlg: 0,
		Message:   err,
		Response:  nil,
	}
	ctx.JSON(http.StatusOK, response)

}

// BasicResponse
//
// 共通レスポンス処理
func BasicResponse[T interface{}](ctx *gin.Context, res T) {

	response := types.CommonResponseType[T]{
		ResultFlg: 1,
		Message:   "",
		Response:  res,
	}
	ctx.JSON(http.StatusOK, response)

}
