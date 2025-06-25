package middleware

import (
	"book-lending-api/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// mockery --name=IJWTService         --dir=internal/middleware         --output=internal/middleware/mocks         --with-expec
type IJWTService interface {
	GenerateToken(userID int64, email string) (string, error)
	VerifyToken(token string) (*CustomClaims, error)
}

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

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

func (s *JwtService) VerifyToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}
	return claims, nil
}
