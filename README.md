# Book Service - gRPC API

This is the **Book Service** of the gRPC microservice architecture. This service manages books, including operations like adding, updating, fetching, and deleting books from the system. It also supports retrieving related category data through a separate Category Service.

## Features

- **Add Book**: Create a new book record in the database.
- **List Books**: Retrieve a list of all books.
- **Get Book by ID**: Retrieve a single book by its unique identifier.
- **Update Book**: Update book details like title, author, and description.
- **Delete Book**: Remove a book by its ID.

## Requirements

- **Go** (version 1.23 or later)
- **gRPC** (Go implementation)
- **PostgreSQL** for database management

## Setup and Installation

### 1. Clone the Repository

```bash
git clone https://github.com/your-repository/book-service.git
cd book-service
```

### 2. Install Dependencies
Ensure you have all required dependencies by running:

```bash
go mod tidy
```

### 3. Setup PostgreSQL Database
Make sure PostgreSQL is installed and configured. You can use the provided dsn in the .env code or modify it as needed:

```bash
host=localhost port=5432 user=postgres dbname=library password=secret sslmode=disable
```

You can also create the database manually:

```bash
CREATE DATABASE library;
```

### 4. Database Migration
If you have any migration tools, apply the database migrations for the books table.

### 5. Running the Service
To start the gRPC server, run:

```bash
go run main.go
```

The service will start on the default gRPC port :1200.