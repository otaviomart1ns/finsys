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
	switch req.HTTPMethod {
	case "POST":
		return Login(ctx, req)
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
