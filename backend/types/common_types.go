package types

type CommonResponseType[T interface{}] struct {
	// ResultFlg 1: success, 0: err
	ResultFlg int `json:"result_flg"`

	// Message shown when ResultFlg is 0
	Message string `json:"message"`

	// Response is main data
	Response T `json:"response"`
}
