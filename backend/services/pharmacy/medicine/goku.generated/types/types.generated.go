package medicine_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/lib/naam"
	"github.com/teejays/goku/generator/lib/panics"
)

// Medicine: <comments>
type Medicine struct {
	ID                scalars.ID           `json:"id"`
	Name              string               `json:"name" validate:"required"`
	CompanyID         scalars.ID           `json:"company_id" validate:"required"`
	PrimaryIngredient Ingredient           `json:"primary_ingredient" validate:"required"`
	Ingredients       []IngredientWithMeta `json:"ingredients" validate:"required"`
	ModeOfDelivery    ModeOfDelivery       `json:"mode_of_delivery" validate:"required"`
	CreatedAt         scalars.Time         `json:"created_at"`
	UpdatedAt         scalars.Time         `json:"updated_at"`
	DeletedAt         *scalars.Time        `json:"deleted_at"`
}

func (t Medicine) GetID() scalars.ID {
	return t.ID
}
func (t Medicine) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t Medicine) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// Ingredient: <comments>
type Ingredient struct {
	DrugID              scalars.ID `json:"drug_id"`
	IsPrimaryIngredient bool       `json:"is_primary_ingredient"`
}

// IngredientWithMeta: <comments>
type IngredientWithMeta struct {
	ParentID            scalars.ID    `json:"parent_id" json:"-"`
	ID                  scalars.ID    `json:"id"`
	DrugID              scalars.ID    `json:"drug_id"`
	IsPrimaryIngredient bool          `json:"is_primary_ingredient"`
	CreatedAt           scalars.Time  `json:"created_at"`
	UpdatedAt           scalars.Time  `json:"updated_at"`
	DeletedAt           *scalars.Time `json:"deleted_at"`
}

func (t IngredientWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t IngredientWithMeta) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t IngredientWithMeta) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// ModeOfDelivery enum: <insert comment>
type ModeOfDelivery int

const (
	ModeOfDelivery_INVALID   ModeOfDelivery = 0
	ModeOfDelivery_Tablet    ModeOfDelivery = 1
	ModeOfDelivery_Syrup     ModeOfDelivery = 2
	ModeOfDelivery_Capsule   ModeOfDelivery = 3
	ModeOfDelivery_Injection ModeOfDelivery = 4
)

func NewModeOfDeliveryFromString(s string) ModeOfDelivery {
	switch s {
	case "INVALID":
		return ModeOfDelivery_INVALID
	case "Tablet":
		return ModeOfDelivery_Tablet
	case "Syrup":
		return ModeOfDelivery_Syrup
	case "Capsule":
		return ModeOfDelivery_Capsule
	case "Injection":
		return ModeOfDelivery_Injection

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "ModeOfDelivery"))
	}
}

// String implements the `fmt.Stringer` interface for ModeOfDelivery. It allows us to print the enum values as strings.
func (f ModeOfDelivery) String() string {
	switch f {
	case ModeOfDelivery_INVALID:
		return "INVALID"
	case ModeOfDelivery_Tablet:
		return "Tablet"
	case ModeOfDelivery_Syrup:
		return "Syrup"
	case ModeOfDelivery_Capsule:
		return "Capsule"
	case ModeOfDelivery_Injection:
		return "Injection"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "ModeOfDelivery"))
	}
}

// Name gives a naam representation of the enum value
func (f ModeOfDelivery) Name() naam.Name {
	switch f {
	case ModeOfDelivery_Tablet:
		return naam.New("tablet")
	case ModeOfDelivery_Syrup:
		return naam.New("syrup")
	case ModeOfDelivery_Capsule:
		return naam.New("capsule")
	case ModeOfDelivery_Injection:
		return naam.New("injection")
	default:
		panics.P("ModeOfDelivery.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f ModeOfDelivery) Value() (driver.Value, error) {
	switch f {
	case ModeOfDelivery_INVALID:
		return nil, nil
	case ModeOfDelivery_Tablet:
		return "Tablet", nil
	case ModeOfDelivery_Syrup:
		return "Syrup", nil
	case ModeOfDelivery_Capsule:
		return "Capsule", nil
	case ModeOfDelivery_Injection:
		return "Injection", nil

	default:
		return nil, fmt.Errorf("Cannot save enum ModeOfDelivery to DB: '%d' is not a valid value for enum ModeOfDelivery", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *ModeOfDelivery) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewModeOfDeliveryFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "ModeOfDelivery")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f ModeOfDelivery) ImplementsGraphQLType(name string) bool {
	return name == "ModeOfDelivery"
}

func (f *ModeOfDelivery) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewModeOfDeliveryFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for ModeOfDelivery: %T", input)
	}
	return err
}

func (f *ModeOfDelivery) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum ModeOfDelivery to a string: %w", err)
	}
	i := NewModeOfDeliveryFromString(enumStr)
	*f = i
	return nil
}

func (f ModeOfDelivery) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil ModeOfDelivery pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type ModeOfDeliveryCondition struct {
	Op     filterlib.Operator
	Values []ModeOfDelivery
}

func (c ModeOfDeliveryCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c ModeOfDeliveryCondition) Len() int {
	return len(c.Values)
}
func (c ModeOfDeliveryCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// IngredientField enum: <insert comment>
type IngredientField int

const (
	IngredientField_INVALID             IngredientField = 0
	IngredientField_ParentID            IngredientField = 1
	IngredientField_ID                  IngredientField = 2
	IngredientField_DrugID              IngredientField = 3
	IngredientField_IsPrimaryIngredient IngredientField = 4
	IngredientField_CreatedAt           IngredientField = 5
	IngredientField_UpdatedAt           IngredientField = 6
	IngredientField_DeletedAt           IngredientField = 7
)

func NewIngredientFieldFromString(s string) IngredientField {
	switch s {
	case "INVALID":
		return IngredientField_INVALID
	case "ParentID":
		return IngredientField_ParentID
	case "ID":
		return IngredientField_ID
	case "DrugID":
		return IngredientField_DrugID
	case "IsPrimaryIngredient":
		return IngredientField_IsPrimaryIngredient
	case "CreatedAt":
		return IngredientField_CreatedAt
	case "UpdatedAt":
		return IngredientField_UpdatedAt
	case "DeletedAt":
		return IngredientField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "IngredientField"))
	}
}

// String implements the `fmt.Stringer` interface for IngredientField. It allows us to print the enum values as strings.
func (f IngredientField) String() string {
	switch f {
	case IngredientField_INVALID:
		return "INVALID"
	case IngredientField_ParentID:
		return "ParentID"
	case IngredientField_ID:
		return "ID"
	case IngredientField_DrugID:
		return "DrugID"
	case IngredientField_IsPrimaryIngredient:
		return "IsPrimaryIngredient"
	case IngredientField_CreatedAt:
		return "CreatedAt"
	case IngredientField_UpdatedAt:
		return "UpdatedAt"
	case IngredientField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "IngredientField"))
	}
}

// Name gives a naam representation of the enum value
func (f IngredientField) Name() naam.Name {
	switch f {
	case IngredientField_ParentID:
		return naam.New("parent_id")
	case IngredientField_ID:
		return naam.New("id")
	case IngredientField_DrugID:
		return naam.New("drug_id")
	case IngredientField_IsPrimaryIngredient:
		return naam.New("is_primary_ingredient")
	case IngredientField_CreatedAt:
		return naam.New("created_at")
	case IngredientField_UpdatedAt:
		return naam.New("updated_at")
	case IngredientField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("IngredientField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f IngredientField) Value() (driver.Value, error) {
	switch f {
	case IngredientField_INVALID:
		return nil, nil
	case IngredientField_ParentID:
		return "ParentID", nil
	case IngredientField_ID:
		return "ID", nil
	case IngredientField_DrugID:
		return "DrugID", nil
	case IngredientField_IsPrimaryIngredient:
		return "IsPrimaryIngredient", nil
	case IngredientField_CreatedAt:
		return "CreatedAt", nil
	case IngredientField_UpdatedAt:
		return "UpdatedAt", nil
	case IngredientField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum IngredientField to DB: '%d' is not a valid value for enum IngredientField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *IngredientField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewIngredientFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "IngredientField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f IngredientField) ImplementsGraphQLType(name string) bool {
	return name == "IngredientField"
}

func (f *IngredientField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewIngredientFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for IngredientField: %T", input)
	}
	return err
}

func (f *IngredientField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum IngredientField to a string: %w", err)
	}
	i := NewIngredientFieldFromString(enumStr)
	*f = i
	return nil
}

func (f IngredientField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil IngredientField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type IngredientFieldCondition struct {
	Op     filterlib.Operator
	Values []IngredientField
}

func (c IngredientFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c IngredientFieldCondition) Len() int {
	return len(c.Values)
}
func (c IngredientFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// MedicineField enum: <insert comment>
type MedicineField int

const (
	MedicineField_INVALID                               MedicineField = 0
	MedicineField_ID                                    MedicineField = 1
	MedicineField_Name                                  MedicineField = 2
	MedicineField_CompanyID                             MedicineField = 3
	MedicineField_PrimaryIngredient                     MedicineField = 4
	MedicineField_PrimaryIngredient_DrugID              MedicineField = 5
	MedicineField_PrimaryIngredient_IsPrimaryIngredient MedicineField = 6
	MedicineField_Ingredients                           MedicineField = 7
	MedicineField_Ingredients_ParentID                  MedicineField = 8
	MedicineField_Ingredients_ID                        MedicineField = 9
	MedicineField_Ingredients_DrugID                    MedicineField = 10
	MedicineField_Ingredients_IsPrimaryIngredient       MedicineField = 11
	MedicineField_Ingredients_CreatedAt                 MedicineField = 12
	MedicineField_Ingredients_UpdatedAt                 MedicineField = 13
	MedicineField_Ingredients_DeletedAt                 MedicineField = 14
	MedicineField_ModeOfDelivery                        MedicineField = 15
	MedicineField_CreatedAt                             MedicineField = 16
	MedicineField_UpdatedAt                             MedicineField = 17
	MedicineField_DeletedAt                             MedicineField = 18
)

func NewMedicineFieldFromString(s string) MedicineField {
	switch s {
	case "INVALID":
		return MedicineField_INVALID
	case "ID":
		return MedicineField_ID
	case "Name":
		return MedicineField_Name
	case "CompanyID":
		return MedicineField_CompanyID
	case "PrimaryIngredient":
		return MedicineField_PrimaryIngredient
	case "PrimaryIngredient_DrugID":
		return MedicineField_PrimaryIngredient_DrugID
	case "PrimaryIngredient_IsPrimaryIngredient":
		return MedicineField_PrimaryIngredient_IsPrimaryIngredient
	case "Ingredients":
		return MedicineField_Ingredients
	case "Ingredients_ParentID":
		return MedicineField_Ingredients_ParentID
	case "Ingredients_ID":
		return MedicineField_Ingredients_ID
	case "Ingredients_DrugID":
		return MedicineField_Ingredients_DrugID
	case "Ingredients_IsPrimaryIngredient":
		return MedicineField_Ingredients_IsPrimaryIngredient
	case "Ingredients_CreatedAt":
		return MedicineField_Ingredients_CreatedAt
	case "Ingredients_UpdatedAt":
		return MedicineField_Ingredients_UpdatedAt
	case "Ingredients_DeletedAt":
		return MedicineField_Ingredients_DeletedAt
	case "ModeOfDelivery":
		return MedicineField_ModeOfDelivery
	case "CreatedAt":
		return MedicineField_CreatedAt
	case "UpdatedAt":
		return MedicineField_UpdatedAt
	case "DeletedAt":
		return MedicineField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "MedicineField"))
	}
}

// String implements the `fmt.Stringer` interface for MedicineField. It allows us to print the enum values as strings.
func (f MedicineField) String() string {
	switch f {
	case MedicineField_INVALID:
		return "INVALID"
	case MedicineField_ID:
		return "ID"
	case MedicineField_Name:
		return "Name"
	case MedicineField_CompanyID:
		return "CompanyID"
	case MedicineField_PrimaryIngredient:
		return "PrimaryIngredient"
	case MedicineField_PrimaryIngredient_DrugID:
		return "PrimaryIngredient_DrugID"
	case MedicineField_PrimaryIngredient_IsPrimaryIngredient:
		return "PrimaryIngredient_IsPrimaryIngredient"
	case MedicineField_Ingredients:
		return "Ingredients"
	case MedicineField_Ingredients_ParentID:
		return "Ingredients_ParentID"
	case MedicineField_Ingredients_ID:
		return "Ingredients_ID"
	case MedicineField_Ingredients_DrugID:
		return "Ingredients_DrugID"
	case MedicineField_Ingredients_IsPrimaryIngredient:
		return "Ingredients_IsPrimaryIngredient"
	case MedicineField_Ingredients_CreatedAt:
		return "Ingredients_CreatedAt"
	case MedicineField_Ingredients_UpdatedAt:
		return "Ingredients_UpdatedAt"
	case MedicineField_Ingredients_DeletedAt:
		return "Ingredients_DeletedAt"
	case MedicineField_ModeOfDelivery:
		return "ModeOfDelivery"
	case MedicineField_CreatedAt:
		return "CreatedAt"
	case MedicineField_UpdatedAt:
		return "UpdatedAt"
	case MedicineField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "MedicineField"))
	}
}

// Name gives a naam representation of the enum value
func (f MedicineField) Name() naam.Name {
	switch f {
	case MedicineField_ID:
		return naam.New("id")
	case MedicineField_Name:
		return naam.New("name")
	case MedicineField_CompanyID:
		return naam.New("company_id")
	case MedicineField_PrimaryIngredient:
		return naam.New("primary_ingredient")
	case MedicineField_PrimaryIngredient_DrugID:
		return naam.New("primary_ingredient___drug_id")
	case MedicineField_PrimaryIngredient_IsPrimaryIngredient:
		return naam.New("primary_ingredient___is_primary_ingredient")
	case MedicineField_Ingredients:
		return naam.New("ingredients")
	case MedicineField_Ingredients_ParentID:
		return naam.New("ingredients___parent_id")
	case MedicineField_Ingredients_ID:
		return naam.New("ingredients___id")
	case MedicineField_Ingredients_DrugID:
		return naam.New("ingredients___drug_id")
	case MedicineField_Ingredients_IsPrimaryIngredient:
		return naam.New("ingredients___is_primary_ingredient")
	case MedicineField_Ingredients_CreatedAt:
		return naam.New("ingredients___created_at")
	case MedicineField_Ingredients_UpdatedAt:
		return naam.New("ingredients___updated_at")
	case MedicineField_Ingredients_DeletedAt:
		return naam.New("ingredients___deleted_at")
	case MedicineField_ModeOfDelivery:
		return naam.New("mode_of_delivery")
	case MedicineField_CreatedAt:
		return naam.New("created_at")
	case MedicineField_UpdatedAt:
		return naam.New("updated_at")
	case MedicineField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("MedicineField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f MedicineField) Value() (driver.Value, error) {
	switch f {
	case MedicineField_INVALID:
		return nil, nil
	case MedicineField_ID:
		return "ID", nil
	case MedicineField_Name:
		return "Name", nil
	case MedicineField_CompanyID:
		return "CompanyID", nil
	case MedicineField_PrimaryIngredient:
		return "PrimaryIngredient", nil
	case MedicineField_PrimaryIngredient_DrugID:
		return "PrimaryIngredient_DrugID", nil
	case MedicineField_PrimaryIngredient_IsPrimaryIngredient:
		return "PrimaryIngredient_IsPrimaryIngredient", nil
	case MedicineField_Ingredients:
		return "Ingredients", nil
	case MedicineField_Ingredients_ParentID:
		return "Ingredients_ParentID", nil
	case MedicineField_Ingredients_ID:
		return "Ingredients_ID", nil
	case MedicineField_Ingredients_DrugID:
		return "Ingredients_DrugID", nil
	case MedicineField_Ingredients_IsPrimaryIngredient:
		return "Ingredients_IsPrimaryIngredient", nil
	case MedicineField_Ingredients_CreatedAt:
		return "Ingredients_CreatedAt", nil
	case MedicineField_Ingredients_UpdatedAt:
		return "Ingredients_UpdatedAt", nil
	case MedicineField_Ingredients_DeletedAt:
		return "Ingredients_DeletedAt", nil
	case MedicineField_ModeOfDelivery:
		return "ModeOfDelivery", nil
	case MedicineField_CreatedAt:
		return "CreatedAt", nil
	case MedicineField_UpdatedAt:
		return "UpdatedAt", nil
	case MedicineField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum MedicineField to DB: '%d' is not a valid value for enum MedicineField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *MedicineField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewMedicineFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "MedicineField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f MedicineField) ImplementsGraphQLType(name string) bool {
	return name == "MedicineField"
}

func (f *MedicineField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewMedicineFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for MedicineField: %T", input)
	}
	return err
}

func (f *MedicineField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum MedicineField to a string: %w", err)
	}
	i := NewMedicineFieldFromString(enumStr)
	*f = i
	return nil
}

func (f MedicineField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil MedicineField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type MedicineFieldCondition struct {
	Op     filterlib.Operator
	Values []MedicineField
}

func (c MedicineFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c MedicineFieldCondition) Len() int {
	return len(c.Values)
}
func (c MedicineFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// IngredientFilter: <comments>
type IngredientFilter struct {
	ParentID            *filterlib.UUIDCondition      `json:"parent_id"`
	ID                  *filterlib.UUIDCondition      `json:"id"`
	DrugID              *filterlib.UUIDCondition      `json:"drug_id"`
	IsPrimaryIngredient *filterlib.BoolCondition      `json:"is_primary_ingredient"`
	CreatedAt           *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt           *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt           *filterlib.TimestampCondition `json:"deleted_at"`
	And                 []IngredientFilter            `json:"and"`
	Or                  []IngredientFilter            `json:"or"`
}

// MedicineFilter: <comments>
type MedicineFilter struct {
	ID                *filterlib.UUIDCondition      `json:"id"`
	Name              *filterlib.StringCondition    `json:"name"`
	CompanyID         *filterlib.UUIDCondition      `json:"company_id"`
	PrimaryIngredient *IngredientFilter             `json:"primary_ingredient"`
	HavingIngredients *IngredientFilter             `json:"having_ingredients"`
	ModeOfDelivery    *ModeOfDeliveryCondition      `json:"mode_of_delivery"`
	CreatedAt         *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt         *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt         *filterlib.TimestampCondition `json:"deleted_at"`
	And               []MedicineFilter              `json:"and"`
	Or                []MedicineFilter              `json:"or"`
}

// UpdateMedicineRequest: <comments>
type UpdateMedicineRequest struct {
	Object        Medicine        `json:"object"`
	Fields        []MedicineField `json:"fields"`
	ExcludeFields []MedicineField `json:"exclude_fields"`
}

// UpdateMedicineResponse: <comments>
type UpdateMedicineResponse struct {
	Object Medicine `json:"object"`
}

// GetMedicineRequest: <comments>
type GetMedicineRequest struct {
	ID scalars.ID `json:"id"`
}

// ListMedicineRequest: <comments>
type ListMedicineRequest struct {
	Filter MedicineFilter `json:"filter"`
}

// ListMedicineResponse: <comments>
type ListMedicineResponse struct {
	Items []Medicine `json:"items"`
	Count int        `json:"count"`
}

// QueryByTextMedicineRequest: <comments>
type QueryByTextMedicineRequest struct {
	QueryText string `json:"query_text"`
}
