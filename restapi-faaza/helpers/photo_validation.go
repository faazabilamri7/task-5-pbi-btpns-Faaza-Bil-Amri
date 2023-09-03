package helpers

import (
    "github.com/go-playground/validator/v10"
    "path/to/your/models" // Ganti dengan package yang sesuai
)

// ValidatePhoto adalah fungsi untuk memvalidasi data foto sebelum penyimpanan
func ValidatePhoto(photo models.Photo) error {
    // Membuat instance validator
    validate := validator.New()

    // Validasi Title
    if err := validate.Var(photo.Title, "required"); err != nil {
        return err
    }

    // Validasi PhotoURL
    if err := validate.Var(photo.PhotoURL, "required,url"); err != nil {
        return err
    }

    // Validasi UserID (pastikan UserID adalah pengguna yang valid)
    if photo.UserID <= 0 {
        return ErrInvalidUserID
    }

    // Validasi lainnya sesuai kebutuhan Anda
    // Contoh: Validasi Caption

    return nil
}
