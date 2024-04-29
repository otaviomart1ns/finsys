package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func Response(statusCode int, data interface{}) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func ErrorResponse(statusCode int, err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       `{"error": "` + err.Error() + `"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}
}
