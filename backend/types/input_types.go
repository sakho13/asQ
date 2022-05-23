package types

// /v1/api/user/create
type CreateUserInput struct {
	FireBaseUID string `json:"firebase_uid"`
	NickName    string `json:"nick_name"`
	Sex         uint   `json:"sex"`
	Age         uint   `json:"age"`
}
