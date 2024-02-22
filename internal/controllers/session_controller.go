package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/INTLSavio/auth-service/internal/dto"
	"github.com/INTLSavio/auth-service/internal/services"
	"github.com/INTLSavio/auth-service/utils"
)

type SessionController struct {
	SessionService *services.SessionService
}

func NewSessionController(sessionService *services.SessionService) *SessionController {
	return &SessionController{
		SessionService: sessionService,
	}
}

func (sessionController *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	var loginParams dto.LoginUserInput

	err := json.NewDecoder(r.Body).Decode(&loginParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := sessionController.SessionService.Login(loginParams.Email, loginParams.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := utils.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	accessToken := dto.LoginResponse{
		AccessToken: result,
	}

	json.NewEncoder(w).Encode(accessToken)
}
