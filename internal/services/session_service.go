package services

import (
	"time"

	"github.com/INTLSavio/auth-service/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	UserRepository *repositories.UserRepository
}

func NewSessionService(userRepository *repositories.UserRepository) *SessionService {
	return &SessionService{
		UserRepository: userRepository,
	}
}

func (sessionService *SessionService) Login(email, password string) (string, error) {
	secretKey := []byte("secret")

	user, err := sessionService.UserRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"clientId": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
