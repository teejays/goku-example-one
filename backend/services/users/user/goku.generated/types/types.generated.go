package user_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/lib/naam"
	"github.com/teejays/goku/generator/lib/panics"

	example_app_types "github.com/teejays/goku/example/backend/goku.generated/types"
)

// User: <comments>
type User struct {
	ID           scalars.ID                     `json:"id"`
	Name         example_app_types.PersonName   `json:"name" validate:"required"`
	Email        string                         `json:"email"`
	PhoneNumber  *example_app_types.PhoneNumber `json:"phone_number"`
	PasswordHash string                         `json:"password_hash" validate:"required"`
	CreatedAt    scalars.Time                   `json:"created_at"`
	UpdatedAt    scalars.Time                   `json:"updated_at"`
	DeletedAt    *scalars.Time                  `json:"deleted_at"`
}

func (t User) GetID() scalars.ID {
	return t.ID
}
func (t User) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t User) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// UserField enum: <insert comment>
type UserField int

const (
	UserField_INVALID                 UserField = 0
	UserField_ID                      UserField = 1
	UserField_Name                    UserField = 2
	UserField_Name_First              UserField = 3
	UserField_Name_MiddleInitial      UserField = 4
	UserField_Name_Last               UserField = 5
	UserField_Email                   UserField = 6
	UserField_PhoneNumber             UserField = 7
	UserField_PhoneNumber_CountryCode UserField = 8
	UserField_PhoneNumber_Number      UserField = 9
	UserField_PhoneNumber_Extension   UserField = 10
	UserField_PasswordHash            UserField = 11
	UserField_CreatedAt               UserField = 12
	UserField_UpdatedAt               UserField = 13
	UserField_DeletedAt               UserField = 14
)

func NewUserFieldFromString(s string) UserField {
	switch s {
	case "INVALID":
		return UserField_INVALID
	case "ID":
		return UserField_ID
	case "Name":
		return UserField_Name
	case "Name_First":
		return UserField_Name_First
	case "Name_MiddleInitial":
		return UserField_Name_MiddleInitial
	case "Name_Last":
		return UserField_Name_Last
	case "Email":
		return UserField_Email
	case "PhoneNumber":
		return UserField_PhoneNumber
	case "PhoneNumber_CountryCode":
		return UserField_PhoneNumber_CountryCode
	case "PhoneNumber_Number":
		return UserField_PhoneNumber_Number
	case "PhoneNumber_Extension":
		return UserField_PhoneNumber_Extension
	case "PasswordHash":
		return UserField_PasswordHash
	case "CreatedAt":
		return UserField_CreatedAt
	case "UpdatedAt":
		return UserField_UpdatedAt
	case "DeletedAt":
		return UserField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "UserField"))
	}
}

// String implements the `fmt.Stringer` interface for UserField. It allows us to print the enum values as strings.
func (f UserField) String() string {
	switch f {
	case UserField_INVALID:
		return "INVALID"
	case UserField_ID:
		return "ID"
	case UserField_Name:
		return "Name"
	case UserField_Name_First:
		return "Name_First"
	case UserField_Name_MiddleInitial:
		return "Name_MiddleInitial"
	case UserField_Name_Last:
		return "Name_Last"
	case UserField_Email:
		return "Email"
	case UserField_PhoneNumber:
		return "PhoneNumber"
	case UserField_PhoneNumber_CountryCode:
		return "PhoneNumber_CountryCode"
	case UserField_PhoneNumber_Number:
		return "PhoneNumber_Number"
	case UserField_PhoneNumber_Extension:
		return "PhoneNumber_Extension"
	case UserField_PasswordHash:
		return "PasswordHash"
	case UserField_CreatedAt:
		return "CreatedAt"
	case UserField_UpdatedAt:
		return "UpdatedAt"
	case UserField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "UserField"))
	}
}

// Name gives a naam representation of the enum value
func (f UserField) Name() naam.Name {
	switch f {
	case UserField_ID:
		return naam.New("id")
	case UserField_Name:
		return naam.New("name")
	case UserField_Name_First:
		return naam.New("name___first")
	case UserField_Name_MiddleInitial:
		return naam.New("name___middle_initial")
	case UserField_Name_Last:
		return naam.New("name___last")
	case UserField_Email:
		return naam.New("email")
	case UserField_PhoneNumber:
		return naam.New("phone_number")
	case UserField_PhoneNumber_CountryCode:
		return naam.New("phone_number___country_code")
	case UserField_PhoneNumber_Number:
		return naam.New("phone_number___number")
	case UserField_PhoneNumber_Extension:
		return naam.New("phone_number___extension")
	case UserField_PasswordHash:
		return naam.New("password_hash")
	case UserField_CreatedAt:
		return naam.New("created_at")
	case UserField_UpdatedAt:
		return naam.New("updated_at")
	case UserField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("UserField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f UserField) Value() (driver.Value, error) {
	switch f {
	case UserField_INVALID:
		return nil, nil
	case UserField_ID:
		return "ID", nil
	case UserField_Name:
		return "Name", nil
	case UserField_Name_First:
		return "Name_First", nil
	case UserField_Name_MiddleInitial:
		return "Name_MiddleInitial", nil
	case UserField_Name_Last:
		return "Name_Last", nil
	case UserField_Email:
		return "Email", nil
	case UserField_PhoneNumber:
		return "PhoneNumber", nil
	case UserField_PhoneNumber_CountryCode:
		return "PhoneNumber_CountryCode", nil
	case UserField_PhoneNumber_Number:
		return "PhoneNumber_Number", nil
	case UserField_PhoneNumber_Extension:
		return "PhoneNumber_Extension", nil
	case UserField_PasswordHash:
		return "PasswordHash", nil
	case UserField_CreatedAt:
		return "CreatedAt", nil
	case UserField_UpdatedAt:
		return "UpdatedAt", nil
	case UserField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum UserField to DB: '%d' is not a valid value for enum UserField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *UserField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewUserFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "UserField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f UserField) ImplementsGraphQLType(name string) bool {
	return name == "UserField"
}

func (f *UserField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewUserFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for UserField: %T", input)
	}
	return err
}

func (f *UserField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum UserField to a string: %w", err)
	}
	i := NewUserFieldFromString(enumStr)
	*f = i
	return nil
}

func (f UserField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil UserField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type UserFieldCondition struct {
	Op     filterlib.Operator
	Values []UserField
}

func (c UserFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c UserFieldCondition) Len() int {
	return len(c.Values)
}
func (c UserFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// UserFilter: <comments>
type UserFilter struct {
	ID           *filterlib.UUIDCondition             `json:"id"`
	Name         *example_app_types.PersonNameFilter  `json:"name"`
	Email        *filterlib.StringCondition           `json:"email"`
	PhoneNumber  *example_app_types.PhoneNumberFilter `json:"phone_number"`
	PasswordHash *filterlib.StringCondition           `json:"password_hash"`
	CreatedAt    *filterlib.TimestampCondition        `json:"created_at"`
	UpdatedAt    *filterlib.TimestampCondition        `json:"updated_at"`
	DeletedAt    *filterlib.TimestampCondition        `json:"deleted_at"`
	And          []UserFilter                         `json:"and"`
	Or           []UserFilter                         `json:"or"`
}

// UpdateUserRequest: <comments>
type UpdateUserRequest struct {
	Object        User        `json:"object"`
	Fields        []UserField `json:"fields"`
	ExcludeFields []UserField `json:"exclude_fields"`
}

// UpdateUserResponse: <comments>
type UpdateUserResponse struct {
	Object User `json:"object"`
}

// GetUserRequest: <comments>
type GetUserRequest struct {
	ID scalars.ID `json:"id"`
}

// ListUserRequest: <comments>
type ListUserRequest struct {
	Filter UserFilter `json:"filter"`
}

// ListUserResponse: <comments>
type ListUserResponse struct {
	Items []User `json:"items"`
	Count int    `json:"count"`
}

// QueryByTextUserRequest: <comments>
type QueryByTextUserRequest struct {
	QueryText string `json:"query_text"`
}
