package drug_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/lib/naam"
	"github.com/teejays/goku/generator/lib/panics"
)

// Drug: <comments>
type Drug struct {
	ID        scalars.ID    `json:"id"`
	Name      string        `json:"name" validate:"required"`
	CreatedAt scalars.Time  `json:"created_at"`
	UpdatedAt scalars.Time  `json:"updated_at"`
	DeletedAt *scalars.Time `json:"deleted_at"`
}

func (t Drug) GetID() scalars.ID {
	return t.ID
}
func (t Drug) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t Drug) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// DrugField enum: <insert comment>
type DrugField int

const (
	DrugField_INVALID   DrugField = 0
	DrugField_ID        DrugField = 1
	DrugField_Name      DrugField = 2
	DrugField_CreatedAt DrugField = 3
	DrugField_UpdatedAt DrugField = 4
	DrugField_DeletedAt DrugField = 5
)

func NewDrugFieldFromString(s string) DrugField {
	switch s {
	case "INVALID":
		return DrugField_INVALID
	case "ID":
		return DrugField_ID
	case "Name":
		return DrugField_Name
	case "CreatedAt":
		return DrugField_CreatedAt
	case "UpdatedAt":
		return DrugField_UpdatedAt
	case "DeletedAt":
		return DrugField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "DrugField"))
	}
}

// String implements the `fmt.Stringer` interface for DrugField. It allows us to print the enum values as strings.
func (f DrugField) String() string {
	switch f {
	case DrugField_INVALID:
		return "INVALID"
	case DrugField_ID:
		return "ID"
	case DrugField_Name:
		return "Name"
	case DrugField_CreatedAt:
		return "CreatedAt"
	case DrugField_UpdatedAt:
		return "UpdatedAt"
	case DrugField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "DrugField"))
	}
}

// Name gives a naam representation of the enum value
func (f DrugField) Name() naam.Name {
	switch f {
	case DrugField_ID:
		return naam.New("id")
	case DrugField_Name:
		return naam.New("name")
	case DrugField_CreatedAt:
		return naam.New("created_at")
	case DrugField_UpdatedAt:
		return naam.New("updated_at")
	case DrugField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("DrugField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f DrugField) Value() (driver.Value, error) {
	switch f {
	case DrugField_INVALID:
		return nil, nil
	case DrugField_ID:
		return "ID", nil
	case DrugField_Name:
		return "Name", nil
	case DrugField_CreatedAt:
		return "CreatedAt", nil
	case DrugField_UpdatedAt:
		return "UpdatedAt", nil
	case DrugField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum DrugField to DB: '%d' is not a valid value for enum DrugField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *DrugField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewDrugFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "DrugField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f DrugField) ImplementsGraphQLType(name string) bool {
	return name == "DrugField"
}

func (f *DrugField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewDrugFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for DrugField: %T", input)
	}
	return err
}

func (f *DrugField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum DrugField to a string: %w", err)
	}
	i := NewDrugFieldFromString(enumStr)
	*f = i
	return nil
}

func (f DrugField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil DrugField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type DrugFieldCondition struct {
	Op     filterlib.Operator
	Values []DrugField
}

func (c DrugFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c DrugFieldCondition) Len() int {
	return len(c.Values)
}
func (c DrugFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// DrugFilter: <comments>
type DrugFilter struct {
	ID        *filterlib.UUIDCondition      `json:"id"`
	Name      *filterlib.StringCondition    `json:"name"`
	CreatedAt *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt *filterlib.TimestampCondition `json:"deleted_at"`
	And       []DrugFilter                  `json:"and"`
	Or        []DrugFilter                  `json:"or"`
}

// UpdateDrugRequest: <comments>
type UpdateDrugRequest struct {
	Object        Drug        `json:"object"`
	Fields        []DrugField `json:"fields"`
	ExcludeFields []DrugField `json:"exclude_fields"`
}

// UpdateDrugResponse: <comments>
type UpdateDrugResponse struct {
	Object Drug `json:"object"`
}

// GetDrugRequest: <comments>
type GetDrugRequest struct {
	ID scalars.ID `json:"id"`
}

// ListDrugRequest: <comments>
type ListDrugRequest struct {
	Filter DrugFilter `json:"filter"`
}

// ListDrugResponse: <comments>
type ListDrugResponse struct {
	Items []Drug `json:"items"`
	Count int    `json:"count"`
}

// QueryByTextDrugRequest: <comments>
type QueryByTextDrugRequest struct {
	QueryText string `json:"query_text"`
}
