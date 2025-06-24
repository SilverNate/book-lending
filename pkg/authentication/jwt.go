package authentication

import (
	"book-lending-api/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(config *config.EnvConfig) *JwtService {
	return &JwtService{
		config.JWTSecret,
		config.JWTIssuer}
}

func (s *JwtService) GenerateToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
		"iss":     s.issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}
