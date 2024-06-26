package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	"github.com/otaviomart1ns/finsys/common/config"
	commonDB "github.com/otaviomart1ns/finsys/common/db/sqlc"
	"github.com/otaviomart1ns/finsys/common/utils"
)

var (
	conn  *sql.DB
	store *commonDB.SQLStore
)

func init() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err = sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store = commonDB.NewStore(conn)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := utils.AddCorsHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	if req.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    headers,
		}, nil
	}

	switch req.HTTPMethod {
	case "POST":
		response, err := AddCategory(ctx, req)
		response.Headers = headers
		return response, err
	case "PUT":
		response, err := UpdateCategory(ctx, req)
		response.Headers = headers
		return response, err
	case "DELETE":
		response, err := DeleteCategory(ctx, req)
		response.Headers = headers
		return response, err
	case "GET":
		id := req.PathParameters["id"]
		var response events.APIGatewayProxyResponse
		var err error
		if id != "" {
			response, err = GetCategoryByID(ctx, req)
		} else {
			response, err = GetCategories(ctx, req)
		}
		response.Headers = headers
		return response, err
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Headers:    headers,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
