package http_handlers_users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teejays/clog"
	gopi "github.com/teejays/gopi"

	"github.com/teejays/goku-util/errutil"
	"github.com/teejays/goku-util/httputil"

	users_methods "github.com/teejays/goku-example-one/backend/services/users/goku.generated/methods"
	users_types "github.com/teejays/goku-example-one/backend/services/users/goku.generated/types"
	user_methods "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/methods"
	user_types "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/types"
)

// GetUsersRoutes returns all the routes for this namespace
func GetUsersRoutes() []gopi.Route {

	routes := []gopi.Route{
		{
			// API Route for POST users/register
			Method:       "POST",
			Version:      1,
			Path:         "users/register",
			HandlerFunc:  RegisterUserHandler,
			Authenticate: false,
		},
		{
			// API Route for POST users/authenticate
			Method:       "POST",
			Version:      1,
			Path:         "users/authenticate",
			HandlerFunc:  AuthenticateUserHandler,
			Authenticate: true,
		},
		{
			// API Route for POST users/user
			Method:       "POST",
			Version:      1,
			Path:         "users/user",
			HandlerFunc:  AddUserHandler,
			Authenticate: true,
		},
		{
			// API Route for PUT users/user
			Method:       "PUT",
			Version:      1,
			Path:         "users/user",
			HandlerFunc:  UpdateUserHandler,
			Authenticate: true,
		},
		{
			// API Route for GET users/user
			Method:       "GET",
			Version:      1,
			Path:         "users/user",
			HandlerFunc:  GetUserHandler,
			Authenticate: true,
		},
		{
			// API Route for GET users/user/list
			Method:       "GET",
			Version:      1,
			Path:         "users/user/list",
			HandlerFunc:  ListUserHandler,
			Authenticate: true,
		},
		{
			// API Route for GET users/user/query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "users/user/query_by_text",
			HandlerFunc:  QueryByTextUserHandler,
			Authenticate: true,
		},
	}

	return routes
}

// RegisterUserHandler is the HTTP handler for the method RegisterUser.
// The method's description: Create a new user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] RegisterUserHandler starting...")
	// Get the req from HTTP body
	var req users_types.RegisterUserRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := users_methods.NewServer()

	resp, err := s.RegisterUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// AuthenticateUserHandler is the HTTP handler for the method AuthenticateUser.
// The method's description: Handle authentication of users
func AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AuthenticateUserHandler starting...")
	// Get the req from HTTP body
	var req users_types.AuthenticateRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := users_methods.NewServer()

	resp, err := s.AuthenticateUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// AddUserHandler is the HTTP handler for the method AddUser.
// The method's description: Adds a new User entity
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AddUserHandler starting...")
	// Get the req from HTTP body
	var req user_types.User
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.AddUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// UpdateUserHandler is the HTTP handler for the method UpdateUser.
// The method's description: Adds a new User entity
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] UpdateUserHandler starting...")
	// Get the req from HTTP body
	var req user_types.UpdateUserRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.UpdateUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// GetUserHandler is the HTTP handler for the method GetUser.
// The method's description: Get a User entity
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] GetUserHandler starting...")
	// Get the req from the HTTP req body
	var req user_types.GetUserRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.GetUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// ListUserHandler is the HTTP handler for the method ListUser.
// The method's description: List User entities
func ListUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] ListUserHandler starting...")
	// Get the req from the HTTP req body
	var req user_types.ListUserRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.ListUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// QueryByTextUserHandler is the HTTP handler for the method QueryByTextUser.
// The method's description: List Users entities by free text search
func QueryByTextUserHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] QueryByTextUserHandler starting...")
	// Get the req from the HTTP req body
	var req user_types.QueryByTextUserRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.QueryByTextUser(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}
