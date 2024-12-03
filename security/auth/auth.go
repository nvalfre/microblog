package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Secret key for signing JWTs (keep this safe in a secure environment).
var jwtSecret = []byte("your-strong-secret-key")

// Claims defines the JWT claims structure.
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateSignedToken creates a new JWT token for the specified user ID.
func GenerateSignedToken(userID string) (string, error) {
	const expiration = 24 * time.Hour
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken parses and validates a JWT token, returning the user ID if valid.
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", jwt.ErrTokenMalformed
	}

	return claims.UserID, nil
}

// JwtSignedTokenMiddleware ensures requests are authenticated.
func JwtSignedTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing Authorization header"})
			c.Abort()
			return
		}

		userID, err := ValidateToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token", "message": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
