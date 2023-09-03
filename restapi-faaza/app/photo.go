package app

import (
    "gorm.io/gorm"
)

// GetPhotoByID mengembalikan foto berdasarkan ID
func GetPhotoByID(db *gorm.DB, photoID uint) (*Photo, error) {
    var photo Photo
    if err := db.First(&photo, photoID).Error; err != nil {
        return nil, err
    }
    return &photo, nil
}

// CreatePhoto membuat foto baru dalam basis data
func CreatePhoto(db *gorm.DB, photo *Photo) error {
    if err := db.Create(photo).Error; err != nil {
        return err
    }
    return nil
}

// UpdatePhoto memperbarui data foto dalam basis data
func UpdatePhoto(db *gorm.DB, photo *Photo) error {
    if err := db.Save(photo).Error; err != nil {
        return err
    }
    return nil
}

// DeletePhoto menghapus foto dari basis data
func DeletePhoto(db *gorm.DB, photoID uint) error {
    if err := db.Delete(&Photo{}, photoID).Error; err != nil {
        return err
    }
    return nil
}
