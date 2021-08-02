.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/main.go

run:
	go run cmd/main.go

docker-build:
	docker build -t go-grafana_app .

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
