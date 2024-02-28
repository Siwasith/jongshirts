package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oddsteam/jongshirts/web/handlers"
)

func TestHomeHandlerRedirect(t *testing.T) {
	// Create a new request with nil body since we don't need to pass any data
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	rr := httptest.NewRecorder()


	// session.Values["username"] = "testuser"

	// Call the HomeHandler function with the mock objects
	handlers.HomeHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusSeeOther {
		t.Errorf("expected status code %d but got %d", http.StatusSeeOther, rr.Code)
	}

}
