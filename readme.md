# __GraphQL API with Go__
This repository contains an example implementation of a GraphQL API using Go (Golang). The API allows performing CRUD operations on a "books" table in a MySQL database.

## Setup
Before running the GraphQL API, you need to have Go (Golang) and MySQL installed on your system. Follow the instructions in the "Prerequisites" section of the README to set up the project.

## Creating the Database
To use the GraphQL API, you need to create a MySQL database and a "books" table. The necessary SQL commands are provided in the "Setup" section of the README.

## Running the API
Once the database is set up, you can build and run the GraphQL API server. The main steps are outlined in the "Setup" section of the README. After running the API server, it will be accessible at http://localhost:8080/graphql.

## Usage
The GraphQL API supports the following operations:

- Get all books: Sends a GET request to /graphql with a query to retrieve all books.

- Get a book by ID: Sends a GET request to /graphql with a query to retrieve a specific book by its ID.

- Create a new book: Sends a POST request to /graphql with a mutation to create a new book.

- Update a book: Sends a POST request to /graphql with a mutation to update an existing book by its ID.

- Delete a book: Sends a POST request to /graphql with a mutation to delete a book by its ID.

## Prerequisites

Before running the GraphQL API, make sure you have the following installed on your system:

- Go (Golang)
- MySQL

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/graphql-api-go.git

2. Navigate to the project directory:

    ```bash
    cd graphql-api-go
3. Install the necessary dependencies:

    ```bash
    go get github.com/graphql-go/graphql
    go get github.com/graphql-go/handler
    go get github.com/go-sql-driver/mysql

4. Update the MySQL connection details:

    Open the main.go file and replace the placeholder "user:password@tcp(localhost:3306)/database" in the sql.Open call with your MySQL connection details.

5. Create the MySQL database and table:

    Run the following SQL commands in your MySQL client to create the necessary database and table:

    ```sql
    CREATE DATABASE booksdb;
    USE booksdb;

    CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
    );

6. Build and run the API server:

    ```bash
    go build
    ./graphql-api-go

7. The GraphQL API server should now be running at http://localhost:8080/graphql.

# Usage
You can interact with the GraphQL API using tools like GraphiQL or Postman. The API supports the following operations:

## Get all books
Send a GET request to http://localhost:8080/graphql with the following query:


```graphql
    query {
    books {
        id
        title
        author
    }
    }
```

## Get a book by ID
Send a GET request to http://localhost:8080/graphql with the following query, replacing <book-id> with the actual book ID:


```graphql
    query {
    book(id: <book-id>) {
        id
        title
        author
    }
    }
```

## Create a new book
Send a POST request to http://localhost:8080/graphql with the following mutation:

```graphql
    mutation {
    createBook(title: "New Book", author: "New Author") {
        id
        title
        author
    }
    }
```

## Update a book
Send a POST request to http://localhost:8080/graphql with the following mutation, replacing <book-id> with the actual book ID and <new-title> and <new-author> with the updated values:

```graphql
    mutation {
    updateBook(id: <book-id>, title: "<new-title>", author: "<new-author>") {
        id
        title
        author
    }
    }
```

## Delete a book
Send a POST request to http://localhost:8080/graphql with the following mutation, replacing <book-id> with the actual book ID:

```graphql
    mutation {
    deleteBook(id: <book-id>) {
        id
        title
        author
    }
    }
```