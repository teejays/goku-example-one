package example_app_types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/lib/naam"
	"github.com/teejays/goku/generator/lib/panics"
)

// Address: <comments>
type Address struct {
	Line1      string            `json:"line_1" validate:"required"`
	Line2      *string           `json:"line_2"`
	City       string            `json:"city" validate:"required"`
	Province   PakistaniProvince `json:"province" validate:"required"`
	PostalCode *string           `json:"postal_code"`
	Country    Country           `json:"country" validate:"required"`
}

// AddressWithMeta: <comments>
type AddressWithMeta struct {
	ParentID   scalars.ID        `json:"parent_id" json:"-"`
	ID         scalars.ID        `json:"id"`
	Line1      string            `json:"line_1" validate:"required"`
	Line2      *string           `json:"line_2"`
	City       string            `json:"city" validate:"required"`
	Province   PakistaniProvince `json:"province" validate:"required"`
	PostalCode *string           `json:"postal_code"`
	Country    Country           `json:"country" validate:"required"`
	CreatedAt  scalars.Time      `json:"created_at"`
	UpdatedAt  scalars.Time      `json:"updated_at"`
	DeletedAt  *scalars.Time     `json:"deleted_at"`
}

func (t AddressWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t AddressWithMeta) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t AddressWithMeta) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// Contact: <comments>
type Contact struct {
	Name    PersonName `json:"name" validate:"required"`
	Email   string     `json:"email" validate:"required"`
	Address Address    `json:"address" validate:"required"`
}

// ContactWithMeta: <comments>
type ContactWithMeta struct {
	ParentID  scalars.ID    `json:"parent_id" json:"-"`
	ID        scalars.ID    `json:"id"`
	Name      PersonName    `json:"name" validate:"required"`
	Email     string        `json:"email" validate:"required"`
	Address   Address       `json:"address" validate:"required"`
	CreatedAt scalars.Time  `json:"created_at"`
	UpdatedAt scalars.Time  `json:"updated_at"`
	DeletedAt *scalars.Time `json:"deleted_at"`
}

func (t ContactWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t ContactWithMeta) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t ContactWithMeta) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// PersonName: <comments>
type PersonName struct {
	First         string  `json:"first" validate:"required"`
	MiddleInitial *string `json:"middle_initial"`
	Last          string  `json:"last" validate:"required"`
}

// PersonNameWithMeta: <comments>
type PersonNameWithMeta struct {
	ParentID      scalars.ID    `json:"parent_id" json:"-"`
	ID            scalars.ID    `json:"id"`
	First         string        `json:"first" validate:"required"`
	MiddleInitial *string       `json:"middle_initial"`
	Last          string        `json:"last" validate:"required"`
	CreatedAt     scalars.Time  `json:"created_at"`
	UpdatedAt     scalars.Time  `json:"updated_at"`
	DeletedAt     *scalars.Time `json:"deleted_at"`
}

func (t PersonNameWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t PersonNameWithMeta) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t PersonNameWithMeta) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// PhoneNumber: <comments>
type PhoneNumber struct {
	CountryCode int     `json:"country_code" validate:"required"`
	Number      string  `json:"number" validate:"required"`
	Extension   *string `json:"extension"`
}

// PhoneNumberWithMeta: <comments>
type PhoneNumberWithMeta struct {
	ParentID    scalars.ID    `json:"parent_id" json:"-"`
	ID          scalars.ID    `json:"id"`
	CountryCode int           `json:"country_code" validate:"required"`
	Number      string        `json:"number" validate:"required"`
	Extension   *string       `json:"extension"`
	CreatedAt   scalars.Time  `json:"created_at"`
	UpdatedAt   scalars.Time  `json:"updated_at"`
	DeletedAt   *scalars.Time `json:"deleted_at"`
}

func (t PhoneNumberWithMeta) GetID() scalars.ID {
	return t.ID
}
func (t PhoneNumberWithMeta) GetUpdatedAt() scalars.Time {
	return t.UpdatedAt
}
func (t PhoneNumberWithMeta) SetUpdatedAt(tim scalars.Time) {
	t.UpdatedAt = tim
}

// Country enum: <insert comment>
type Country int

const (
	Country_INVALID  Country = 0
	Country_Pakistan Country = 1
	Country_USA      Country = 2
)

func NewCountryFromString(s string) Country {
	switch s {
	case "INVALID":
		return Country_INVALID
	case "Pakistan":
		return Country_Pakistan
	case "USA":
		return Country_USA

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "Country"))
	}
}

// String implements the `fmt.Stringer` interface for Country. It allows us to print the enum values as strings.
func (f Country) String() string {
	switch f {
	case Country_INVALID:
		return "INVALID"
	case Country_Pakistan:
		return "Pakistan"
	case Country_USA:
		return "USA"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "Country"))
	}
}

// Name gives a naam representation of the enum value
func (f Country) Name() naam.Name {
	switch f {
	case Country_Pakistan:
		return naam.New("pakistan")
	case Country_USA:
		return naam.New("usa")
	default:
		panics.P("Country.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f Country) Value() (driver.Value, error) {
	switch f {
	case Country_INVALID:
		return nil, nil
	case Country_Pakistan:
		return "Pakistan", nil
	case Country_USA:
		return "USA", nil

	default:
		return nil, fmt.Errorf("Cannot save enum Country to DB: '%d' is not a valid value for enum Country", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *Country) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewCountryFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "Country")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f Country) ImplementsGraphQLType(name string) bool {
	return name == "Country"
}

func (f *Country) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewCountryFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for Country: %T", input)
	}
	return err
}

func (f *Country) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum Country to a string: %w", err)
	}
	i := NewCountryFromString(enumStr)
	*f = i
	return nil
}

func (f Country) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil Country pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type CountryCondition struct {
	Op     filterlib.Operator
	Values []Country
}

func (c CountryCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c CountryCondition) Len() int {
	return len(c.Values)
}
func (c CountryCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// PakistaniProvince enum: <insert comment>
type PakistaniProvince int

const (
	PakistaniProvince_INVALID            PakistaniProvince = 0
	PakistaniProvince_Punjab             PakistaniProvince = 1
	PakistaniProvince_Sindh              PakistaniProvince = 2
	PakistaniProvince_KhyberPakhtoonkhwa PakistaniProvince = 3
	PakistaniProvince_Balochistan        PakistaniProvince = 4
	PakistaniProvince_GilgitBaltistan    PakistaniProvince = 5
	PakistaniProvince_AzadKashmir        PakistaniProvince = 6
)

func NewPakistaniProvinceFromString(s string) PakistaniProvince {
	switch s {
	case "INVALID":
		return PakistaniProvince_INVALID
	case "Punjab":
		return PakistaniProvince_Punjab
	case "Sindh":
		return PakistaniProvince_Sindh
	case "KhyberPakhtoonkhwa":
		return PakistaniProvince_KhyberPakhtoonkhwa
	case "Balochistan":
		return PakistaniProvince_Balochistan
	case "GilgitBaltistan":
		return PakistaniProvince_GilgitBaltistan
	case "AzadKashmir":
		return PakistaniProvince_AzadKashmir

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "PakistaniProvince"))
	}
}

// String implements the `fmt.Stringer` interface for PakistaniProvince. It allows us to print the enum values as strings.
func (f PakistaniProvince) String() string {
	switch f {
	case PakistaniProvince_INVALID:
		return "INVALID"
	case PakistaniProvince_Punjab:
		return "Punjab"
	case PakistaniProvince_Sindh:
		return "Sindh"
	case PakistaniProvince_KhyberPakhtoonkhwa:
		return "KhyberPakhtoonkhwa"
	case PakistaniProvince_Balochistan:
		return "Balochistan"
	case PakistaniProvince_GilgitBaltistan:
		return "GilgitBaltistan"
	case PakistaniProvince_AzadKashmir:
		return "AzadKashmir"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "PakistaniProvince"))
	}
}

// Name gives a naam representation of the enum value
func (f PakistaniProvince) Name() naam.Name {
	switch f {
	case PakistaniProvince_Punjab:
		return naam.New("punjab")
	case PakistaniProvince_Sindh:
		return naam.New("sindh")
	case PakistaniProvince_KhyberPakhtoonkhwa:
		return naam.New("khyber_pakhtoonkhwa")
	case PakistaniProvince_Balochistan:
		return naam.New("balochistan")
	case PakistaniProvince_GilgitBaltistan:
		return naam.New("gilgit_baltistan")
	case PakistaniProvince_AzadKashmir:
		return naam.New("azad_kashmir")
	default:
		panics.P("PakistaniProvince.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f PakistaniProvince) Value() (driver.Value, error) {
	switch f {
	case PakistaniProvince_INVALID:
		return nil, nil
	case PakistaniProvince_Punjab:
		return "Punjab", nil
	case PakistaniProvince_Sindh:
		return "Sindh", nil
	case PakistaniProvince_KhyberPakhtoonkhwa:
		return "KhyberPakhtoonkhwa", nil
	case PakistaniProvince_Balochistan:
		return "Balochistan", nil
	case PakistaniProvince_GilgitBaltistan:
		return "GilgitBaltistan", nil
	case PakistaniProvince_AzadKashmir:
		return "AzadKashmir", nil

	default:
		return nil, fmt.Errorf("Cannot save enum PakistaniProvince to DB: '%d' is not a valid value for enum PakistaniProvince", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *PakistaniProvince) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewPakistaniProvinceFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "PakistaniProvince")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f PakistaniProvince) ImplementsGraphQLType(name string) bool {
	return name == "PakistaniProvince"
}

func (f *PakistaniProvince) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewPakistaniProvinceFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for PakistaniProvince: %T", input)
	}
	return err
}

func (f *PakistaniProvince) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum PakistaniProvince to a string: %w", err)
	}
	i := NewPakistaniProvinceFromString(enumStr)
	*f = i
	return nil
}

func (f PakistaniProvince) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil PakistaniProvince pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type PakistaniProvinceCondition struct {
	Op     filterlib.Operator
	Values []PakistaniProvince
}

func (c PakistaniProvinceCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c PakistaniProvinceCondition) Len() int {
	return len(c.Values)
}
func (c PakistaniProvinceCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// AddressField enum: <insert comment>
type AddressField int

const (
	AddressField_INVALID    AddressField = 0
	AddressField_ParentID   AddressField = 1
	AddressField_ID         AddressField = 2
	AddressField_Line1      AddressField = 3
	AddressField_Line2      AddressField = 4
	AddressField_City       AddressField = 5
	AddressField_Province   AddressField = 6
	AddressField_PostalCode AddressField = 7
	AddressField_Country    AddressField = 8
	AddressField_CreatedAt  AddressField = 9
	AddressField_UpdatedAt  AddressField = 10
	AddressField_DeletedAt  AddressField = 11
)

func NewAddressFieldFromString(s string) AddressField {
	switch s {
	case "INVALID":
		return AddressField_INVALID
	case "ParentID":
		return AddressField_ParentID
	case "ID":
		return AddressField_ID
	case "Line1":
		return AddressField_Line1
	case "Line2":
		return AddressField_Line2
	case "City":
		return AddressField_City
	case "Province":
		return AddressField_Province
	case "PostalCode":
		return AddressField_PostalCode
	case "Country":
		return AddressField_Country
	case "CreatedAt":
		return AddressField_CreatedAt
	case "UpdatedAt":
		return AddressField_UpdatedAt
	case "DeletedAt":
		return AddressField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "AddressField"))
	}
}

// String implements the `fmt.Stringer` interface for AddressField. It allows us to print the enum values as strings.
func (f AddressField) String() string {
	switch f {
	case AddressField_INVALID:
		return "INVALID"
	case AddressField_ParentID:
		return "ParentID"
	case AddressField_ID:
		return "ID"
	case AddressField_Line1:
		return "Line1"
	case AddressField_Line2:
		return "Line2"
	case AddressField_City:
		return "City"
	case AddressField_Province:
		return "Province"
	case AddressField_PostalCode:
		return "PostalCode"
	case AddressField_Country:
		return "Country"
	case AddressField_CreatedAt:
		return "CreatedAt"
	case AddressField_UpdatedAt:
		return "UpdatedAt"
	case AddressField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "AddressField"))
	}
}

// Name gives a naam representation of the enum value
func (f AddressField) Name() naam.Name {
	switch f {
	case AddressField_ParentID:
		return naam.New("parent_id")
	case AddressField_ID:
		return naam.New("id")
	case AddressField_Line1:
		return naam.New("line_1")
	case AddressField_Line2:
		return naam.New("line_2")
	case AddressField_City:
		return naam.New("city")
	case AddressField_Province:
		return naam.New("province")
	case AddressField_PostalCode:
		return naam.New("postal_code")
	case AddressField_Country:
		return naam.New("country")
	case AddressField_CreatedAt:
		return naam.New("created_at")
	case AddressField_UpdatedAt:
		return naam.New("updated_at")
	case AddressField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("AddressField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f AddressField) Value() (driver.Value, error) {
	switch f {
	case AddressField_INVALID:
		return nil, nil
	case AddressField_ParentID:
		return "ParentID", nil
	case AddressField_ID:
		return "ID", nil
	case AddressField_Line1:
		return "Line1", nil
	case AddressField_Line2:
		return "Line2", nil
	case AddressField_City:
		return "City", nil
	case AddressField_Province:
		return "Province", nil
	case AddressField_PostalCode:
		return "PostalCode", nil
	case AddressField_Country:
		return "Country", nil
	case AddressField_CreatedAt:
		return "CreatedAt", nil
	case AddressField_UpdatedAt:
		return "UpdatedAt", nil
	case AddressField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum AddressField to DB: '%d' is not a valid value for enum AddressField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *AddressField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewAddressFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "AddressField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f AddressField) ImplementsGraphQLType(name string) bool {
	return name == "AddressField"
}

func (f *AddressField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewAddressFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for AddressField: %T", input)
	}
	return err
}

func (f *AddressField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum AddressField to a string: %w", err)
	}
	i := NewAddressFieldFromString(enumStr)
	*f = i
	return nil
}

func (f AddressField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil AddressField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type AddressFieldCondition struct {
	Op     filterlib.Operator
	Values []AddressField
}

func (c AddressFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c AddressFieldCondition) Len() int {
	return len(c.Values)
}
func (c AddressFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// ContactField enum: <insert comment>
type ContactField int

const (
	ContactField_INVALID            ContactField = 0
	ContactField_ParentID           ContactField = 1
	ContactField_ID                 ContactField = 2
	ContactField_Name               ContactField = 3
	ContactField_Name_First         ContactField = 4
	ContactField_Name_MiddleInitial ContactField = 5
	ContactField_Name_Last          ContactField = 6
	ContactField_Email              ContactField = 7
	ContactField_Address            ContactField = 8
	ContactField_Address_Line1      ContactField = 9
	ContactField_Address_Line2      ContactField = 10
	ContactField_Address_City       ContactField = 11
	ContactField_Address_Province   ContactField = 12
	ContactField_Address_PostalCode ContactField = 13
	ContactField_Address_Country    ContactField = 14
	ContactField_CreatedAt          ContactField = 15
	ContactField_UpdatedAt          ContactField = 16
	ContactField_DeletedAt          ContactField = 17
)

func NewContactFieldFromString(s string) ContactField {
	switch s {
	case "INVALID":
		return ContactField_INVALID
	case "ParentID":
		return ContactField_ParentID
	case "ID":
		return ContactField_ID
	case "Name":
		return ContactField_Name
	case "Name_First":
		return ContactField_Name_First
	case "Name_MiddleInitial":
		return ContactField_Name_MiddleInitial
	case "Name_Last":
		return ContactField_Name_Last
	case "Email":
		return ContactField_Email
	case "Address":
		return ContactField_Address
	case "Address_Line1":
		return ContactField_Address_Line1
	case "Address_Line2":
		return ContactField_Address_Line2
	case "Address_City":
		return ContactField_Address_City
	case "Address_Province":
		return ContactField_Address_Province
	case "Address_PostalCode":
		return ContactField_Address_PostalCode
	case "Address_Country":
		return ContactField_Address_Country
	case "CreatedAt":
		return ContactField_CreatedAt
	case "UpdatedAt":
		return ContactField_UpdatedAt
	case "DeletedAt":
		return ContactField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "ContactField"))
	}
}

// String implements the `fmt.Stringer` interface for ContactField. It allows us to print the enum values as strings.
func (f ContactField) String() string {
	switch f {
	case ContactField_INVALID:
		return "INVALID"
	case ContactField_ParentID:
		return "ParentID"
	case ContactField_ID:
		return "ID"
	case ContactField_Name:
		return "Name"
	case ContactField_Name_First:
		return "Name_First"
	case ContactField_Name_MiddleInitial:
		return "Name_MiddleInitial"
	case ContactField_Name_Last:
		return "Name_Last"
	case ContactField_Email:
		return "Email"
	case ContactField_Address:
		return "Address"
	case ContactField_Address_Line1:
		return "Address_Line1"
	case ContactField_Address_Line2:
		return "Address_Line2"
	case ContactField_Address_City:
		return "Address_City"
	case ContactField_Address_Province:
		return "Address_Province"
	case ContactField_Address_PostalCode:
		return "Address_PostalCode"
	case ContactField_Address_Country:
		return "Address_Country"
	case ContactField_CreatedAt:
		return "CreatedAt"
	case ContactField_UpdatedAt:
		return "UpdatedAt"
	case ContactField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "ContactField"))
	}
}

// Name gives a naam representation of the enum value
func (f ContactField) Name() naam.Name {
	switch f {
	case ContactField_ParentID:
		return naam.New("parent_id")
	case ContactField_ID:
		return naam.New("id")
	case ContactField_Name:
		return naam.New("name")
	case ContactField_Name_First:
		return naam.New("name___first")
	case ContactField_Name_MiddleInitial:
		return naam.New("name___middle_initial")
	case ContactField_Name_Last:
		return naam.New("name___last")
	case ContactField_Email:
		return naam.New("email")
	case ContactField_Address:
		return naam.New("address")
	case ContactField_Address_Line1:
		return naam.New("address___line_1")
	case ContactField_Address_Line2:
		return naam.New("address___line_2")
	case ContactField_Address_City:
		return naam.New("address___city")
	case ContactField_Address_Province:
		return naam.New("address___province")
	case ContactField_Address_PostalCode:
		return naam.New("address___postal_code")
	case ContactField_Address_Country:
		return naam.New("address___country")
	case ContactField_CreatedAt:
		return naam.New("created_at")
	case ContactField_UpdatedAt:
		return naam.New("updated_at")
	case ContactField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("ContactField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f ContactField) Value() (driver.Value, error) {
	switch f {
	case ContactField_INVALID:
		return nil, nil
	case ContactField_ParentID:
		return "ParentID", nil
	case ContactField_ID:
		return "ID", nil
	case ContactField_Name:
		return "Name", nil
	case ContactField_Name_First:
		return "Name_First", nil
	case ContactField_Name_MiddleInitial:
		return "Name_MiddleInitial", nil
	case ContactField_Name_Last:
		return "Name_Last", nil
	case ContactField_Email:
		return "Email", nil
	case ContactField_Address:
		return "Address", nil
	case ContactField_Address_Line1:
		return "Address_Line1", nil
	case ContactField_Address_Line2:
		return "Address_Line2", nil
	case ContactField_Address_City:
		return "Address_City", nil
	case ContactField_Address_Province:
		return "Address_Province", nil
	case ContactField_Address_PostalCode:
		return "Address_PostalCode", nil
	case ContactField_Address_Country:
		return "Address_Country", nil
	case ContactField_CreatedAt:
		return "CreatedAt", nil
	case ContactField_UpdatedAt:
		return "UpdatedAt", nil
	case ContactField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum ContactField to DB: '%d' is not a valid value for enum ContactField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *ContactField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewContactFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "ContactField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f ContactField) ImplementsGraphQLType(name string) bool {
	return name == "ContactField"
}

func (f *ContactField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewContactFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for ContactField: %T", input)
	}
	return err
}

func (f *ContactField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum ContactField to a string: %w", err)
	}
	i := NewContactFieldFromString(enumStr)
	*f = i
	return nil
}

func (f ContactField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil ContactField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type ContactFieldCondition struct {
	Op     filterlib.Operator
	Values []ContactField
}

func (c ContactFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c ContactFieldCondition) Len() int {
	return len(c.Values)
}
func (c ContactFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// PersonNameField enum: <insert comment>
type PersonNameField int

const (
	PersonNameField_INVALID       PersonNameField = 0
	PersonNameField_ParentID      PersonNameField = 1
	PersonNameField_ID            PersonNameField = 2
	PersonNameField_First         PersonNameField = 3
	PersonNameField_MiddleInitial PersonNameField = 4
	PersonNameField_Last          PersonNameField = 5
	PersonNameField_CreatedAt     PersonNameField = 6
	PersonNameField_UpdatedAt     PersonNameField = 7
	PersonNameField_DeletedAt     PersonNameField = 8
)

func NewPersonNameFieldFromString(s string) PersonNameField {
	switch s {
	case "INVALID":
		return PersonNameField_INVALID
	case "ParentID":
		return PersonNameField_ParentID
	case "ID":
		return PersonNameField_ID
	case "First":
		return PersonNameField_First
	case "MiddleInitial":
		return PersonNameField_MiddleInitial
	case "Last":
		return PersonNameField_Last
	case "CreatedAt":
		return PersonNameField_CreatedAt
	case "UpdatedAt":
		return PersonNameField_UpdatedAt
	case "DeletedAt":
		return PersonNameField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "PersonNameField"))
	}
}

// String implements the `fmt.Stringer` interface for PersonNameField. It allows us to print the enum values as strings.
func (f PersonNameField) String() string {
	switch f {
	case PersonNameField_INVALID:
		return "INVALID"
	case PersonNameField_ParentID:
		return "ParentID"
	case PersonNameField_ID:
		return "ID"
	case PersonNameField_First:
		return "First"
	case PersonNameField_MiddleInitial:
		return "MiddleInitial"
	case PersonNameField_Last:
		return "Last"
	case PersonNameField_CreatedAt:
		return "CreatedAt"
	case PersonNameField_UpdatedAt:
		return "UpdatedAt"
	case PersonNameField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "PersonNameField"))
	}
}

// Name gives a naam representation of the enum value
func (f PersonNameField) Name() naam.Name {
	switch f {
	case PersonNameField_ParentID:
		return naam.New("parent_id")
	case PersonNameField_ID:
		return naam.New("id")
	case PersonNameField_First:
		return naam.New("first")
	case PersonNameField_MiddleInitial:
		return naam.New("middle_initial")
	case PersonNameField_Last:
		return naam.New("last")
	case PersonNameField_CreatedAt:
		return naam.New("created_at")
	case PersonNameField_UpdatedAt:
		return naam.New("updated_at")
	case PersonNameField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("PersonNameField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f PersonNameField) Value() (driver.Value, error) {
	switch f {
	case PersonNameField_INVALID:
		return nil, nil
	case PersonNameField_ParentID:
		return "ParentID", nil
	case PersonNameField_ID:
		return "ID", nil
	case PersonNameField_First:
		return "First", nil
	case PersonNameField_MiddleInitial:
		return "MiddleInitial", nil
	case PersonNameField_Last:
		return "Last", nil
	case PersonNameField_CreatedAt:
		return "CreatedAt", nil
	case PersonNameField_UpdatedAt:
		return "UpdatedAt", nil
	case PersonNameField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum PersonNameField to DB: '%d' is not a valid value for enum PersonNameField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *PersonNameField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewPersonNameFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "PersonNameField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f PersonNameField) ImplementsGraphQLType(name string) bool {
	return name == "PersonNameField"
}

func (f *PersonNameField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewPersonNameFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for PersonNameField: %T", input)
	}
	return err
}

func (f *PersonNameField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum PersonNameField to a string: %w", err)
	}
	i := NewPersonNameFieldFromString(enumStr)
	*f = i
	return nil
}

func (f PersonNameField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil PersonNameField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type PersonNameFieldCondition struct {
	Op     filterlib.Operator
	Values []PersonNameField
}

func (c PersonNameFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c PersonNameFieldCondition) Len() int {
	return len(c.Values)
}
func (c PersonNameFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// PhoneNumberField enum: <insert comment>
type PhoneNumberField int

const (
	PhoneNumberField_INVALID     PhoneNumberField = 0
	PhoneNumberField_ParentID    PhoneNumberField = 1
	PhoneNumberField_ID          PhoneNumberField = 2
	PhoneNumberField_CountryCode PhoneNumberField = 3
	PhoneNumberField_Number      PhoneNumberField = 4
	PhoneNumberField_Extension   PhoneNumberField = 5
	PhoneNumberField_CreatedAt   PhoneNumberField = 6
	PhoneNumberField_UpdatedAt   PhoneNumberField = 7
	PhoneNumberField_DeletedAt   PhoneNumberField = 8
)

func NewPhoneNumberFieldFromString(s string) PhoneNumberField {
	switch s {
	case "INVALID":
		return PhoneNumberField_INVALID
	case "ParentID":
		return PhoneNumberField_ParentID
	case "ID":
		return PhoneNumberField_ID
	case "CountryCode":
		return PhoneNumberField_CountryCode
	case "Number":
		return PhoneNumberField_Number
	case "Extension":
		return PhoneNumberField_Extension
	case "CreatedAt":
		return PhoneNumberField_CreatedAt
	case "UpdatedAt":
		return PhoneNumberField_UpdatedAt
	case "DeletedAt":
		return PhoneNumberField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "PhoneNumberField"))
	}
}

// String implements the `fmt.Stringer` interface for PhoneNumberField. It allows us to print the enum values as strings.
func (f PhoneNumberField) String() string {
	switch f {
	case PhoneNumberField_INVALID:
		return "INVALID"
	case PhoneNumberField_ParentID:
		return "ParentID"
	case PhoneNumberField_ID:
		return "ID"
	case PhoneNumberField_CountryCode:
		return "CountryCode"
	case PhoneNumberField_Number:
		return "Number"
	case PhoneNumberField_Extension:
		return "Extension"
	case PhoneNumberField_CreatedAt:
		return "CreatedAt"
	case PhoneNumberField_UpdatedAt:
		return "UpdatedAt"
	case PhoneNumberField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "PhoneNumberField"))
	}
}

// Name gives a naam representation of the enum value
func (f PhoneNumberField) Name() naam.Name {
	switch f {
	case PhoneNumberField_ParentID:
		return naam.New("parent_id")
	case PhoneNumberField_ID:
		return naam.New("id")
	case PhoneNumberField_CountryCode:
		return naam.New("country_code")
	case PhoneNumberField_Number:
		return naam.New("number")
	case PhoneNumberField_Extension:
		return naam.New("extension")
	case PhoneNumberField_CreatedAt:
		return naam.New("created_at")
	case PhoneNumberField_UpdatedAt:
		return naam.New("updated_at")
	case PhoneNumberField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("PhoneNumberField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f PhoneNumberField) Value() (driver.Value, error) {
	switch f {
	case PhoneNumberField_INVALID:
		return nil, nil
	case PhoneNumberField_ParentID:
		return "ParentID", nil
	case PhoneNumberField_ID:
		return "ID", nil
	case PhoneNumberField_CountryCode:
		return "CountryCode", nil
	case PhoneNumberField_Number:
		return "Number", nil
	case PhoneNumberField_Extension:
		return "Extension", nil
	case PhoneNumberField_CreatedAt:
		return "CreatedAt", nil
	case PhoneNumberField_UpdatedAt:
		return "UpdatedAt", nil
	case PhoneNumberField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum PhoneNumberField to DB: '%d' is not a valid value for enum PhoneNumberField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *PhoneNumberField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewPhoneNumberFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "PhoneNumberField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f PhoneNumberField) ImplementsGraphQLType(name string) bool {
	return name == "PhoneNumberField"
}

func (f *PhoneNumberField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewPhoneNumberFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for PhoneNumberField: %T", input)
	}
	return err
}

func (f *PhoneNumberField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum PhoneNumberField to a string: %w", err)
	}
	i := NewPhoneNumberFieldFromString(enumStr)
	*f = i
	return nil
}

func (f PhoneNumberField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil PhoneNumberField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type PhoneNumberFieldCondition struct {
	Op     filterlib.Operator
	Values []PhoneNumberField
}

func (c PhoneNumberFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c PhoneNumberFieldCondition) Len() int {
	return len(c.Values)
}
func (c PhoneNumberFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// AddressFilter: <comments>
type AddressFilter struct {
	ParentID   *filterlib.UUIDCondition      `json:"parent_id"`
	ID         *filterlib.UUIDCondition      `json:"id"`
	Line1      *filterlib.StringCondition    `json:"line_1"`
	Line2      *filterlib.StringCondition    `json:"line_2"`
	City       *filterlib.StringCondition    `json:"city"`
	Province   *PakistaniProvinceCondition   `json:"province"`
	PostalCode *filterlib.StringCondition    `json:"postal_code"`
	Country    *CountryCondition             `json:"country"`
	CreatedAt  *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt  *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt  *filterlib.TimestampCondition `json:"deleted_at"`
	And        []AddressFilter               `json:"and"`
	Or         []AddressFilter               `json:"or"`
}

// ContactFilter: <comments>
type ContactFilter struct {
	ParentID  *filterlib.UUIDCondition      `json:"parent_id"`
	ID        *filterlib.UUIDCondition      `json:"id"`
	Name      *PersonNameFilter             `json:"name"`
	Email     *filterlib.StringCondition    `json:"email"`
	Address   *AddressFilter                `json:"address"`
	CreatedAt *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt *filterlib.TimestampCondition `json:"deleted_at"`
	And       []ContactFilter               `json:"and"`
	Or        []ContactFilter               `json:"or"`
}

// PersonNameFilter: <comments>
type PersonNameFilter struct {
	ParentID      *filterlib.UUIDCondition      `json:"parent_id"`
	ID            *filterlib.UUIDCondition      `json:"id"`
	First         *filterlib.StringCondition    `json:"first"`
	MiddleInitial *filterlib.StringCondition    `json:"middle_initial"`
	Last          *filterlib.StringCondition    `json:"last"`
	CreatedAt     *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt     *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt     *filterlib.TimestampCondition `json:"deleted_at"`
	And           []PersonNameFilter            `json:"and"`
	Or            []PersonNameFilter            `json:"or"`
}

// PhoneNumberFilter: <comments>
type PhoneNumberFilter struct {
	ParentID    *filterlib.UUIDCondition      `json:"parent_id"`
	ID          *filterlib.UUIDCondition      `json:"id"`
	CountryCode *filterlib.IntCondition       `json:"country_code"`
	Number      *filterlib.StringCondition    `json:"number"`
	Extension   *filterlib.StringCondition    `json:"extension"`
	CreatedAt   *filterlib.TimestampCondition `json:"created_at"`
	UpdatedAt   *filterlib.TimestampCondition `json:"updated_at"`
	DeletedAt   *filterlib.TimestampCondition `json:"deleted_at"`
	And         []PhoneNumberFilter           `json:"and"`
	Or          []PhoneNumberFilter           `json:"or"`
}
