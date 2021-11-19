package delivery

import (
	"dripapp/internal/dripapp/models"
	"dripapp/internal/pkg/logger"
	_p "dripapp/internal/pkg/permissions"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetUserRouting(loggger logger.Logger, router *mux.Router, us models.UserUsecase, su models.SessionUsecase, ur models.UserRepository) {
	userHandler := &UserHandler{
		Logger:       loggger,
		UserUCase:    us,
		SessionUcase: su,
	}

	userMid := _p.UserMiddlware{
		UserRepo: ur,
	}

	router.HandleFunc("/api/v1/profile", _p.SetCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.CurrentUser)))).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/profile", _p.CheckCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.EditProfileHandler)))).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/profile", _p.SetCSRF(userHandler.SignupHandler)).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/v1/profile/photo", _p.CheckCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.UploadPhoto)))).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/profile/photo", _p.CheckCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.DeletePhoto)))).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/v1/user/cards", _p.SetCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.NextUserHandler)))).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/v1/matches", _p.SetCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.MatchesHandler)))).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/v1/likes", _p.CheckCSRF(_p.CheckAuthenticated(userMid.GetCurrentUser(userHandler.ReactionHandler)))).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/v1/tags", _p.SetCSRF(_p.CheckAuthenticated(userHandler.GetAllTags))).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/v1/notifications", permissions.SetCSRF(permissions.CheckAuthenticated(userHandler.Notifications)))

	router.PathPrefix("/api/documentation/").Handler(httpSwagger.WrapHandler)
}
