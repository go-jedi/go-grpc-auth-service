package auth

import (
	"github.com/go-jedi/auth-service/internal/service"
	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
