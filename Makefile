run:
	go run main.go

down:
	docker-compose -f docker-compose.yml down

up:
	docker-compose -f docker-compose.yml up --build


.PHONY: up run down