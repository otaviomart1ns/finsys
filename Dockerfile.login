FROM public.ecr.aws/docker/library/golang:1.22 AS builder

WORKDIR /src

COPY . /src

RUN GOARCH=amd64 GOOS=linux go build -o lambda-handler ./login

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=builder /src/lambda-handler .

ENTRYPOINT ./lambda-handler