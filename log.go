package lambda_common

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func LogRequestPretty(event events.APIGatewayProxyRequest) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		log.Printf("Unable to marshall request: %s\n", err)
	} else {
		log.Printf("Received request: %s\n", string(bytes))
	}
}

func LogRequest(event events.APIGatewayProxyRequest) {
	bytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Unable to marshall request: %s\n", err)
	} else {
		log.Printf("Received request: %s\n", string(bytes))
	}
}
