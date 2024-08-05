# Bookstore API

This project implements a simple bookstore API using the Go programming language and the Gin web framework. The API allows users to manage a collection of books by providing endpoints to create, read, and update book records.

## Features

- Retrieve all books
- Retrieve a book by its ID
- Create a new book
- Check out a book (decrement quantity)
- Return a book (increment quantity)

## Endpoints

- `GET /books`: Retrieve all books
- `GET /books/:id`: Retrieve a book by its ID
- `POST /books`: Create a new book
- `PUT /checkout`: Check out a book
- `PUT /return`: Return a book

## Data Structure

The data structure for a book is defined as follows:

```go
type book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
}
