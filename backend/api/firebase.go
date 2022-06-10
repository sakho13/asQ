package api

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"
	option "google.golang.org/api/option"
)

var firebaseApp *firebase.App
var firebaseClient *auth.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase_config.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err.Error())
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	firebaseApp = app
	firebaseClient = client
}

func DecodeFirebaseToken(token string) (*auth.Token, error) {
	decodedToken, err := firebaseClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return decodedToken, nil
}
