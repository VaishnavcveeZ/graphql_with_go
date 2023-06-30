# __Introduction to GraphQL__
GraphQL is an open-source query language and runtime for APIs (Application Programming Interfaces) that was developed by Facebook. It provides a flexible and efficient way to query and manipulate data, offering several advantages over traditional RESTful APIs. This guide provides a beginner-friendly introduction to GraphQL, explaining its core concepts and features.

## _What is GraphQL?_
GraphQL is a query language that allows clients to specify the structure of the data they need from an API. It provides a single endpoint for clients to send their queries and receive the requested data in a predictable and efficient manner. Unlike RESTful APIs, where multiple requests might be required to fetch related data, GraphQL enables clients to retrieve all the required data in a single request.

## _Core Concepts of GraphQL_
1. Schema: At the heart of GraphQL is the schema, which defines the capabilities of the API and describes the data models and operations that can be performed. It consists of two main parts: the object types and the root types.

2. Object Types: GraphQL defines object types, which represent entities in the system being modeled. Each object type has fields that specify the data it contains. For example, in an e-commerce system, there might be object types for Product, User, and Order.

3. Queries: GraphQL queries are used to request data from the API. Clients can specify exactly what data they need and the structure of the response by defining the fields they want to retrieve. Queries are executed against the API's schema, and the response mirrors the structure of the query.

4. Mutations: Mutations are used to modify data on the server. They allow clients to create, update, or delete data. Like queries, mutations are defined in the schema and can be executed by clients.

5. Resolvers: Resolvers are functions that are responsible for fetching the data requested in a query. Each field in a GraphQL query is associated with a resolver, which determines how the data should be retrieved. Resolvers can fetch data from databases, call external services, or perform any other necessary operations to fulfill the query.

## _Advantages of GraphQL_
1. Efficiency: GraphQL reduces over-fetching and under-fetching of data. Clients can request only the data they need, eliminating unnecessary data transfer and reducing bandwidth usage.

2. Flexibility: With GraphQL, clients have fine-grained control over the data they receive. They can request multiple related resources in a single query and avoid making multiple round trips to the server.

3. Versioning: GraphQL avoids versioning issues common in RESTful APIs. As the schema evolves, clients can add or modify their queries without relying on changes to the server's API endpoints.

4. Documentation: GraphQL provides self-documenting APIs. The schema serves as a contract between the server and clients, making it easy to understand the available data and operations. Tools can generate comprehensive documentation automatically.

## _Getting Started with GraphQL_
To start using GraphQL, you typically need to:

1. Define the Schema: Define the object types, queries, and mutations in the GraphQL schema. Specify the fields and their types, along with any relationships between object types.

2. Implement Resolvers: Implement the resolver functions for each field in the schema. Resolvers fetch the requested data from the appropriate data sources, such as databases or external services.

3. Set up the Server: Set up a GraphQL server that uses the schema and resolvers. The server receives queries and mutations, resolves them, and returns the requested data.

4. Interact with the API: Clients can interact with the GraphQL API by sending queries and mutations to the server. They receive the requested data in response, structured according to their query.

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