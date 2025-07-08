package auth

import (
	"os"
	"time"

	"icomphub-api/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

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
	tokenString, err := token.SignedString(JwtSecret)
	return "Bearer " + tokenString, err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
