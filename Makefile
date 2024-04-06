postgresql:
	docker run --name db_finsys -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pgpwd2024 -d -p 5432:5432 -v postgres_data:/var/lib/postgresql/data postgres:latest

createdb:
	docker exec -it db_finsys createdb --username=postgres --owner=postgres finsys

migrationup:
	migrate -path db/migration -database "postgresql://postgres:pgpwd2024@localhost:5432/finsys?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://postgres:pgpwd2024@localhost:5432/finsys?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgresql createdb migrationup migrationdown sqlc test server