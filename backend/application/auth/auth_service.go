package auth

import (
	"sub-watch-backend/application/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
    secretKey []byte
}

func NewJWTService(cfg *config.ConfigMap) *JWTService {
    return &JWTService{
        secretKey: []byte(cfg.JWTSecretKey),
    }
}

func (s *JWTService) GenerateAccessToken(userID, email string) (string, error) {
    expirationTime := time.Now().Add(AccessTokenDuration)
    
    claims := &UserClaims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Subject:   userID,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(s.secretKey)
}