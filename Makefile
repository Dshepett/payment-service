run:
	go run ./cmd
build:
	go build -o /application ./cmd
docker:
	docker-compose up -d --build
docker-stop:
	docker-compose down
test:
	go test -v ./...