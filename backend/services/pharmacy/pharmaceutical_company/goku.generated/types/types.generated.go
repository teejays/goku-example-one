package pharmaceutical_company_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku-util/filter"
	"github.com/teejays/goku-util/naam"
	"github.com/teejays/goku-util/panics"
	"github.com/teejays/goku-util/scalars"
)

// PharmaceuticalCompany: <comments>
type PharmaceuticalCompany struct {
	ID        scalars.ID    `json:"id"`
	Name      string        `json:"name" validate:"required"`
	CreatedAt scalars.Time  `json:"created_at"`
	UpdatedAt scalars.Time  `json:"updated_at"`
	DeletedAt *scalars.Time `json:"deleted_at"`
}

func (t PharmaceuticalCompany) GetID() scalars.ID {
	return t.ID
}
func (t PharmaceuticalCompany) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t PharmaceuticalCompany) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// PharmaceuticalCompanyField enum: <insert comment>
type PharmaceuticalCompanyField int

const (
	PharmaceuticalCompanyField_INVALID   PharmaceuticalCompanyField = 0
	PharmaceuticalCompanyField_ID        PharmaceuticalCompanyField = 1
	PharmaceuticalCompanyField_Name      PharmaceuticalCompanyField = 2
	PharmaceuticalCompanyField_CreatedAt PharmaceuticalCompanyField = 3
	PharmaceuticalCompanyField_UpdatedAt PharmaceuticalCompanyField = 4
	PharmaceuticalCompanyField_DeletedAt PharmaceuticalCompanyField = 5
)

func NewPharmaceuticalCompanyFieldFromString(s string) PharmaceuticalCompanyField {
	switch s {
	case "INVALID":
		return PharmaceuticalCompanyField_INVALID
	case "ID":
		return PharmaceuticalCompanyField_ID
	case "Name":
		return PharmaceuticalCompanyField_Name
	case "CreatedAt":
		return PharmaceuticalCompanyField_CreatedAt
	case "UpdatedAt":
		return PharmaceuticalCompanyField_UpdatedAt
	case "DeletedAt":
		return PharmaceuticalCompanyField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "PharmaceuticalCompanyField"))
	}
}

// String implements the `fmt.Stringer` interface for PharmaceuticalCompanyField. It allows us to print the enum values as strings.
func (f PharmaceuticalCompanyField) String() string {
	switch f {
	case PharmaceuticalCompanyField_INVALID:
		return "INVALID"
	case PharmaceuticalCompanyField_ID:
		return "ID"
	case PharmaceuticalCompanyField_Name:
		return "Name"
	case PharmaceuticalCompanyField_CreatedAt:
		return "CreatedAt"
	case PharmaceuticalCompanyField_UpdatedAt:
		return "UpdatedAt"
	case PharmaceuticalCompanyField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "PharmaceuticalCompanyField"))
	}
}

// Name gives a naam representation of the enum value
func (f PharmaceuticalCompanyField) Name() naam.Name {
	switch f {
	case PharmaceuticalCompanyField_ID:
		return naam.New("id")
	case PharmaceuticalCompanyField_Name:
		return naam.New("name")
	case PharmaceuticalCompanyField_CreatedAt:
		return naam.New("created_at")
	case PharmaceuticalCompanyField_UpdatedAt:
		return naam.New("updated_at")
	case PharmaceuticalCompanyField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("PharmaceuticalCompanyField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f PharmaceuticalCompanyField) Value() (driver.Value, error) {
	switch f {
	case PharmaceuticalCompanyField_INVALID:
		return nil, nil
	case PharmaceuticalCompanyField_ID:
		return "ID", nil
	case PharmaceuticalCompanyField_Name:
		return "Name", nil
	case PharmaceuticalCompanyField_CreatedAt:
		return "CreatedAt", nil
	case PharmaceuticalCompanyField_UpdatedAt:
		return "UpdatedAt", nil
	case PharmaceuticalCompanyField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum PharmaceuticalCompanyField to DB: '%d' is not a valid value for enum PharmaceuticalCompanyField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *PharmaceuticalCompanyField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewPharmaceuticalCompanyFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "PharmaceuticalCompanyField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f PharmaceuticalCompanyField) ImplementsGraphQLType(name string) bool {
	return name == "PharmaceuticalCompanyField"
}

func (f *PharmaceuticalCompanyField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewPharmaceuticalCompanyFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for PharmaceuticalCompanyField: %T", input)
	}
	return err
}

func (f *PharmaceuticalCompanyField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum PharmaceuticalCompanyField to a string: %w", err)
	}
	i := NewPharmaceuticalCompanyFieldFromString(enumStr)
	*f = i
	return nil
}

func (f PharmaceuticalCompanyField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil PharmaceuticalCompanyField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type PharmaceuticalCompanyFieldCondition struct {
	Op     filterlib.Operator
	Values []PharmaceuticalCompanyField
}

func (c PharmaceuticalCompanyFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c PharmaceuticalCompanyFieldCondition) Len() int {
	return len(c.Values)
}
func (c PharmaceuticalCompanyFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// PharmaceuticalCompanyFilter: <comments>
type PharmaceuticalCompanyFilter struct {
	ID        *filterlib.UUIDCondition      `json:"id"`
	Name      *filterlib.StringCondition    `json:"name"`
	CreatedAt *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt *filterlib.TimestampCondition `json:"deleted_at"`
	And       []PharmaceuticalCompanyFilter `json:"and"`
	Or        []PharmaceuticalCompanyFilter `json:"or"`
}

// UpdatePharmaceuticalCompanyRequest: <comments>
type UpdatePharmaceuticalCompanyRequest struct {
	Object        PharmaceuticalCompany        `json:"object"`
	Fields        []PharmaceuticalCompanyField `json:"fields"`
	ExcludeFields []PharmaceuticalCompanyField `json:"exclude_fields"`
}

// UpdatePharmaceuticalCompanyResponse: <comments>
type UpdatePharmaceuticalCompanyResponse struct {
	Object PharmaceuticalCompany `json:"object"`
}

// GetPharmaceuticalCompanyRequest: <comments>
type GetPharmaceuticalCompanyRequest struct {
	ID scalars.ID `json:"id"`
}

// ListPharmaceuticalCompanyRequest: <comments>
type ListPharmaceuticalCompanyRequest struct {
	Filter PharmaceuticalCompanyFilter `json:"filter"`
}

// ListPharmaceuticalCompanyResponse: <comments>
type ListPharmaceuticalCompanyResponse struct {
	Items []PharmaceuticalCompany `json:"items"`
	Count int                     `json:"count"`
}

// QueryByTextPharmaceuticalCompanyRequest: <comments>
type QueryByTextPharmaceuticalCompanyRequest struct {
	QueryText string `json:"query_text"`
}
