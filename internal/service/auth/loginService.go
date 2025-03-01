package auth

import (
	"pumplepet-server/internal/model"
	"pumplepet-server/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginUser(email, password string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := checkPassword(user.Password, password); err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}