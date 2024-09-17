# AI Agent Microservices Blueprint

## Overview

AI Agent Microservices Blueprint is a robust, scalable backend service built with Go and PocketBase. It provides a solid foundation for developing AI agents in a microservices architecture, with features like client authentication, project management, and extensible AI agent integration.

## Features

- Client Authentication: Secure API access using client ID and secret
- Project Management: Create, list, and check status of AI agent projects
- AI Agent Integration: Extensible architecture for integrating various AI agents
- Database Migrations: Automated schema updates and initial data seeding
- Microservices Architecture: Designed for scalability and modularity
- CRUD Testing: Bash script for testing API endpoints

## Prerequisites

- Go 1.22.4 or later
- PocketBase
- Make (for using Makefile commands)
- Bash (for running the test script)
- curl (for API testing)
- jq (for parsing JSON responses in the test script)
- Docker (for containerization and microservices deployment)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/ai-fab/agent-blueprint.git
   cd agent-blueprint
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

5. Build Docker images for microservices:
   ```
   docker-compose build
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

## Running the Services

Start the microservices:

```
docker-compose up
```

The agent API will be available at `http://localhost:8090/api/...`.
Pocketbase Admin UI will be available at `http://localhost:8090/_/`.


## API Endpoints

- `POST /api/projects`: Create a new AI agent project
- `GET /api/projects`: List AI agent projects for a client
- `GET /api/projects/:id/status`: Get the status of a specific AI agent project
- `POST /api/agents/:agent_type/execute`: Execute an AI agent task

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
- `agent/`: AI agent integration code
- `test_crud.sh`: CRUD testing script
- `docker-compose.yml`: Docker Compose configuration for microservices

### Adding New AI Agents

1. Create a new directory in `agent/` for the AI agent type
2. Implement the agent's logic and API in the new directory
3. Update the main integrate the new AI agent logic
4. Add new routes and handlers for the AI agent's functionality
5. Update the `test_crud.sh` script to test the new AI agent endpoints

## Makefile Commands

- `make build`: Build the main service
- `make run`: Run the main service (for local development)
- `make test`: Run Go tests
- `make clean`: Remove the built binary
- `make clean-db`: Remove the PocketBase data directory
- `make migrate`: Run database migrations
- `make migrate-down`: Revert database migrations
- `make migrate-fresh`: Revert and re-run all migrations
- `make test-crud`: Run the CRUD test script
- `make docker-build`: Build all Docker images
- `make docker-up`: Start all Docker services
- `make docker-down`: Stop all Docker services

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
