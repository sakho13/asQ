package types

type CreateUserOutput struct {
	// through sign in
	Through bool `json:"through"`
	// initialized account
	Initialized bool `json:"initialized"`
}
