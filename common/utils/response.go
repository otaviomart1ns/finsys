package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func AddCorsHeaders(headers map[string]string) map[string]string {
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
	headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
	return headers
}

func Response(statusCode int, data interface{}) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	headers = AddCorsHeaders(headers)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers:    headers,
	}, nil
}

func ErrorResponse(statusCode int, err error) events.APIGatewayProxyResponse {
	headers := map[string]string{"Content-Type": "application/json"}
	headers = AddCorsHeaders(headers)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       `{"error": "` + err.Error() + `"}`,
		Headers:    headers,
	}
}
