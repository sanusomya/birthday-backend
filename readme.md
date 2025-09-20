# Birthday API

A Go-based backend for managing [birthday entries](https://github.com/sanusomya/birthday-cli/tree/feature/local-api), supporting add, edit, delete, and query operations. Designed for local development with MongoDB and easy integration with CLI or web clients.

## Features

- Add, list, edit, and delete birthday entries (name, date, month, mobile).
- Validation for names, dates, and mobile numbers.
- RESTful API endpoints via Echo server.
- Environment-based configuration.
- Unit tests for validation and database logic.

## Prerequisites

- Go 1.21+
- [MongoDB](#run)
- Docker
- [Birthday CLI](https://github.com/sanusomya/birthday-cli/releases)

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/sanusomya/birthday.git
   cd birthday
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

## Usage

- Use the API endpoints to add, edit, delete, and list birthdays.
- Integrate with the [birthday-cli](https://github.com/sanusomya/birthday-cli) for command-line access.

## Testing

Run unit tests:
```sh
go test ./...
```

## Project Structure

- `birthday/` – Birthday model definitions
- `database/` – Database logic and interfaces
- `server/` – HTTP server, endpoints, middleware
- `utils/` – Validation and custom error types

## Example `.env`

```
db_url=mongodb://localhost:27017
database=birthdaydb
db_coll=birthdays
birthday_app_port=8002
```

## Run
```sh
make start
```

- This will build the Go binary for Alpine Linux and start the API server in a Docker container.
- A MongoDB instance will also be created and attached to a volume on your system, ensuring data persists across container restarts and providing redundancy.
- The API will be available at [http://localhost:8002](http://localhost:8002).

### Stop and Clean Up

```sh
make clean
```

- This will remove the built binary and stop the Docker containers.

> **Note:** Ensure Docker is installed and running on your machine before using these commands.

