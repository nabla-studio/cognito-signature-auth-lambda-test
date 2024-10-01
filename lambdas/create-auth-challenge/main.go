package main

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func handler(event *events.CognitoEventUserPoolsCreateAuthChallenge) (*events.CognitoEventUserPoolsCreateAuthChallenge, error) {
	log.Println("New login try for user with username: ", event.UserName)
	log.Println(event)

	// Generating a random nonce. We'll store it on REDIS/DynamoDB.
	// It is used to avoid reply attacks.
	nonce, err := uuid.NewRandom()
	if err != nil {
		return event, errors.New("cannot generate a nonce")
	}

	nonceString := nonce.String()

	log.Println("Generated nonce: ", nonceString)
	event.Response = events.CognitoEventUserPoolsCreateAuthChallengeResponse{
		PublicChallengeParameters: map[string]string{
			"nonce": nonceString,
		},
		PrivateChallengeParameters: map[string]string{
			"nonce": nonceString,
		},
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
