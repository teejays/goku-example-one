package methods_user

import (
	"context"
	"fmt"

	"github.com/teejays/clog"

	"github.com/teejays/goku/generator/external/client/db"

	user_dal "github.com/teejays/goku/example/backend/services/users/user/goku.generated/dal"
	user_types "github.com/teejays/goku/example/backend/services/users/user/goku.generated/types"
)

// UserEntity_Server provides all the methods that fall under this entity
type UserEntity_Server struct{}

func NewServer() *UserEntity_Server {
	return &UserEntity_Server{}
}

func (s UserEntity_Server) AddUser(ctx context.Context, req user_types.User) (user_types.User, error) {
	var resp user_types.User
	var err error

	clog.Infof("[Method] AddUser() starting with Request\n%+v", req)
	conn, err := db.NewConnection("users")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "users", err)
	}
	// Get the DAL wrapper
	d := user_dal.UserEntityDAL{}
	resp, err = d.AddUser(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s UserEntity_Server) UpdateUser(ctx context.Context, req user_types.UpdateUserRequest) (user_types.UpdateUserResponse, error) {
	var resp user_types.UpdateUserResponse
	var err error

	clog.Infof("[Method] UpdateUser() starting with Request\n%+v", req)
	conn, err := db.NewConnection("users")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "users", err)
	}
	// Get the DAL wrapper
	d := user_dal.UserEntityDAL{}
	resp, err = d.UpdateUser(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s UserEntity_Server) GetUser(ctx context.Context, req user_types.GetUserRequest) (user_types.User, error) {
	var resp user_types.User
	var err error

	clog.Infof("[Method] GetUser() starting with Request\n%+v", req)
	conn, err := db.NewConnection("users")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "users", err)
	}
	// Get the DAL wrapper
	d := user_dal.UserEntityDAL{}
	resp, err = d.GetUser(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s UserEntity_Server) ListUser(ctx context.Context, req user_types.ListUserRequest) (user_types.ListUserResponse, error) {
	var resp user_types.ListUserResponse
	var err error

	clog.Infof("[Method] ListUser() starting with Request\n%+v", req)
	conn, err := db.NewConnection("users")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "users", err)
	}
	// Get the DAL wrapper
	d := user_dal.UserEntityDAL{}
	resp, err = d.ListUser(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s UserEntity_Server) QueryByTextUser(ctx context.Context, req user_types.QueryByTextUserRequest) (user_types.ListUserResponse, error) {
	var resp user_types.ListUserResponse
	var err error

	clog.Infof("[Method] QueryByTextUser() starting with Request\n%+v", req)
	conn, err := db.NewConnection("users")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "users", err)
	}
	// Get the DAL wrapper
	d := user_dal.UserEntityDAL{}
	resp, err = d.QueryByTextUser(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}
