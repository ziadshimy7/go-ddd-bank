docker_up:
	docker-compose up

start:
	go run main.go

docker_down:
	docker-compose down

tests:
	go test ./... -cover -v


migrate_up:
	migrate -path ./pkg/migrations -database "mysql://ziadshimy7:example-password@tcp(127.0.0.1:3306)/users_db?charset=utf8"  up 

migrate_down:
	migrate -path ./pkg/migrations -database "mysql://ziadshimy7:example-password@tcp(127.0.0.1:3306)/users_db?charset=utf8"  down 