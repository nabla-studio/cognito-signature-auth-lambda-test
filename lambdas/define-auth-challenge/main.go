package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event *events.CognitoEventUserPoolsDefineAuthChallenge) (*events.CognitoEventUserPoolsDefineAuthChallenge, error) {
	if event.Request.UserNotFound {
		log.Println("User not found")
		event.Response = events.CognitoEventUserPoolsDefineAuthChallengeResponse{
			IssueTokens:        false,
			FailAuthentication: true,
			ChallengeName:      "CUSTOM_CHALLENGE",
		}
		return event, nil
	}

	if len(event.Request.Session) == 0 {
		log.Println("New session, issuing custom challenge")
		event.Response = events.CognitoEventUserPoolsDefineAuthChallengeResponse{
			IssueTokens:        false,
			FailAuthentication: false,
			ChallengeName:      "CUSTOM_CHALLENGE",
		}
		return event, nil
	}

	for _, attempt := range event.Request.Session {
		if attempt.ChallengeName != "CUSTOM_CHALLENGE" {
			log.Printf("Unexpected challenge type: %s\n", attempt.ChallengeName)
			event.Response = events.CognitoEventUserPoolsDefineAuthChallengeResponse{
				IssueTokens:        false,
				FailAuthentication: true,
				ChallengeName:      "CUSTOM_CHALLENGE",
			}
			return event, nil
		}
	}

	lastSession := event.Request.Session[len(event.Request.Session)-1]
	if !lastSession.ChallengeResult {
		log.Println("Last challenge failed")
		event.Response = events.CognitoEventUserPoolsDefineAuthChallengeResponse{
			IssueTokens:        false,
			FailAuthentication: true,
			ChallengeName:      "CUSTOM_CHALLENGE",
		}
		return event, nil
	}

	log.Println("All challenges passed, issuing tokens")
	event.Response = events.CognitoEventUserPoolsDefineAuthChallengeResponse{
		IssueTokens:        true,
		FailAuthentication: false,
	}

	return event, nil
}

func main() {
	lambda.Start(handler)
}
