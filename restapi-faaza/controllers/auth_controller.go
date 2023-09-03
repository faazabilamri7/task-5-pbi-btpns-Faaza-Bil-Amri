package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "path/to/your/models" // Ganti dengan package yang sesuai
    "path/to/your/helpers" // Ganti dengan package yang sesuai
)

// LoginUserHandler menangani permintaan POST untuk otentikasi pengguna dan pembuatan token JWT
func LoginUserHandler(c *gin.Context) {
    // Ambil data login dari permintaan
    var loginData models.User
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cari pengguna berdasarkan email
    user, err := models.GetUserByEmail(db, loginData.Email)
    if err != nil || !helpers.VerifyPassword(loginData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Email atau password salah"})
        return
    }

    // Buat token JWT
    token, err := helpers.CreateToken(*user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil otentikasi, kirim token JWT sebagai respons
    c.JSON(http.StatusOK, gin.H{"token": token})
}
