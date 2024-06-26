package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"

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

	path := req.Path
	switch req.HTTPMethod {
	case "POST":
		response, err := AddAccount(ctx, req)
		response.Headers = headers
		return response, err
	case "PUT":
		response, err := UpdateAccount(ctx, req)
		response.Headers = headers
		return response, err
	case "DELETE":
		response, err := DeleteAccount(ctx, req)
		response.Headers = headers
		return response, err
	case "GET":
		var response events.APIGatewayProxyResponse
		var err error
		if path == "/accounts" {
			response, err = GetAccounts(ctx, req)
		} else {
			pathParts := strings.Split(path, "/")
			if len(pathParts) == 3 && pathParts[2] != "" {
				response, err = GetAccountByID(ctx, req)
			} else if len(pathParts) == 5 && pathParts[1] == "accounts" {
				if pathParts[2] == "graph" {
					response, err = GetAccountGraph(ctx, req)
				} else if pathParts[2] == "reports" {
					response, err = GetAccountReports(ctx, req)
				}
			} else {
				response = events.APIGatewayProxyResponse{
					StatusCode: http.StatusNotFound,
					Headers:    headers,
					Body:       "Endpoint not found",
				}
				return response, nil
			}
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
