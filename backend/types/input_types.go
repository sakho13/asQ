package types

// GET /v1/api/user/user
type GetUserInput struct {
	UserID string `json:"user_id"`
}

// POST /v1/api/user/create
type CreateUserInput struct {
	FirebaseJWT string `json:"firebase_jwt"`
}
