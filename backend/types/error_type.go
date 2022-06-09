package types

type BasicError struct {
	Msg  string
	code uint
}

func (e BasicError) Error() error {
	// return
}
