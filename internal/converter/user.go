package converter

import (
	"github.com/go-jedi/auth-service/internal/model"
	desc "github.com/go-jedi/auth-service/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToGetProtoFromService(user *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:                   user.ID,
		Username:             user.Username,
		Password:             user.Password,
		CreatedAt:            timestamppb.New(user.CreatedAt),
		UpdatedAt:            updatedAt,
		PasswordLastChangeAt: timestamppb.New(user.PasswordLastChangeAt),
	}
}

func ToUpdateNameServiceFromProto(updateNameRequest *desc.UpdateNameRequest) *model.UpdateNameRequest {
	return &model.UpdateNameRequest{
		ID:       updateNameRequest.Id,
		Username: updateNameRequest.Username,
	}
}

func ToUpdatePasswordServiceFromProto(updatePasswordRequest *desc.UpdatePasswordRequest) *model.UpdatePasswordRequest {
	return &model.UpdatePasswordRequest{
		ID:       updatePasswordRequest.Id,
		Password: updatePasswordRequest.Password,
	}
}
