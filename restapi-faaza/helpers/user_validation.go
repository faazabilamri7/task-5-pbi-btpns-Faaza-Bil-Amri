package helpers

import (
    "github.com/go-playground/validator/v10"
    "path/to/your/models" // Ganti dengan package yang sesuai
)

// ValidateUser adalah fungsi untuk memvalidasi data pengguna sebelum penyimpanan
func ValidateUser(user models.User) error {
    // Membuat instance validator
    validate := validator.New()

    // Validasi Email
    if err := validate.Var(user.Email, "required,email"); err != nil {
        return err
    }

    // Validasi Password
    if len(user.Password) < 6 {
        return ErrPasswordTooShort
    }

    // Validasi lainnya sesuai kebutuhan Anda
    // Contoh: Validasi Username

    return nil
}
