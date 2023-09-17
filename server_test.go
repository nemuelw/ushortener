package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShortenURL(t *testing.T) {
	r := gin.Default()
	r.POST("/shorten", shortenURL)

	// test request with JSON payload
	req, err := http.NewRequest("POST", "/shorten", 
						strings.NewReader(`{"longURL":"https:some.long/url"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// test response recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedJSON := `{"shortURL":`
	if !strings.Contains(w.Body.String(), expectedJSON) {
		t.Errorf("Expected response body to contain %s, got %s", expectedJSON, w.Body.String())
	}
}

func TestRedirectURL(t *testing.T) {
	r := gin.Default()
	r.GET("/:hash", redirectToURL)

	// test case 1
	urlMap["exists"] = "https://url.exists"
	req1, err := http.NewRequest("GET", "/exists", nil)
	if err != nil {
		t.Fatal(err)
	}

	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)
	if w1.Code != http.StatusPermanentRedirect {
		t.Errorf("Expected status code %d, got %d", http.StatusPermanentRedirect, w1.Code)
	}
	expectedURL := "https://url.exists"
	if location := w1.Header().Get("Location"); location != expectedURL {
		t.Errorf("Expected Location header to be %s, got %s", expectedURL, location)
	}

	// test case 2
	req2, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	if w2.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w2.Code)
	}
}
