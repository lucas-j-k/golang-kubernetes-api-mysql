package user

import (
	"net/http"

	"github.com/lucas-j-k/kube-go-api/authTools"

	"github.com/lucas-j-k/kube-go-api/httpTools"

	"github.com/go-chi/render"
)

// Logout checks for the presence of a session Id in the request cookies. If one is found, it clears any
// session data for the session in the cache, and clears the session cookie. Subsequent requests to
// protected routes should fail
func Logout(service *UserService, sessionManager authTools.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// If no sessionId found in cookies, ignore and return
		sessionCookie, err := r.Cookie("user-session")
		if err != nil {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, nil)
			return
		}

		sessionId := sessionCookie.Value

		// Delete session cookie
		authTools.ClearAuthCookie(w)

		// Clear the session data from the cache based on the session Id
		err = sessionManager.ClearSession(ctx, sessionId)
		if err != nil {
			render.Render(w, r, httpTools.ErrInternalServer())
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, nil)
	}
}
