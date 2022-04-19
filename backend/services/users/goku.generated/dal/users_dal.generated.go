package users

import (
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect

	dal_global "github.com/teejays/goku-example-one/backend/goku.generated/dal"
)

// UsersServiceDAL encapsulates DAL methods for types that fall under Users
type UsersServiceDAL struct {
	dal_global.ExampleAppAppDAL
}
