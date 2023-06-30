package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

func init() {
	// Initialize the database connection
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/booksdb")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func main() {
	// Close the database connection when the application exits
	defer db.Close()

	// Define GraphQL object types
	bookType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	// Define GraphQL queries
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:        bookType,
				Description: "Get a book by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(int)
					if ok {
						// Query the database for the book with the given ID
						book := Book{}
						err := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
						if err != nil {
							if err == sql.ErrNoRows {
								return nil, nil // Book not found
							}
							return nil, err
						}
						return book, nil
					}
					return nil, nil
				},
			},
			"books": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get all books",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Query the database for all books
					rows, err := db.Query("SELECT id, title, author FROM books")
					if err != nil {
						return nil, err
					}
					defer rows.Close()

					books := []Book{}
					for rows.Next() {
						book := Book{}
						err := rows.Scan(&book.ID, &book.Title, &book.Author)
						if err != nil {
							return nil, err
						}
						books = append(books, book)
					}
					return books, nil
				},
			},
		},
	})

	// Define GraphQL mutations
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createBook": &graphql.Field{
				Type:        bookType,
				Description: "Create a new book",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					title, _ := params.Args["title"].(string)
					author, _ := params.Args["author"].(string)

					// Insert the new book into the database
					result, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", title, author)
					if err != nil {
						return nil, err
					}

					// Get the ID of the newly created book
					id, err := result.LastInsertId()
					if err != nil {
						return nil, err
					}

					book := Book{
						ID:     int(id),
						Title:  title,
						Author: author,
					}
					return book, nil
				},
			},
			"updateBook": &graphql.Field{
				Type:        bookType,
				Description: "Update a book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					title, titleOk := params.Args["title"].(string)
					author, authorOk := params.Args["author"].(string)

					// Update the book in the database
					updateQuery := "UPDATE books SET "
					updateParams := make([]interface{}, 0)

					if titleOk {
						updateQuery += "title = ?, "
						updateParams = append(updateParams, title)
					}
					if authorOk {
						updateQuery += "author = ?, "
						updateParams = append(updateParams, author)
					}

					// Trim the trailing comma and space from the update query
					updateQuery = updateQuery[:len(updateQuery)-2]

					// Add the WHERE clause to specify the book ID
					updateQuery += " WHERE id = ?"
					updateParams = append(updateParams, id)

					_, err := db.Exec(updateQuery, updateParams...)
					if err != nil {
						return nil, err
					}

					// Get the updated book from the database
					book := Book{}
					err = db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
					if err != nil {
						if err == sql.ErrNoRows {
							return nil, nil // Book not found
						}
						return nil, err
					}

					return book, nil
				},
			},
			"deleteBook": &graphql.Field{
				Type:        bookType,
				Description: "Delete a book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					// Get the book from the database before deleting it
					book := Book{}
					err := db.QueryRow("SELECT id, title, author FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Author)
					if err != nil {
						if err == sql.ErrNoRows {
							return nil, nil // Book not found
						}
						return nil, err
					}

					// Delete the book from the database
					_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
					if err != nil {
						return nil, err
					}

					return book, nil
				},
			},
		},
	})

	// Define the GraphQL schema
	schemaConfig := graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create the GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Serve the GraphQL API over HTTP
	http.Handle("/graphql", h)
	log.Println("Server running at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
