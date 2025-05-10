# User Registration and Login API

This is a WOLFTAGON BACKEND ASSESMENT TEST , this project provides a simple API for user registration and login functionality using the Gin framework in Go. It includes JWT-based authentication and uses GORM for database interaction.

## Features

- User Registration
- User Login
- Password hashing using bcrypt
- JWT token generation for authentication
- Environment variable support through `.env`

---

## Prerequisites

- Go installed on your machine
- PostgreSQL database
- `.env` file with the following variables:
  ```env
  DB_URL=your_postgres_connection_string
  SECRET=your_jwt_secret_key
  ```

---

## Installation

1. Clone the repository:

   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up your `.env` file:
   Create a `.env` file in the root of your project and add the required environment variables:

   ```env
   DB_URL=your_postgres_connection_string
   SECRET=your_jwt_secret_key
   ```

4. Run database migrations:
   ```bash
   go run main.go
   ```

---

## Running the Application

1. Start the server:

   ```bash
   go run main.go
   ```

2. The server will run on `http://localhost:8080`.

---

## API Endpoints

### 1. Register a User

**Endpoint:**  
`POST /users/register`

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```

**Response:**

- Success (`200`):

```json
{
  "user": {
    "ID": 1,
    "CreatedAt": "2025-05-10T18:50:14Z",
    "UpdatedAt": "2025-05-10T18:50:14Z",
    "DeletedAt": null,
    "Name": "John Doe",
    "Email": "john.doe@example.com",
    "Password": "hashed_password"
  }
}
```

- Error (`400`):
  - If the email already exists:
  ```json
  {
    "error": "Email already exists"
  }
  ```
  - If the password cannot be hashed:
  ```json
  {
    "error": "Cannot Generate Password"
  }
  ```

---

### 2. Login a User

**Endpoint:**  
`POST /users/login`

**Request Body:**

```json
{
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```

**Response:**

- Success (`200`):

```json
{
  "token": "your_jwt_token",
  "userEmail": "john.doe@example.com"
}
```

- Error (`400`):
  - If the email is invalid:
  ```json
  {
    "error": "Invalid Email or password"
  }
  ```
  - If the password is incorrect:
  ```json
  {
    "error": "Incorrect Password"
  }
  ```

---

## Project Structure

```
.
├── controllers
│   └── UserController.go         # Handles user registration and login
├── initializers
│   ├── database.go           # Handles database connection
│   ├── loadEnvVariables.go          # Loads environment variables
├── models
│   └── model.go         # Defines the User model
├── migrate
│   └── migrate.go      # automatically creates the `users` table in the database.
├── main.go             # Entry point of the application
├── .env                # Environment variables (not included in the repo)
└── go.mod              # Go module file
```

---

## Notes

- Ensure that the `SECRET` in your `.env` file is a secure string for JWT token signing.
- The database connection string in `DB_URL` should be valid and point to your PostgreSQL database.
- The `initializers.DB.AutoMigrate(&models.Users{})` in the migration script automatically creates the `users` table in the database.

---

## License

This project is licensed under the MIT License. Feel free to use and modify it as per your needs.
