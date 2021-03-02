build:
	# Service HTTP
	env GOOS=linux go build -ldflags="-s -w" -o bin/service cmd/service/*

clean:
	rm -rf bin/*

test:
	go test ./... -v

get-aws:
	go get github.com/aws/aws-lambda-go github.com/awslabs/aws-lambda-go-api-proxy

init:
	go run boilerplate.go $(pkg)