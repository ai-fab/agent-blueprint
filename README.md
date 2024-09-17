# Service Blueprint

## Overview

Service Blueprint is a robust, scalable backend service built with Go and PocketBase. It provides a solid foundation for building client-specific project management applications with features like client authentication, project creation, listing, and status retrieval.

## Features

- Client Authentication: Secure API access using client ID and secret
- Project Management: Create, list, and check status of projects
- Database Migrations: Automated schema updates and initial data seeding
- Extensible Architecture: Easy to add new features and endpoints
- CRUD Testing: Bash script for testing API endpoints

## Prerequisites

- Go 1.22.4 or later
- PocketBase
- Make (for using Makefile commands)
- Bash (for running the test script)
- curl (for API testing)
- jq (for parsing JSON responses in the test script)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/service-blueprint.git
   cd service-blueprint
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Copy the example environment file and edit it with your settings:
   ```
   cp .env.example .env
   ```

4. Build the project:
   ```
   make build
   ```

## Configuration

Edit the `.env` file to set your environment variables:

- `ADMIN_EMAIL`: Email for the initial admin user
- `ADMIN_PASSWORD`: Password for the initial admin user

## Database Migrations

Run migrations to set up the database schema and seed initial data:

```
make migrate
```

This will create the necessary tables and add two test clients:
- Client 1: ID: `test_client_1`, Secret: `test_secret_1`
- Client 2: ID: `test_client_2`, Secret: `test_secret_2`

## Running the Service

Start the service:

```
make run
```

The service will be available at `http://localhost:8090`.

## API Endpoints

- `POST /api/projects`: Create a new project
- `GET /api/projects`: List projects for a client
- `GET /api/projects/:id/status`: Get the status of a specific project

All endpoints require client authentication headers:
- `X-Client-ID`: The client's ID
- `X-Client-Secret`: The client's secret

## Testing

Run the CRUD test script:

```
make test-crud
```

This will start the server, run a series of API tests, and then stop the server.

## Development

### Project Structure

- `main.go`: Entry point of the application
- `config/`: Configuration related code
- `handlers/`: HTTP request handlers
- `middleware/`: Custom middleware (e.g., client authentication)
- `migrations/`: Database migration files
- `models/`: Data models
- `test_crud.sh`: CRUD testing script

### Adding New Features

1. Create new migration files in the `migrations/` directory if needed
2. Add new handlers in the `handlers/` directory
3. Register new routes in `main.go`
4. Update the `test_crud.sh` script to test new endpoints

## Makefile Commands

- `make build`: Build the service
- `make run`: Run the service
- `make test`: Run Go tests
- `make clean`: Remove the built binary
- `make clean-db`: Remove the PocketBase data directory
- `make migrate`: Run database migrations
- `make migrate-down`: Revert database migrations
- `make migrate-fresh`: Revert and re-run all migrations
- `make test-crud`: Run the CRUD test script

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [PocketBase](https://pocketbase.io/) for the backend framework
- [Echo](https://echo.labstack.com/) for the web framework
- All contributors who have helped shape this project
