package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oddsteam/jongshirts/internal/sessions"
	"github.com/oddsteam/jongshirts/web/handlers"
)

func TestHomeHandlerRedirectHaveSession(t *testing.T) {
	// Create a new request with nil body since we don't need to pass any data
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the HomeHandler function with the mock objects
	handlers.HomeHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusSeeOther {
		t.Errorf("expected status code %d but got %d", http.StatusSeeOther, rr.Code)
	}

}

func TestHomeHandler(t *testing.T) {
	handlers.SetTemplateDir("../templates")
	// Create a new request with nil body since we don't need to pass any data
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	rr := httptest.NewRecorder()

	session,_ := sessions.NewSession(req)

	session.Values["username"] = "testuser"
	session.Save(req, rr)
	// session.Values["username"] = "testuser"

	// Call the HomeHandler function with the mock objects
	handlers.HomeHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rr.Code)
	}

}