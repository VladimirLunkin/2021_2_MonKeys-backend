package middleware

import (
	"context"
	"dripapp/configs"
	"dripapp/internal/dripapp/models"
	"dripapp/internal/pkg/logger"
	"net/http"
)

type sessionMiddleware struct {
	sessionRepo models.SessionRepository
}

func (s *sessionMiddleware) SessionMiddleware(l logger.Logger) (mw func(http.Handler) http.Handler) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.DebugLogging("auth middlware")
			var userSession models.Session
			session, err := r.Cookie("sessionId")
			if err != nil {
				userSession = models.Session{
					UserID: 0,
					Cookie: "",
				}
			} else {
				userSession, err = s.sessionRepo.GetSessionByCookie(session.Value)
				if err != nil {
					userSession = models.Session{
						UserID: 0,
						Cookie: "",
					}
				}
			}
			r = r.WithContext(context.WithValue(r.Context(), configs.ContextUserID, userSession))
			next.ServeHTTP(w, r)
		})
	}
}
