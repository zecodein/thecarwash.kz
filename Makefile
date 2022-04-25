run:
	docker-compose build
	docker image prune
	docker-compose up -d

start:
	docker-compose start

stop:
	docker-compose stop

down:
	docker-compose down

delete:
	docker-compose down
	docker rmi redis:6.2-alpine postgres:latest thecarwashkz_app:latest golang:1.17 alpine:latest
	docker volume rm thecarwashkz_cache thecarwashkz_pg-data