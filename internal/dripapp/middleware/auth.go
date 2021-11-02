package middleware

import (
	"context"
	"dripapp/configs"
	"dripapp/internal/dripapp/models"
	"net/http"
)

type sessionMiddleware struct {
	sessionRepo models.SessionRepository
}

func (s *sessionMiddleware) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		r = r.WithContext(context.WithValue(r.Context(), configs.ForContext, userSession))
		next.ServeHTTP(w, r)
	})
}
