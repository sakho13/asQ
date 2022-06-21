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

// Initialize Firebase Parameters
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

// Decoding Firebase Token
func DecodeFirebaseToken(token string) (*auth.Token, error) {
	decodedToken, err := firebaseClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return decodedToken, nil
}

// Getting Firebase User Info
func GetUserInfo(token string) (*auth.UserRecord, error) {
	decodedToken, err := DecodeFirebaseToken(token)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	userInfo, err := firebaseClient.GetUser(context.Background(), decodedToken.UID)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return userInfo, nil
}
