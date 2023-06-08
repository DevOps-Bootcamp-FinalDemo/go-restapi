package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRootRoute(t *testing.T) {
	r := setupRouter()

	// Create a new request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	// Convert the response body to a map
	var expectedBody = gin.H{
		"message": "Hello from FINAL DEMO presentation (June 8)!!!",
	}

	currentBody := gin.H{}
	if err := json.Unmarshal(rr.Body.Bytes(), &currentBody); err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}

	// Check the response body
	if !reflect.DeepEqual(currentBody, expectedBody) {
		jsonExpected, _ := json.Marshal(expectedBody)
		t.Errorf("Expected body %v, got %v\n", string(jsonExpected), rr.Body)
	}
}
