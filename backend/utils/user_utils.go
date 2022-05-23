package utils

import (
	"errors"
	"strings"

	"github.com/sakho13/backend/types"
)

type UserUtil struct{}

func (u UserUtil) CheckCreateUserInput(input *types.CreateUserInput) error {

	if len(strings.TrimSpace(input.FireBaseUID)) == 0 {
		return errors.New("")
	}

	if len(strings.TrimSpace(input.NickName)) == 0 {
		return errors.New("")
	}

	return nil
}
