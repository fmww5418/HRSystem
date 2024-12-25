# Simple HR System

This is a simple HR system written using Clean Architecture. The system uses MySQL and Redis as the database and caching system, and is deployed via Docker Compose.

## System Architecture

- **Clean Architecture**: Divides the system into multiple layers to ensure the independence and testability of each layer.
- **MySQL**: Used as the main database system.
- **Redis**: Used for caching and other data that requires fast access.
- **Docker Compose**: Used to manage and deploy multiple containerized services.

## Makefile Commands

Below are some commonly used Makefile commands and their purposes:

- `make docker-up`: Start all Docker containers.
- `make docker-down`: Stop all Docker containers.
- `make migrate`: Generate database auto migration files.
   - Ensure you have installed https://github.com/ariga/atlas before running it
- `make seed`: Execute database seeding.
- `make test`: Run unit tests.
- `make run`: Run the main program.
- `make clean`: Clean up build output files.
- `make mockery`: Generate mock files.
   - Ensure you have installed mockery by executing `go install github.com/vektra/mockery/v2@v2.50.1` before running it
- `make swagger`: Update and generate swagger files.
   - Ensure you have installed swag by executing `go install github.com/swaggo/swag/cmd/swag@v1.16.3 && swag -v` before running it  

## Deployment

1. Ensure Docker and Docker Compose are installed.
2. Use the `make docker-up` command to start all services.
3. The system will run on `http://localhost:8080`.
4. You can view and interact with the swagger documentation at `http://localhost:8080/docs/index.html`.

## Environment Variables

The following environment variables are defined in `docker-compose.yml`:

- `DB_HOST`: Database host.
- `DB_PORT`: Database port.
- `DB_USERNAME`: Database username.
- `DB_PASSWORD`: Database password.
- `DB_NAME`: Database name.
- `REDIS_HOST`: Redis host.
- `REDIS_PORT`: Redis port.

Please ensure these variables are set in the `.env` file.

## Notes

1. For newly registered accounts that have not joined an organization, please use the Create Organization API to establish an organization. Otherwise, there will be no data in the department or employee sections.
2. The token required for API calls can be obtained after logging in and should be used in the authorization header in the format "Bearer ${token}".

- Please ensure MySQL and Redis are correctly configured in the local environment.
- Before performing database migration or seeding operations, ensure the database connection is normal.