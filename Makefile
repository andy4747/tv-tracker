dev:
	go run main.go

sqlgen:
	sqlc generate

test:
	go test -v -cover ./...

build:
	go build -o out/tv-tracker .

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/tv-tracker-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/tv-tracker-windows-amd64 main.go

up:
	docker-compose up

down:
	docker-compose down
	docker image rm tv-tracker_api

psql:
	psql -U postgres -h localhost -d tv-tracker