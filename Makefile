docker_up:
	docker-compose up

build:
	go run main.go

docker_down:
	docker-compose down

tests:
	go test ./... -cover -v
