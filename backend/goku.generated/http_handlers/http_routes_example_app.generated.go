package http_handlers_example_app

import (
	gopi "github.com/teejays/gopi"

	pharmacy_http_handlers "github.com/teejays/goku-example-one/backend/services/pharmacy/goku.generated/http_handlers"
	users_http_handlers "github.com/teejays/goku-example-one/backend/services/users/goku.generated/http_handlers"
)

// GetRoutes returns all the routes for this application
func GetRoutes() []gopi.Route {
	var routes []gopi.Route
	routes = append(routes, pharmacy_http_handlers.GetRoutes()...)
	routes = append(routes, users_http_handlers.GetRoutes()...)
	return routes
}
