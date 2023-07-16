package authTools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
)

// UserSession represents the JSON object which is saved under the session in our cache
type UserSession struct {
	UserID int64 `json:"userId"`
}

// SessionManager declares the interface for managing user sessions
type SessionManager interface {

	// SetSession saves the current user session details into cache, under the sessionId key.
	SetSession(ctx context.Context, sessionId string, sessionDetails UserSession) error

	// GetSession queries the cache for a session stored under a given sessionId
	// and parses the stored session JSON back into a usable struct
	GetSession(ctx context.Context, sessionId string) (UserSession, error)

	// ClearSession deletes a session saved in cache under a given sessionId key
	ClearSession(ctx context.Context, sessionId string) error
}

// RedisSessionManager implements SessionManager, with Redis as the store
type RedisSessionManager struct {
	Client *redis.Client
}

func (sessionManager RedisSessionManager) SetSession(ctx context.Context, sessionId string, sessionDetails UserSession) error {
	sessionString, err := json.Marshal(sessionDetails)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("session:%v", sessionId)
	err = sessionManager.Client.Set(ctx, cacheKey, sessionString, 10*time.Minute).Err()

	if err != nil {
		return err
	}

	return nil
}

func (sessionManager RedisSessionManager) GetSession(ctx context.Context, sessionId string) (UserSession, error) {
	sessionDetails := UserSession{}
	cacheKey := fmt.Sprintf("session:%v", sessionId)

	value, err := sessionManager.Client.Get(ctx, cacheKey).Result()

	if err != nil {
		return sessionDetails, err
	}

	// unmarshall redis JSON object into a usable struct
	err = json.Unmarshal([]byte(value), &sessionDetails)
	if err != nil {
		return sessionDetails, err
	}

	return sessionDetails, nil

}

func (sessionManager RedisSessionManager) ClearSession(ctx context.Context, sessionId string) error {
	cacheKey := fmt.Sprintf("session:%v", sessionId)
	err := sessionManager.Client.Del(ctx, cacheKey).Err()
	return err
}

type CacheMiddleware struct {
	SessionManager SessionManager
}

// SessionGuard checks for a session cookie on an incoming request. If found, it checks for a valid session under that ID in the cache.
// If session is valid, it forwards the request to the next handler in the chain. If not, it blocks with a 401 unauthorized
func (middleware *CacheMiddleware) SessionGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("user-session")

		// No cookie found, return immediately and block
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, nil)
			return
		}

		sessionId := cookie.Value

		// Attempt to pull the session data from our session cache (redis)
		sessionData, err := middleware.SessionManager.GetSession(r.Context(), sessionId)

		if err != nil {
			// User session is not present or has expired
			ClearAuthCookie(w)
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, nil)
			return
		}

		// Assume there is now a valid session. Append the userId to the context to make it available in upstream handlers
		ctx := context.WithValue(r.Context(), "userId", sessionData.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
