package main

import (
	"errors"
	"internal/helpers"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event events.CognitoEventUserPoolsPreSignup) (events.CognitoEventUserPoolsPreSignup, error) {
	log.Println("New signup try for user with username: ", event.UserName)
	log.Println(event)

	// Verify the public key is provided
	pubKeyString, ok := event.Request.UserAttributes["custom:pubKSecp256k1"]
	if !ok {
		return event, errors.New("user did not provide its public key")
	}

	// Verify the signature is provided
	signatureString, ok := event.Request.ClientMetadata["custom:signature"]
	if !ok {
		return event, errors.New("user did not provide the signature")
	}

	// Verify the provided signature
	if err := helpers.VerifySignature(event.UserName, pubKeyString, signatureString, event.UserName); err != nil {
		return event, err
	}

	log.Println("Signature verified for user: ", event.UserName)
	// If the signature is valid, confirm the user
	event.Response = events.CognitoEventUserPoolsPreSignupResponse{
		AutoConfirmUser: true,
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
