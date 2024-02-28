package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

func NewSession(r *http.Request) (*sessions.Session, error) {
	session, err := Store.Get(r, "session")
	if session.IsNew {
		session.Options.Domain = "localhost"
		session.Options.MaxAge = 0
		session.Options.Secure = true
	}
	return session, err
}
