PROJECT:=admin-server

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ${PROJECT} main.go
build-sqlite:
	go build -tags sqlite3 -o ${PROJECT} main.go
build-lux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${PROJECT} main.go