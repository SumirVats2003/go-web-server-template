# Go Web Server Template

A template for building web servers in Go, featuring a structured layout and common functionalities.

## Features

*   **Structured Project Layout:** Organized into `internal`, `models`, and `utils` directories for better code management.
*   **Authentication:** Includes JWT-based authentication for securing endpoints.
*   **CORS Middleware:** Provides Cross-Origin Resource Sharing (CORS) middleware for handling requests from different origins. It is configured to allow all origins (`*`), and the following HTTP methods: `GET, POST, PUT, PATCH, DELETE, OPTIONS`. It also allows the `Content-Type` and `Authorization` headers.
*   **Database Integration:** Uses `pgx/v5` for connecting to a PostgreSQL database.
*   **Environment Variable Management:** Utilizes `godotenv` for managing environment variables. It loads environment variables from a `.env` file into the application. The project also includes a custom `GetEnv` function in the `utils` package to retrieve environment variables with a default value if they are not set.

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

*   Go (version 1.24.1 or later)
*   PostgreSQL

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/SumirVats2003/go-web-server-template.git
    ```

2.  **Navigate to the project directory:**

    ```sh
    cd go-web-server-template
    ```

3.  **Install dependencies:**

    ```sh
    go mod tidy
    ```

4.  **Set up environment variables:**

    Create a `.env` file in the root directory and add the following variables:

    ```env
    DATABASE_URL="your_database_url"
    JWT_SECRET="your_jwt_secret"
    ```

5.  **Run the application:**

    ```sh
    go run main.go
    ```

## Environment Variables

*   `DATABASE_URL`: The connection string for your PostgreSQL database.
*   `JWT_SECRET`: The secret key for signing JWT tokens.

## Dependencies

*   **`go-chi/chi`**: A lightweight, idiomatic and composable router for building Go HTTP services.
*   **`golang-jwt/jwt`**: A Go implementation of JSON Web Tokens (JWT).
*   **`jackc/pgx`**: A pure Go driver and toolkit for PostgreSQL.
*   **`joho/godotenv`**: A Go port of the Ruby dotenv library (loads environment variables from a `.env` file).

## API Endpoints

*   **`/`**: Heartbeat endpoint to check the server's status.
*   **/auth/login**: Log in an existing user.
*   **/auth/signup**: Register a new user.

## Database

This project uses **PostgreSQL** as its database. The `pgx/v5` driver is used for database connectivity.
