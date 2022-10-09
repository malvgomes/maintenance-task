up:
	docker-compose -f ./docker-compose.yml up --build

down:
	docker-compose down -v

test:
	docker exec -ti maintenance-task_app_1  go test -cover ./...