run:
	docker-compose build
	docker-compose up
	go run ./cmd/

delete:
	docker-compose down
	docker rmi redis:6.2-alpine postgres:latest
	docker volume rm thecarwashkz_cache thecarwashkz_pg-data