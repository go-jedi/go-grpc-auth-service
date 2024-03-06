package converter

import (
	"github.com/go-jedi/auth-service/internal/model"
	desc "github.com/go-jedi/auth-service/pkg/auth_v1"
)

func ToRegisterServiceFromProto(registerRequest *desc.RegisterRequest) *model.RegisterRequest {
	return &model.RegisterRequest{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}
}

func ToLoginServiceFromProto(loginRequest *desc.LoginRequest) *model.LoginRequest {
	return &model.LoginRequest{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}
}

func ToLoginProtoFromService(loginResponse *model.LoginResponse) *desc.LoginResponse {
	return &desc.LoginResponse{
		AccessToken:  loginResponse.AccessToken,
		RefreshToken: loginResponse.RefreshToken,
	}
}
