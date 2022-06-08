run:
	go run ./cmd
build:
	go build -o bin/app ./cmd
build-win:
	go build -o bin/app.exe ./cmd