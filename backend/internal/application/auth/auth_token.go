package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

const AccessTokenDuration = time.Minute * 15 
const RefreshTokenDuration = time.Hour * 24 * 7