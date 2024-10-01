package main

import (
	"errors"
	"internal/helpers"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event *events.CognitoEventUserPoolsVerifyAuthChallenge) (*events.CognitoEventUserPoolsVerifyAuthChallenge, error) {
	log.Println("New verification try for user with username: ", event.UserName)
	log.Println(event)

	// Get the address of the user
	addressString := event.UserName

	// Get the public key from the private challenge parameters
	pubKeyString, ok := event.Request.UserAttributes["custom:pubKSecp256k1"]
	if !ok {
		log.Println("Cannot extract public key from challenge parameters")
		event.Response = events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
		return event, errors.New("internal server error")
	}

	// Get the nonce from the private challenge parameters
	nonceString, ok := event.Request.PrivateChallengeParameters["nonce"]
	if !ok {
		log.Println("Cannot extract nonce from challenge parameters")
		event.Response = events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
		return event, errors.New("internal server error")
	}

	// Extract the signature provided by the user
	signatureString, ok := event.Request.ChallengeAnswer.(string)
	if !ok {
		log.Println("Cannot extract signature from challenge answer")
		event.Response = events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
		return event, errors.New("cannot extract signature from challenge answer")
	}

	// Verify the provided signature
	if err := helpers.VerifySignature(addressString, pubKeyString, signatureString, nonceString); err != nil {
		log.Println(err)
		log.Println("Signature verification failed")
		event.Response = events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
			AnswerCorrect: false,
		}
		return event, err
	}

	// If signature verification is successful, return answer correct
	log.Println("Signature verification successful")
	// If everything is good, return answer correct
	event.Response = events.CognitoEventUserPoolsVerifyAuthChallengeResponse{
		AnswerCorrect: true,
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
