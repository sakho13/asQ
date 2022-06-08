package types

// /v1/api/user/create
type CreateUserInput struct {
	FirebaseJWT string `json:"firebase_jwt"`
}
