package auth

import (
	"pumplepet-server/internal/model"
	"pumplepet-server/pkg/database"
	"pumplepet-server/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginUser(email, password string) (*LoginResponse, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := checkPassword(user.Password, password); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Don't return the password
	user.Password = ""

	return &LoginResponse{
		User:  &user,
		Token: token,
	}, nil
}
