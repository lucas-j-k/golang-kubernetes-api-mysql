package note

import (
	"net/http"
	"time"
)

// TODO :: move this all into a separate module

// SetAuthCookie writes a sessionId into a http-only cookie
func SetAuthCookie(sessionId string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "user-session",
		Domain:   ".links.localhost",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().UTC().Add(24 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}

// ClearAuthCookie deletes the session cookie, overwriting the expiry
func ClearAuthCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "user-session",
		Domain:   ".links.localhost",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, &cookie)
}
