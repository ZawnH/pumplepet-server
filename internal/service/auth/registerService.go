package auth

import (
	"pumplepet-server/internal/model"
	"pumplepet-server/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(username, email, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: username,
		Email: email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
} 