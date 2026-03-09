# TCP Tunnels Manager

This project implements a TCP tunnels manager in Go. It allows for the creation, management, and monitoring of TCP tunnels, facilitating secure and efficient communication between different network endpoints.

## Features

- **Tunnel Management:** Create, update, and delete TCP tunnel configurations.
- **Database Integration:** Stores tunnel configurations and logs in a PostgreSQL database.
- **Tunnel Logging:** Records connection events and data transfer for each tunnel.
- **Client-Side Tunneling:** Provides functionality for establishing and managing TCP tunnel clients.

## Project Structure

The project follows a clean architecture pattern, separating concerns into distinct layers:

- `cmd/app`: Contains the main application entry point.
- `configs`: Handles application configuration and initialization.
- `internal/application`: Implements the core business logic and use cases (e.g., `tunnels_manager.go`).
- `internal/domain`: Defines the core entities and interfaces (e.g., `tunnel_log.go`, `tunnel_row.go`).
- `internal/infrastructure`: Provides implementations for external concerns like database access (`db`) and TCP tunneling (`tcptunnels`).

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)
- PostgreSQL database

### Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Sinaionline/tcp-tunnels.git
    cd tcp-tunnels
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the project root with your database connection string and other necessary configurations. Refer to `.env.example` (if available) for required variables.

3.  **Initialize the database schema:**
    Run the SQL script to create the necessary tables:
    ```bash
    psql -U your_user -d your_database -f schema.sql
    ```

4.  **Run the application:**
    ```bash
    go run cmd/app/main.go
    ```

## Database Schema

The `schema.sql` file defines the database tables for `tunnel_rows` and `tunnel_logs`:

- `tunnel_rows`: Stores the configuration for each TCP tunnel.
- `tunnel_logs`: Records events and statistics related to tunnel activity.

## Dependencies

- `github.com/lib/pq`: PostgreSQL driver for Go.
- `github.com/joho/godotenv`: For loading environment variables from `.env` file.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.