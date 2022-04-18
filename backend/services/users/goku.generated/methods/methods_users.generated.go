package methods_users

import (
	"context"

	"github.com/teejays/clog"

	users_types "github.com/teejays/goku/example/backend/services/users/goku.generated/types"
	custom_methods "github.com/teejays/goku/example/backend/services/users/methods"
	user_methods "github.com/teejays/goku/example/backend/services/users/user/goku.generated/methods"
)

// UsersService_Server provides all the service methods, including methods from its entities.
type UsersService_Server struct {
	*user_methods.UserEntity_Server
}

func NewServer() *UsersService_Server {
	return &UsersService_Server{
		UserEntity_Server: user_methods.NewServer(),
	}
}

func (s UsersService_Server) AuthenticateUser(ctx context.Context, req users_types.AuthenticateRequest) (users_types.AuthenticateResponse, error) {
	var resp users_types.AuthenticateResponse
	var err error

	clog.Infof("[Method] AuthenticateUser() starting with Request\n%+v", req)
	resp, err = custom_methods.AuthenticateUser(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func (s UsersService_Server) RegisterUser(ctx context.Context, req users_types.RegisterUserRequest) (users_types.AuthenticateResponse, error) {
	var resp users_types.AuthenticateResponse
	var err error

	clog.Infof("[Method] RegisterUser() starting with Request\n%+v", req)
	resp, err = custom_methods.RegisterUser(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}
