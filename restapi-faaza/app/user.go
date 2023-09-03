package app

import (
    "gorm.io/gorm"
)

// GetUserByID mengembalikan pengguna berdasarkan ID
func GetUserByID(db *gorm.DB, userID uint) (*User, error) {
    var user User
    if err := db.First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// CreateUser membuat pengguna baru dalam basis data
func CreateUser(db *gorm.DB, user *User) error {
    if err := db.Create(user).Error; err != nil {
        return err
    }
    return nil
}

// UpdateUser memperbarui data pengguna dalam basis data
func UpdateUser(db *gorm.DB, user *User) error {
    if err := db.Save(user).Error; err != nil {
        return err
    }
    return nil
}

// DeleteUser menghapus pengguna dari basis data
func DeleteUser(db *gorm.DB, userID uint) error {
    if err := db.Delete(&User{}, userID).Error; err != nil {
        return err
    }
    return nil
}
