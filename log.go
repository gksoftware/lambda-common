package lambda_common

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
)

func LogRequestPretty(event events.APIGatewayProxyRequest) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Unable to marshall request: %s\n", err)
	} else {
		_, err = fmt.Fprintf(os.Stdout, "Received request: %s\n", string(bytes))
	}
}

func LogRequest(event events.APIGatewayProxyRequest) {
	bytes, err := json.Marshal(event)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Unable to marshall request: %s\n", err)
	} else {
		_, err = fmt.Fprintf(os.Stdout, "Received request: %s\n", string(bytes))
	}
}
