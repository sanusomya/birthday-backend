.PHONY: start clean

start: build run
build:
	@echo "Building for alpine linux"
	env GOOS=linux GOARCH=amd64 go build -o main main.go
	@echo "Build successfull"
run: build
	@echo "building API"
	docker compose up -d
	@echo "API is running at port 8002, API-URL = http://localhost:8002"

clean:
	rm main
	docker compose down
	@echo "Cleaned up"