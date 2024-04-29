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
	path := req.Path
	switch req.HTTPMethod {
	case "POST":
		return AddAccount(ctx, req)
	case "PUT":
		return UpdateAccount(ctx, req)
	case "DELETE":
		return DeleteAccount(ctx, req)
	case "GET":
		if path == "/accounts" {
			return GetAccounts(ctx, req)
		}
		pathParts := strings.Split(path, "/")
		if len(pathParts) == 3 && pathParts[2] != "" {
			return GetAccountByID(ctx, req)
		} else if len(pathParts) == 5 && pathParts[1] == "accounts" {
			if pathParts[2] == "graph" {
				return GetAccountGraph(ctx, req)
			} else if pathParts[2] == "reports" {
				return GetAccountReports(ctx, req)
			}
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Endpoint not found",
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
