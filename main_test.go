package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Assert that the response body is as expected
	expected := `[{"id":"1","title":"In Search of Lost Time","author":"Marcel Proust","quantity":2},{"id":"2","title":"The Great Gatsby","author":"F. Scott Fitzgerald","quantity":5},{"id":"3","title":"War and Peace","author":"Leo Tolstoy","quantity":6}]
		`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s, but got %s", expected, w.Body.String())
	}
}
