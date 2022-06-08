package utilities

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"
	option "google.golang.org/api/option"
)

type UserUtil struct{}

func (u UserUtil) DecodeFirebaseToken(token string) (*auth.Token, error) {

	opt := option.WithCredentialsFile("../firebase_config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	decodedToken, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return decodedToken, nil
}
