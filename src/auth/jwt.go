package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"icomphub-api/codes"
	"icomphub-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID     uint64 `json:"user_id"`
	SystemRole string `json:"system_role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
	claims := Claims{
		UserID:     user.ID,
		SystemRole: string(user.SystemRole),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GetTokenData(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return getJWTSecret(), nil
	})
}

func GetUserIDFromToken(context *gin.Context) (uint64, codes.Code, error) {
	tokenString := extractTokenFromHeader(context)
	if tokenString == "" {
		return 0, codes.AuthMissingToken, errors.New("missing or malformed token")
	}

	token, err := GetTokenData(tokenString)

	if err != nil || !token.Valid {
		return 0, codes.AuthInvalidToken, fmt.Errorf("invalid token: %v", err)
	}

	claims := token.Claims.(*Claims)

	return claims.UserID, codes.AuthValidToken, nil
}

func extractTokenFromHeader(context *gin.Context) string {
	authHeader := context.GetHeader("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		return authHeader[7:]
	}
	return ""
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return []byte(secret)
}
