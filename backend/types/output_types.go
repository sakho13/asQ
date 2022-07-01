package types

type CreateUserOutput struct {
	UserId string `json:"user_id"`
}

type GetUserOutput struct {
	UserId    string `json:"user_id"`
	TwitterID string `json:"twitter_id"`
}
