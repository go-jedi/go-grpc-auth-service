package user

import (
	"github.com/go-jedi/auth-service/internal/service"
	desc "github.com/go-jedi/auth-service/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
