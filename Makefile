.PHONY: start clean
start: test run

backend_url=$(shell echo $$backend_url)
test:
	@echo "determining dependencies"
	@echo "checking docker installation"
	@if ! command -v docker &> /dev/null; then \
		echo "Docker is not installed. Please install Docker to proceed."; \
		exit 1; \
	fi
	@echo "Docker is installed."
	@echo "checking birthday-cli installation"
	@if ! command -v ./birthday-cli &> /dev/null; then \
		echo "birthday-cli is not installed. Run 'make install' to install it."; \
		exit 1; \
	fi
	@echo "birthday-cli is installed."
	@echo "checking for backend_url environment variable"

	@if [[ -z "${backend_url}" ]]; then \
		echo "backend_url environment variable is not set. Please set it to proceed, 'export backend_url="http://localhost:8002/"'"; \
		exit 1; \
	fi
	@echo "All dependencies are satisfied."

install:
	@echo "installing dependencies"
	@echo "check how to install docker-dekstop and go version ${required_go_version} or above for your native OS"
	@echo "installing birthday-cli"
	@curl -LO "https://github.com/sanusomya/birthday-cli/releases/download/v0.0.1/birthday-cli"
	@chmod +x birthday-cli
	@echo "dependencies installed"

run: test
	@echo "building API"
	docker compose up -d
	@echo "API is running at port 8002, API-URL = http://localhost:8002"

clean:
	docker compose down
	rm -f birthday-cli
	@echo "Cleaned up"