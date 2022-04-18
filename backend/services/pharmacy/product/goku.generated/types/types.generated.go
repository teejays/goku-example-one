package product_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/lib/naam"
	"github.com/teejays/goku/generator/lib/panics"
)

// Product: <comments>
type Product struct {
	ID         scalars.ID    `json:"id"`
	MedicineID scalars.ID    `json:"medicine_id" validate:"required"`
	Mass       int           `json:"mass" validate:"required"`
	Count      int           `json:"count" validate:"required"`
	Name       *string       `json:"name"`
	CreatedAt  scalars.Time  `json:"created_at"`
	UpdatedAt  scalars.Time  `json:"updated_at"`
	DeletedAt  *scalars.Time `json:"deleted_at"`
}

func (t Product) GetID() scalars.ID {
	return t.ID
}
func (t Product) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t Product) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// ProductField enum: <insert comment>
type ProductField int

const (
	ProductField_INVALID    ProductField = 0
	ProductField_ID         ProductField = 1
	ProductField_MedicineID ProductField = 2
	ProductField_Mass       ProductField = 3
	ProductField_Count      ProductField = 4
	ProductField_Name       ProductField = 5
	ProductField_CreatedAt  ProductField = 6
	ProductField_UpdatedAt  ProductField = 7
	ProductField_DeletedAt  ProductField = 8
)

func NewProductFieldFromString(s string) ProductField {
	switch s {
	case "INVALID":
		return ProductField_INVALID
	case "ID":
		return ProductField_ID
	case "MedicineID":
		return ProductField_MedicineID
	case "Mass":
		return ProductField_Mass
	case "Count":
		return ProductField_Count
	case "Name":
		return ProductField_Name
	case "CreatedAt":
		return ProductField_CreatedAt
	case "UpdatedAt":
		return ProductField_UpdatedAt
	case "DeletedAt":
		return ProductField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "ProductField"))
	}
}

// String implements the `fmt.Stringer` interface for ProductField. It allows us to print the enum values as strings.
func (f ProductField) String() string {
	switch f {
	case ProductField_INVALID:
		return "INVALID"
	case ProductField_ID:
		return "ID"
	case ProductField_MedicineID:
		return "MedicineID"
	case ProductField_Mass:
		return "Mass"
	case ProductField_Count:
		return "Count"
	case ProductField_Name:
		return "Name"
	case ProductField_CreatedAt:
		return "CreatedAt"
	case ProductField_UpdatedAt:
		return "UpdatedAt"
	case ProductField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "ProductField"))
	}
}

// Name gives a naam representation of the enum value
func (f ProductField) Name() naam.Name {
	switch f {
	case ProductField_ID:
		return naam.New("id")
	case ProductField_MedicineID:
		return naam.New("medicine_id")
	case ProductField_Mass:
		return naam.New("mass")
	case ProductField_Count:
		return naam.New("count")
	case ProductField_Name:
		return naam.New("name")
	case ProductField_CreatedAt:
		return naam.New("created_at")
	case ProductField_UpdatedAt:
		return naam.New("updated_at")
	case ProductField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("ProductField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f ProductField) Value() (driver.Value, error) {
	switch f {
	case ProductField_INVALID:
		return nil, nil
	case ProductField_ID:
		return "ID", nil
	case ProductField_MedicineID:
		return "MedicineID", nil
	case ProductField_Mass:
		return "Mass", nil
	case ProductField_Count:
		return "Count", nil
	case ProductField_Name:
		return "Name", nil
	case ProductField_CreatedAt:
		return "CreatedAt", nil
	case ProductField_UpdatedAt:
		return "UpdatedAt", nil
	case ProductField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum ProductField to DB: '%d' is not a valid value for enum ProductField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *ProductField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewProductFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "ProductField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f ProductField) ImplementsGraphQLType(name string) bool {
	return name == "ProductField"
}

func (f *ProductField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewProductFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for ProductField: %T", input)
	}
	return err
}

func (f *ProductField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum ProductField to a string: %w", err)
	}
	i := NewProductFieldFromString(enumStr)
	*f = i
	return nil
}

func (f ProductField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil ProductField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type ProductFieldCondition struct {
	Op     filterlib.Operator
	Values []ProductField
}

func (c ProductFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c ProductFieldCondition) Len() int {
	return len(c.Values)
}
func (c ProductFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// ProductFilter: <comments>
type ProductFilter struct {
	ID         *filterlib.UUIDCondition      `json:"id"`
	MedicineID *filterlib.UUIDCondition      `json:"medicine_id"`
	Mass       *filterlib.IntCondition       `json:"mass"`
	Count      *filterlib.IntCondition       `json:"count"`
	Name       *filterlib.StringCondition    `json:"name"`
	CreatedAt  *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt  *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt  *filterlib.TimestampCondition `json:"deleted_at"`
	And        []ProductFilter               `json:"and"`
	Or         []ProductFilter               `json:"or"`
}

// UpdateProductRequest: <comments>
type UpdateProductRequest struct {
	Object        Product        `json:"object"`
	Fields        []ProductField `json:"fields"`
	ExcludeFields []ProductField `json:"exclude_fields"`
}

// UpdateProductResponse: <comments>
type UpdateProductResponse struct {
	Object Product `json:"object"`
}

// GetProductRequest: <comments>
type GetProductRequest struct {
	ID scalars.ID `json:"id"`
}

// ListProductRequest: <comments>
type ListProductRequest struct {
	Filter ProductFilter `json:"filter"`
}

// ListProductResponse: <comments>
type ListProductResponse struct {
	Items []Product `json:"items"`
	Count int       `json:"count"`
}

// QueryByTextProductRequest: <comments>
type QueryByTextProductRequest struct {
	QueryText string `json:"query_text"`
}
