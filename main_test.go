package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetBooks tests the /books endpoint
func TestGetBooks(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a router with the API routes
	router := gin.Default()
	router.GET("/books", getBooks)

	// Create a request to the /books endpoint
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Define the expected books
	expectedBooks := []book{
		{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
		{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	}

	// Unmarshal the response body
	var responseBooks []book
	err := json.Unmarshal(w.Body.Bytes(), &responseBooks)
	assert.NoError(t, err)

	// Assert that the response body is as expected
	assert.Equal(t, expectedBooks, responseBooks)
}

// TestBookById tests the /books/:id endpoint
func TestBookById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/books/:id", bookById)

	req, _ := http.NewRequest(http.MethodGet, "/books/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBook := book{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2}

	var responseBook book
	err := json.Unmarshal(w.Body.Bytes(), &responseBook)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, responseBook)
}

// TestCreateBook tests the /books endpoint for creating a new book
func TestCreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/books", createBook)

	newBook := `{"id":"4","title":"The Catcher in the Rye","author":"J.D. Salinger","quantity":3}`
	req, _ := http.NewRequest(http.MethodPost, "/books", strings.NewReader(newBook))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdBook book
	err := json.Unmarshal(w.Body.Bytes(), &createdBook)
	assert.NoError(t, err)

	expectedBook := book{ID: "4", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 3}
	assert.Equal(t, expectedBook, createdBook)
}

// TestCheckoutBook tests the /checkout endpoint
func TestCheckoutBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/checkout", checkoutBook)

	req, _ := http.NewRequest(http.MethodPut, "/checkout?id=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBook := book{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 1}

	var responseBook book
	err := json.Unmarshal(w.Body.Bytes(), &responseBook)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, responseBook)
}

// TestReturnBook tests the /return endpoint
func TestReturnBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/return", returnBook)

	req, _ := http.NewRequest(http.MethodPut, "/return?id=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBook := book{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 3}

	var responseBook book
	err := json.Unmarshal(w.Body.Bytes(), &responseBook)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, responseBook)
}
