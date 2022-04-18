package users_types

import (
	example_app_types "github.com/teejays/goku/example/backend/goku.generated/types"
)

// AuthenticateRequest: <comments>
type AuthenticateRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateResponse: <comments>
type AuthenticateResponse struct {
	Token string `json:"token" validate:"required"`
}

// RegisterUserRequest: <comments>
type RegisterUserRequest struct {
	Email       string                        `json:"email" validate:"required"`
	Name        example_app_types.PersonName  `json:"name" validate:"required"`
	PhoneNumber example_app_types.PhoneNumber `json:"phone_number" validate:"required"`
	Password    string                        `json:"password" validate:"required"`
}
