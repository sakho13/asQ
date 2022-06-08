package utils

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	option "google.golang.org/api/option"
)

type UserUtil struct{}

func (u UserUtil) CheckFirebaseToken(token string) error {

	opt := option.WithCredentialsFile("../firebase_config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
