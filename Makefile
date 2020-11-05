build:
	go build -o bin/raytraycer main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go

test:
	go test ./...

cover:
	go test ./... --coverprofile="cover.out"
	go tool cover -html="cover.out"
