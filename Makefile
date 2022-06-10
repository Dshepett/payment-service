run:
	go run ./cmd
build:
	go build -o /application ./cmd
docker:
	docker-compose up -d --build