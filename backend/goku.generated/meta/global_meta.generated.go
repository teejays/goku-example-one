package global_types_meta

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/teejays/clog"
	"github.com/teejays/goku/generator/external/client/db"
	"github.com/teejays/goku/generator/external/dalutil"
	"github.com/teejays/goku/generator/external/scalars"
	"github.com/teejays/goku/generator/external/types"
	"github.com/teejays/goku/generator/lib/naam"

	example_app_types "github.com/teejays/goku/example/backend/goku.generated/types"
)

type AddressMeta struct {
	types.BasicTypeMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField]
}

func GetAddressMeta() AddressMeta {
	objMeta := types.BasicTypeMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField]{
		Name: naam.New("address"),
		Fields: []example_app_types.AddressField{
			example_app_types.AddressField_ParentID,
			example_app_types.AddressField_ID,
			example_app_types.AddressField_Line1,
			example_app_types.AddressField_Line2,
			example_app_types.AddressField_City,
			example_app_types.AddressField_Province,
			example_app_types.AddressField_PostalCode,
			example_app_types.AddressField_Country,
			example_app_types.AddressField_CreatedAt,
			example_app_types.AddressField_UpdatedAt,
			example_app_types.AddressField_DeletedAt,
		},
	}
	return AddressMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta AddressMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField] {
	return meta.BasicTypeMetaBase
}

func (meta AddressMeta) SetMetaFieldValues(obj example_app_types.AddressWithMeta, now time.Time) example_app_types.AddressWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Addresses: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Addresses: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta AddressMeta) ConvertTimestampColumnsToUTC(obj example_app_types.AddressWithMeta) example_app_types.AddressWithMeta {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta AddressMeta) SetDefaultFieldValues(obj example_app_types.AddressWithMeta) example_app_types.AddressWithMeta {
	return obj
}

type AddressDALMeta struct {
	AddressMeta
	dalutil.BasicTypeDALMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField]
}

func GetAddressDALMeta() AddressDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField]{
		DatabaseColumnFields: []example_app_types.AddressField{
			example_app_types.AddressField_ParentID,
			example_app_types.AddressField_ID,
			example_app_types.AddressField_Line1,
			example_app_types.AddressField_Line2,
			example_app_types.AddressField_City,
			example_app_types.AddressField_Province,
			example_app_types.AddressField_PostalCode,
			example_app_types.AddressField_Country,
			example_app_types.AddressField_CreatedAt,
			example_app_types.AddressField_UpdatedAt,
			example_app_types.AddressField_DeletedAt,
		},
		DatabaseSubTableFields: []example_app_types.AddressField{},
		MutableOnlyByDALFields: []example_app_types.AddressField{
			example_app_types.AddressField_UpdatedAt,
			example_app_types.AddressField_DeletedAt,
		},
		NonMutableFields: []example_app_types.AddressField{
			example_app_types.AddressField_ParentID,
			example_app_types.AddressField_ID,
			example_app_types.AddressField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []example_app_types.AddressField{
			example_app_types.AddressField_CreatedAt,
			example_app_types.AddressField_UpdatedAt,
			example_app_types.AddressField_DeletedAt,
		},
		UpdatedAtField: example_app_types.AddressField_UpdatedAt,
	}
	return AddressDALMeta{
		AddressMeta:          GetAddressMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta AddressDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[example_app_types.AddressWithMeta, example_app_types.AddressField] {
	return meta.BasicTypeDALMetaBase
}

func (meta AddressDALMeta) GetDirectDBValues(obj example_app_types.AddressWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.Line1,
		obj.Line2,
		obj.City,
		obj.Province,
		obj.PostalCode,
		obj.Country,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta AddressDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj example_app_types.AddressWithMeta) (example_app_types.AddressWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta AddressDALMeta) ScanDBNextRow(rows *sql.Rows) (example_app_types.AddressWithMeta, error) {
	var elem example_app_types.AddressWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.Line1,
		&elem.Line2,
		&elem.City,
		&elem.Province,
		&elem.PostalCode,
		&elem.Country,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta AddressDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []example_app_types.AddressWithMeta) ([]example_app_types.AddressWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta AddressDALMeta) GetChangedFieldsAndValues(old, new example_app_types.AddressWithMeta, allowedFields []example_app_types.AddressField) ([]example_app_types.AddressField, []interface{}) {

	var colsWithValueChange []example_app_types.AddressField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(example_app_types.AddressField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_Line1, allowedFields) {
		if old.Line1 != new.Line1 {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_Line1)
			vals = append(vals, new.Line1)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_Line2, allowedFields) {
		if old.Line2 != new.Line2 {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_Line2)
			vals = append(vals, new.Line2)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_City, allowedFields) {
		if old.City != new.City {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_City)
			vals = append(vals, new.City)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_Province, allowedFields) {
		if old.Province != new.Province {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_Province)
			vals = append(vals, new.Province)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_PostalCode, allowedFields) {
		if old.PostalCode != new.PostalCode {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_PostalCode)
			vals = append(vals, new.PostalCode)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_Country, allowedFields) {
		if old.Country != new.Country {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_Country)
			vals = append(vals, new.Country)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.AddressField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.AddressField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta AddressDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[example_app_types.AddressWithMeta, example_app_types.AddressField], allowedFields []example_app_types.AddressField, elem example_app_types.AddressWithMeta) (example_app_types.AddressWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type ContactMeta struct {
	types.BasicTypeMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField]
}

func GetContactMeta() ContactMeta {
	objMeta := types.BasicTypeMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField]{
		Name: naam.New("contact"),
		Fields: []example_app_types.ContactField{
			example_app_types.ContactField_ParentID,
			example_app_types.ContactField_ID,
			example_app_types.ContactField_Name,
			example_app_types.ContactField_Name_First,
			example_app_types.ContactField_Name_MiddleInitial,
			example_app_types.ContactField_Name_Last,
			example_app_types.ContactField_Email,
			example_app_types.ContactField_Address,
			example_app_types.ContactField_Address_Line1,
			example_app_types.ContactField_Address_Line2,
			example_app_types.ContactField_Address_City,
			example_app_types.ContactField_Address_Province,
			example_app_types.ContactField_Address_PostalCode,
			example_app_types.ContactField_Address_Country,
			example_app_types.ContactField_CreatedAt,
			example_app_types.ContactField_UpdatedAt,
			example_app_types.ContactField_DeletedAt,
		},
	}
	return ContactMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta ContactMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField] {
	return meta.BasicTypeMetaBase
}

func (meta ContactMeta) SetMetaFieldValues(obj example_app_types.ContactWithMeta, now time.Time) example_app_types.ContactWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Contacts: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Contacts: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta ContactMeta) ConvertTimestampColumnsToUTC(obj example_app_types.ContactWithMeta) example_app_types.ContactWithMeta {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta ContactMeta) SetDefaultFieldValues(obj example_app_types.ContactWithMeta) example_app_types.ContactWithMeta {
	return obj
}

type ContactDALMeta struct {
	ContactMeta
	dalutil.BasicTypeDALMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField]
}

func GetContactDALMeta() ContactDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField]{
		DatabaseColumnFields: []example_app_types.ContactField{
			example_app_types.ContactField_ParentID,
			example_app_types.ContactField_ID,
			example_app_types.ContactField_Name_First,
			example_app_types.ContactField_Name_MiddleInitial,
			example_app_types.ContactField_Name_Last,
			example_app_types.ContactField_Email,
			example_app_types.ContactField_Address_Line1,
			example_app_types.ContactField_Address_Line2,
			example_app_types.ContactField_Address_City,
			example_app_types.ContactField_Address_Province,
			example_app_types.ContactField_Address_PostalCode,
			example_app_types.ContactField_Address_Country,
			example_app_types.ContactField_CreatedAt,
			example_app_types.ContactField_UpdatedAt,
			example_app_types.ContactField_DeletedAt,
		},
		DatabaseSubTableFields: []example_app_types.ContactField{},
		MutableOnlyByDALFields: []example_app_types.ContactField{
			example_app_types.ContactField_UpdatedAt,
			example_app_types.ContactField_DeletedAt,
		},
		NonMutableFields: []example_app_types.ContactField{
			example_app_types.ContactField_ParentID,
			example_app_types.ContactField_ID,
			example_app_types.ContactField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []example_app_types.ContactField{
			example_app_types.ContactField_CreatedAt,
			example_app_types.ContactField_UpdatedAt,
			example_app_types.ContactField_DeletedAt,
		},
		UpdatedAtField: example_app_types.ContactField_UpdatedAt,
	}
	return ContactDALMeta{
		ContactMeta:          GetContactMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta ContactDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[example_app_types.ContactWithMeta, example_app_types.ContactField] {
	return meta.BasicTypeDALMetaBase
}

func (meta ContactDALMeta) GetDirectDBValues(obj example_app_types.ContactWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.Name.First,
		obj.Name.MiddleInitial,
		obj.Name.Last,
		obj.Email,
		obj.Address.Line1,
		obj.Address.Line2,
		obj.Address.City,
		obj.Address.Province,
		obj.Address.PostalCode,
		obj.Address.Country,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta ContactDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj example_app_types.ContactWithMeta) (example_app_types.ContactWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta ContactDALMeta) ScanDBNextRow(rows *sql.Rows) (example_app_types.ContactWithMeta, error) {
	var elem example_app_types.ContactWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.Name.First,
		&elem.Name.MiddleInitial,
		&elem.Name.Last,
		&elem.Email,
		&elem.Address.Line1,
		&elem.Address.Line2,
		&elem.Address.City,
		&elem.Address.Province,
		&elem.Address.PostalCode,
		&elem.Address.Country,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta ContactDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []example_app_types.ContactWithMeta) ([]example_app_types.ContactWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta ContactDALMeta) GetChangedFieldsAndValues(old, new example_app_types.ContactWithMeta, allowedFields []example_app_types.ContactField) ([]example_app_types.ContactField, []interface{}) {

	var colsWithValueChange []example_app_types.ContactField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(example_app_types.ContactField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Name_First, allowedFields) {
		if old.Name.First != new.Name.First {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Name_First)
			vals = append(vals, new.Name.First)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Name_MiddleInitial, allowedFields) {
		if old.Name.MiddleInitial != new.Name.MiddleInitial {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Name_MiddleInitial)
			vals = append(vals, new.Name.MiddleInitial)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Name_Last, allowedFields) {
		if old.Name.Last != new.Name.Last {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Name_Last)
			vals = append(vals, new.Name.Last)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Email, allowedFields) {
		if old.Email != new.Email {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Email)
			vals = append(vals, new.Email)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_Line1, allowedFields) {
		if old.Address.Line1 != new.Address.Line1 {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_Line1)
			vals = append(vals, new.Address.Line1)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_Line2, allowedFields) {
		if old.Address.Line2 != new.Address.Line2 {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_Line2)
			vals = append(vals, new.Address.Line2)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_City, allowedFields) {
		if old.Address.City != new.Address.City {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_City)
			vals = append(vals, new.Address.City)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_Province, allowedFields) {
		if old.Address.Province != new.Address.Province {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_Province)
			vals = append(vals, new.Address.Province)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_PostalCode, allowedFields) {
		if old.Address.PostalCode != new.Address.PostalCode {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_PostalCode)
			vals = append(vals, new.Address.PostalCode)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_Address_Country, allowedFields) {
		if old.Address.Country != new.Address.Country {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_Address_Country)
			vals = append(vals, new.Address.Country)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.ContactField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.ContactField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta ContactDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[example_app_types.ContactWithMeta, example_app_types.ContactField], allowedFields []example_app_types.ContactField, elem example_app_types.ContactWithMeta) (example_app_types.ContactWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type PersonNameMeta struct {
	types.BasicTypeMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField]
}

func GetPersonNameMeta() PersonNameMeta {
	objMeta := types.BasicTypeMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField]{
		Name: naam.New("person_name"),
		Fields: []example_app_types.PersonNameField{
			example_app_types.PersonNameField_ParentID,
			example_app_types.PersonNameField_ID,
			example_app_types.PersonNameField_First,
			example_app_types.PersonNameField_MiddleInitial,
			example_app_types.PersonNameField_Last,
			example_app_types.PersonNameField_CreatedAt,
			example_app_types.PersonNameField_UpdatedAt,
			example_app_types.PersonNameField_DeletedAt,
		},
	}
	return PersonNameMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta PersonNameMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField] {
	return meta.BasicTypeMetaBase
}

func (meta PersonNameMeta) SetMetaFieldValues(obj example_app_types.PersonNameWithMeta, now time.Time) example_app_types.PersonNameWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting PersonNames: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting PersonNames: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta PersonNameMeta) ConvertTimestampColumnsToUTC(obj example_app_types.PersonNameWithMeta) example_app_types.PersonNameWithMeta {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta PersonNameMeta) SetDefaultFieldValues(obj example_app_types.PersonNameWithMeta) example_app_types.PersonNameWithMeta {
	return obj
}

type PersonNameDALMeta struct {
	PersonNameMeta
	dalutil.BasicTypeDALMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField]
}

func GetPersonNameDALMeta() PersonNameDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField]{
		DatabaseColumnFields: []example_app_types.PersonNameField{
			example_app_types.PersonNameField_ParentID,
			example_app_types.PersonNameField_ID,
			example_app_types.PersonNameField_First,
			example_app_types.PersonNameField_MiddleInitial,
			example_app_types.PersonNameField_Last,
			example_app_types.PersonNameField_CreatedAt,
			example_app_types.PersonNameField_UpdatedAt,
			example_app_types.PersonNameField_DeletedAt,
		},
		DatabaseSubTableFields: []example_app_types.PersonNameField{},
		MutableOnlyByDALFields: []example_app_types.PersonNameField{
			example_app_types.PersonNameField_UpdatedAt,
			example_app_types.PersonNameField_DeletedAt,
		},
		NonMutableFields: []example_app_types.PersonNameField{
			example_app_types.PersonNameField_ParentID,
			example_app_types.PersonNameField_ID,
			example_app_types.PersonNameField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []example_app_types.PersonNameField{
			example_app_types.PersonNameField_CreatedAt,
			example_app_types.PersonNameField_UpdatedAt,
			example_app_types.PersonNameField_DeletedAt,
		},
		UpdatedAtField: example_app_types.PersonNameField_UpdatedAt,
	}
	return PersonNameDALMeta{
		PersonNameMeta:       GetPersonNameMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta PersonNameDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField] {
	return meta.BasicTypeDALMetaBase
}

func (meta PersonNameDALMeta) GetDirectDBValues(obj example_app_types.PersonNameWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.First,
		obj.MiddleInitial,
		obj.Last,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta PersonNameDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj example_app_types.PersonNameWithMeta) (example_app_types.PersonNameWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta PersonNameDALMeta) ScanDBNextRow(rows *sql.Rows) (example_app_types.PersonNameWithMeta, error) {
	var elem example_app_types.PersonNameWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.First,
		&elem.MiddleInitial,
		&elem.Last,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta PersonNameDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []example_app_types.PersonNameWithMeta) ([]example_app_types.PersonNameWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta PersonNameDALMeta) GetChangedFieldsAndValues(old, new example_app_types.PersonNameWithMeta, allowedFields []example_app_types.PersonNameField) ([]example_app_types.PersonNameField, []interface{}) {

	var colsWithValueChange []example_app_types.PersonNameField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(example_app_types.PersonNameField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_First, allowedFields) {
		if old.First != new.First {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_First)
			vals = append(vals, new.First)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_MiddleInitial, allowedFields) {
		if old.MiddleInitial != new.MiddleInitial {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_MiddleInitial)
			vals = append(vals, new.MiddleInitial)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_Last, allowedFields) {
		if old.Last != new.Last {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_Last)
			vals = append(vals, new.Last)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.PersonNameField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PersonNameField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta PersonNameDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[example_app_types.PersonNameWithMeta, example_app_types.PersonNameField], allowedFields []example_app_types.PersonNameField, elem example_app_types.PersonNameWithMeta) (example_app_types.PersonNameWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type PhoneNumberMeta struct {
	types.BasicTypeMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField]
}

func GetPhoneNumberMeta() PhoneNumberMeta {
	objMeta := types.BasicTypeMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField]{
		Name: naam.New("phone_number"),
		Fields: []example_app_types.PhoneNumberField{
			example_app_types.PhoneNumberField_ParentID,
			example_app_types.PhoneNumberField_ID,
			example_app_types.PhoneNumberField_CountryCode,
			example_app_types.PhoneNumberField_Number,
			example_app_types.PhoneNumberField_Extension,
			example_app_types.PhoneNumberField_CreatedAt,
			example_app_types.PhoneNumberField_UpdatedAt,
			example_app_types.PhoneNumberField_DeletedAt,
		},
	}
	return PhoneNumberMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta PhoneNumberMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField] {
	return meta.BasicTypeMetaBase
}

func (meta PhoneNumberMeta) SetMetaFieldValues(obj example_app_types.PhoneNumberWithMeta, now time.Time) example_app_types.PhoneNumberWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting PhoneNumbers: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting PhoneNumbers: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta PhoneNumberMeta) ConvertTimestampColumnsToUTC(obj example_app_types.PhoneNumberWithMeta) example_app_types.PhoneNumberWithMeta {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta PhoneNumberMeta) SetDefaultFieldValues(obj example_app_types.PhoneNumberWithMeta) example_app_types.PhoneNumberWithMeta {
	return obj
}

type PhoneNumberDALMeta struct {
	PhoneNumberMeta
	dalutil.BasicTypeDALMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField]
}

func GetPhoneNumberDALMeta() PhoneNumberDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField]{
		DatabaseColumnFields: []example_app_types.PhoneNumberField{
			example_app_types.PhoneNumberField_ParentID,
			example_app_types.PhoneNumberField_ID,
			example_app_types.PhoneNumberField_CountryCode,
			example_app_types.PhoneNumberField_Number,
			example_app_types.PhoneNumberField_Extension,
			example_app_types.PhoneNumberField_CreatedAt,
			example_app_types.PhoneNumberField_UpdatedAt,
			example_app_types.PhoneNumberField_DeletedAt,
		},
		DatabaseSubTableFields: []example_app_types.PhoneNumberField{},
		MutableOnlyByDALFields: []example_app_types.PhoneNumberField{
			example_app_types.PhoneNumberField_UpdatedAt,
			example_app_types.PhoneNumberField_DeletedAt,
		},
		NonMutableFields: []example_app_types.PhoneNumberField{
			example_app_types.PhoneNumberField_ParentID,
			example_app_types.PhoneNumberField_ID,
			example_app_types.PhoneNumberField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []example_app_types.PhoneNumberField{
			example_app_types.PhoneNumberField_CreatedAt,
			example_app_types.PhoneNumberField_UpdatedAt,
			example_app_types.PhoneNumberField_DeletedAt,
		},
		UpdatedAtField: example_app_types.PhoneNumberField_UpdatedAt,
	}
	return PhoneNumberDALMeta{
		PhoneNumberMeta:      GetPhoneNumberMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta PhoneNumberDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField] {
	return meta.BasicTypeDALMetaBase
}

func (meta PhoneNumberDALMeta) GetDirectDBValues(obj example_app_types.PhoneNumberWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.CountryCode,
		obj.Number,
		obj.Extension,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta PhoneNumberDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj example_app_types.PhoneNumberWithMeta) (example_app_types.PhoneNumberWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta PhoneNumberDALMeta) ScanDBNextRow(rows *sql.Rows) (example_app_types.PhoneNumberWithMeta, error) {
	var elem example_app_types.PhoneNumberWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.CountryCode,
		&elem.Number,
		&elem.Extension,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta PhoneNumberDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []example_app_types.PhoneNumberWithMeta) ([]example_app_types.PhoneNumberWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta PhoneNumberDALMeta) GetChangedFieldsAndValues(old, new example_app_types.PhoneNumberWithMeta, allowedFields []example_app_types.PhoneNumberField) ([]example_app_types.PhoneNumberField, []interface{}) {

	var colsWithValueChange []example_app_types.PhoneNumberField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(example_app_types.PhoneNumberField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_CountryCode, allowedFields) {
		if old.CountryCode != new.CountryCode {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_CountryCode)
			vals = append(vals, new.CountryCode)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_Number, allowedFields) {
		if old.Number != new.Number {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_Number)
			vals = append(vals, new.Number)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_Extension, allowedFields) {
		if old.Extension != new.Extension {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_Extension)
			vals = append(vals, new.Extension)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(example_app_types.PhoneNumberField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, example_app_types.PhoneNumberField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta PhoneNumberDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[example_app_types.PhoneNumberWithMeta, example_app_types.PhoneNumberField], allowedFields []example_app_types.PhoneNumberField, elem example_app_types.PhoneNumberWithMeta) (example_app_types.PhoneNumberWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
