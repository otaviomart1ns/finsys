migrationup:
	migrate -path common/db/migration -database "postgresql://postgres:pgpwd2024@db-finsys.cj6goom8aq3j.us-east-1.rds.amazonaws.com:5432/finsys" -verbose up

migrationdown:
	migrate -path common/db/migration -database "postgresql://postgres:pgpwd2024@db-finsys.cj6goom8aq3j.us-east-1.rds.amazonaws.com:5432/finsys" -verbose down

build-accounts:
	GOARCH=amd64 GOOS=linux go build -o bootstrap ./accounts
	mv bootstrap ./accounts

build-categories:
	GOARCH=amd64 GOOS=linux go build -o bootstrap ./categories
	mv bootstrap ./categories

build-login:
	GOARCH=amd64 GOOS=linux go build -o bootstrap ./login
	mv bootstrap ./login
	
build-users:
	GOARCH=amd64 GOOS=linux go build -o bootstrap ./users
	mv bootstrap ./users

run-local:
	sam local start-api
deploy-aws:
	sam build --profile otavio
	sam deploy --profile otavio

drop-stack:
	aws cloudformation delete-stack --stack-name finsys


.PHONY: migrationup migrationdown build-accounts build-categories build-login build-users run-local deploy-aws drop-stack