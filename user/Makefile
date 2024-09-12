build:
	@go build -o user

run:
	@go run .

test:
	@go test -v ./...

testCover:
	@go test -v ./... -cover

# swagger:
# 	@"$HOME/go/bin/swag init -g ./src/routes/routes.go"

dockerize:
	@docker build -t user:latest .

dockerrun:
	@docker run --name grpcapp -p 4000:4000 user:latest