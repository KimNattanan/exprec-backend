package dto

import "github.com/KimNattanan/exprec-backend/internal/entities"

func ToUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

func ToUserResponseList(users []*entities.User) []*UserResponse {
	res := make([]*UserResponse, len(users))
	for i, u := range users {
		res[i] = ToUserResponse(u)
	}
	return res
}

func ToUserEntity(req *RegisterRequest) *entities.User {
	return &entities.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
}
