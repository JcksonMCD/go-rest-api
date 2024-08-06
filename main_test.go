package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
