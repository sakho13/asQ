package types

// /v1/api/user/create
type CreateUserInput struct {
	FireBaseUID string
	NickName    string
	Sex         uint
	Age         uint
}

type EditUserInput struct {
	//
}
