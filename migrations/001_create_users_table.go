package migrations

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Create enum type for roles
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE role AS ENUM ('admin', 'user');
		EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`).Error; err != nil {
		return err
	}

	// Create users table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Create admin user if not exists
	var adminCount int64
	if err := db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&adminCount).Error; err != nil {
		return err
	}

	if adminCount == 0 {
		admin := models.User{
			Name:     "Admin",
			Email:    "admin@example.com",
			Password: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // password: admin123
			Role:     models.RoleAdmin,
		}
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
	}

	return nil
} 