migrationup:
	migrate -path common/db/migration -database "postgresql://postgres:pgpwd2024@finsys.cj6goom8aq3j.us-east-1.rds.amazonaws.com:5432/finsys" -verbose up

migrationdown:
	migrate -path common/db/migration -database "postgresql://postgres:pgpwd2024@finsys.cj6goom8aq3j.us-east-1.rds.amazonaws.com:5432/finsys" -verbose down

build:
	sam build --profile otavio

run-local:
	sam local start-api
	
deploy-aws: 
	sam deploy --profile otavio

drop-stack:
	aws cloudformation delete-stack --stack-name finsys


.PHONY: migrationup migrationdown run-local build-aws deploy-aws drop-stack