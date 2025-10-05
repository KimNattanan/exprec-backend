package dto

import "github.com/KimNattanan/exprec-backend/internal/entities"

func ToUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		Email: user.Email,
		Name:  user.Name,
	}
}

func ToUserResponseList(users []*entities.User) []*UserResponse {
	result := make([]*UserResponse, len(users))
	for i, u := range users {
		result[i] = ToUserResponse(u)
	}
	return result
}

func ToUserEntity(req *RegisterRequest) *entities.User {
	return &entities.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
}
