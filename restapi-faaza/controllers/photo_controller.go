package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "path/to/your/models" // Ganti dengan package yang sesuai
    "path/to/your/helpers" // Ganti dengan package yang sesuai untuk otentikasi JWT
)

// CreatePhotoHandler menangani permintaan POST untuk membuat foto baru
func CreatePhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil data foto dari permintaan
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lakukan validasi data foto
    if err := helpers.ValidatePhoto(photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan foto dalam basis data
    if err := models.CreatePhoto(db, &photo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil membuat foto baru
    c.JSON(http.StatusCreated, gin.H{"message": "Foto berhasil dibuat"})
}

// GetPhotoHandler menangani permintaan GET untuk mendapatkan data foto
func GetPhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil ID foto dari parameter URL
    photoID := c.Param("photoId")

    // Cari foto dalam basis data berdasarkan ID
    photo, err := models.GetPhotoByID(db, photoID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
        return
    }

    // Berhasil mendapatkan data foto
    c.JSON(http.StatusOK, photo)
}

// UpdatePhotoHandler menangani permintaan PUT untuk memperbarui data foto
func UpdatePhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil data foto dari permintaan
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan perubahan data foto dalam basis data
    if err := models.UpdatePhoto(db, &photo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil memperbarui data foto
    c.JSON(http.StatusOK, gin.H{"message": "Data foto berhasil diperbarui"})
}

// DeletePhotoHandler menangani permintaan DELETE untuk menghapus foto
func DeletePhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil ID foto dari parameter URL
    photoID := c.Param("photoId")

    // Hapus foto dari basis data
    if err := models.DeletePhoto(db, photoID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil menghapus foto
    c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil dihapus"})
}

// UploadPhotoHandler menangani permintaan POST untuk mengunggah gambar (foto)
func UploadPhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil data foto dari permintaan
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Lakukan validasi data foto
    if err := helpers.ValidatePhoto(photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan foto dalam basis data
    if err := models.CreatePhoto(db, &photo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil mengunggah gambar (foto)
    c.JSON(http.StatusCreated, gin.H{"message": "Gambar berhasil diunggah"})


	// DeletePhotoHandler menangani permintaan DELETE untuk menghapus gambar (foto)
func DeletePhotoHandler(c *gin.Context) {
    // Lakukan otentikasi JWT di sini jika diperlukan

    // Ambil ID foto dari parameter URL
    photoID := c.Param("photoId")

    // Hapus foto dari basis data
    if err := models.DeletePhoto(db, photoID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Berhasil menghapus gambar (foto)
    c.JSON(http.StatusOK, gin.H{"message": "Gambar berhasil dihapus"})
}
