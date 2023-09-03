package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "path/to/your/models" // Ganti dengan package yang sesuai
    "path/to/your/helpers" // Ganti dengan package yang sesuai untuk otentikasi JWT
)

// RegisterUserHandler menangani permintaan POST untuk mendaftarkan pengguna baru
func RegisterUserHandler(c *gin.Context) {
    // Ambil data pengguna dari permintaan
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lakukan validasi data pengguna
    if err := helpers.ValidateUser(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan pengguna dalam basis data
    if err := models.CreateUser(db, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil mendaftarkan pengguna
    c.JSON(http.StatusCreated, gin.H{"message": "Pengguna berhasil didaftarkan"})
}

// GetUserHandler menangani permintaan GET untuk mendapatkan data pengguna
func GetUserHandler(c *gin.Context) {
    // Ambil ID pengguna dari parameter URL
    userID := c.Param("userId")

    // Lakukan otentikasi JWT di sini jika diperlukan

    // Cari pengguna dalam basis data berdasarkan ID
    user, err := models.GetUserByID(db, userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
        return
    }

    // Berhasil mendapatkan data pengguna
    c.JSON(http.StatusOK, user)
}

// UpdateUserHandler menangani permintaan PUT untuk memperbarui data pengguna
func UpdateUserHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil data pengguna dari permintaan
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lakukan validasi data pengguna
    if err := helpers.ValidateUser(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan perubahan data pengguna dalam basis data
    if err := models.UpdateUser(db, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil memperbarui data pengguna
    c.JSON(http.StatusOK, gin.H{"message": "Data pengguna berhasil diperbarui"})
}

// DeleteUserHandler menangani permintaan DELETE untuk menghapus pengguna
func DeleteUserHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil ID pengguna dari parameter URL
    userID := c.Param("userId")

    // Hapus pengguna dari basis data
    if err := models.DeleteUser(db, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil menghapus pengguna
    c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}
