package converter

import (
	"github.com/go-jedi/auth-service/internal/model"
	modelRepo "github.com/go-jedi/auth-service/internal/repository/auth/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:                   user.ID,
		Username:             user.Username,
		Password:             user.Password,
		CreatedAt:            user.CreatedAt,
		UpdatedAt:            user.UpdatedAt,
		PasswordLastChangeAt: user.PasswordLastChangeAt,
	}
}
