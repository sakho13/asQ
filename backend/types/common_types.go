package types

type CommonResponseType[T any] struct {
	result_flg int
	response   T
}
