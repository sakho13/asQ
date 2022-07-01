package types

type CommonResponseType[T interface{}] struct {
	// ResultFlg 1: success, 0: err
	ResultFlg int `json:"result_flg"`

	ErrorNo uint `json:"error_no"`

	// Message shown when ResultFlg is 0
	Message string `json:"message"`

	// Response is main data
	Response T `json:"response"`
}
