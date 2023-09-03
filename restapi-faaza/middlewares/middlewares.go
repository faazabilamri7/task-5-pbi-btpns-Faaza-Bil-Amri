package middlewares

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "path/to/your/helpers" // Ganti dengan package yang sesuai
)

// AuthMiddleware adalah middleware untuk otentikasi JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Mendapatkan token dari header Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak ada"})
            c.Abort()
            return
        }

        // Mengambil token dari header Authorization
        tokenString := strings.Split(authHeader, " ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret_key"), nil // Ganti dengan secret key yang sesuai
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid"})
            c.Abort()
            return
        }

        if token.Valid {
            // Token valid, lanjutkan ke endpoint yang dilindungi
            c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid"})
            c.Abort()
        }
    }
}
