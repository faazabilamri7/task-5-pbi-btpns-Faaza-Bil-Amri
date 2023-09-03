package helpers

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "path/to/your/models" // Ganti dengan package yang sesuai
)

var jwtKey = []byte("secret_key") // Ganti dengan secret key yang sesuai

// Claims adalah struktur yang akan digunakan untuk mengirim data tambahan dalam token JWT
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// CreateToken membuat token JWT berdasarkan ID pengguna
func CreateToken(user models.User) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Durasi token berlaku (misal: 24 jam)

    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
