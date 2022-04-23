package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/teejays/clog"

	"github.com/teejays/goku-util/client/db"
	"github.com/teejays/goku-util/dalutil"
	"github.com/teejays/goku-util/naam"
	"github.com/teejays/goku-util/scalars"
	"github.com/teejays/goku-util/types"

	app_types "github.com/teejays/goku-example-one/backend/goku.generated/types"
	entity_types "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/types"
)

type UserEntityDALMeta struct {
	dalutil.EntityDALMetaBase[entity_types.User, entity_types.UserField]
}

func GetUserEntityDALMeta() UserEntityDALMeta {
	meta := dalutil.EntityDALMetaBase[entity_types.User, entity_types.UserField]{
		DbTableName:      naam.New("tb_user"),
		BasicTypeDALMeta: GetUserDALMeta(),
	}
	return UserEntityDALMeta{
		EntityDALMetaBase: meta,
	}
}

type UserMeta struct {
	types.BasicTypeMetaBase[entity_types.User, entity_types.UserField]
}

func GetUserMeta() UserMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.User, entity_types.UserField]{
		Name: naam.New("user"),
		Fields: []entity_types.UserField{
			entity_types.UserField_ID,
			entity_types.UserField_Name,
			entity_types.UserField_Name_First,
			entity_types.UserField_Name_MiddleInitial,
			entity_types.UserField_Name_Last,
			entity_types.UserField_Email,
			entity_types.UserField_PhoneNumber,
			entity_types.UserField_PhoneNumber_CountryCode,
			entity_types.UserField_PhoneNumber_Number,
			entity_types.UserField_PhoneNumber_Extension,
			entity_types.UserField_PasswordHash,
			entity_types.UserField_CreatedAt,
			entity_types.UserField_UpdatedAt,
			entity_types.UserField_DeletedAt,
		},
	}
	return UserMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta UserMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.User, entity_types.UserField] {
	return meta.BasicTypeMetaBase
}

func (meta UserMeta) SetMetaFieldValues(obj entity_types.User, now time.Time) entity_types.User {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Users: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Users: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta UserMeta) ConvertTimestampColumnsToUTC(obj entity_types.User) entity_types.User {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta UserMeta) SetDefaultFieldValues(obj entity_types.User) entity_types.User {
	if obj.Email == "" {
		obj.Email = "no@email.com"
	}
	return obj
}

type UserDALMeta struct {
	UserMeta
	dalutil.BasicTypeDALMetaBase[entity_types.User, entity_types.UserField]
}

func GetUserDALMeta() UserDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.User, entity_types.UserField]{
		DatabaseColumnFields: []entity_types.UserField{
			entity_types.UserField_ID,
			entity_types.UserField_Name_First,
			entity_types.UserField_Name_MiddleInitial,
			entity_types.UserField_Name_Last,
			entity_types.UserField_Email,
			entity_types.UserField_PhoneNumber_CountryCode,
			entity_types.UserField_PhoneNumber_Number,
			entity_types.UserField_PhoneNumber_Extension,
			entity_types.UserField_PasswordHash,
			entity_types.UserField_CreatedAt,
			entity_types.UserField_UpdatedAt,
			entity_types.UserField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.UserField{},
		MutableOnlyByDALFields: []entity_types.UserField{
			entity_types.UserField_UpdatedAt,
			entity_types.UserField_DeletedAt,
		},
		NonMutableFields: []entity_types.UserField{
			entity_types.UserField_ID,
			entity_types.UserField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.UserField{
			entity_types.UserField_CreatedAt,
			entity_types.UserField_UpdatedAt,
			entity_types.UserField_DeletedAt,
		},
		UpdatedAtField: entity_types.UserField_UpdatedAt,
	}
	return UserDALMeta{
		UserMeta:             GetUserMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta UserDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.User, entity_types.UserField] {
	return meta.BasicTypeDALMetaBase
}

func (meta UserDALMeta) GetDirectDBValues(obj entity_types.User) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values
	if obj.PhoneNumber == nil {
		obj.PhoneNumber = &app_types.PhoneNumber{}
	}

	var vals = []interface{}{
		obj.ID,
		obj.Name.First,
		obj.Name.MiddleInitial,
		obj.Name.Last,
		obj.Email,
		obj.PhoneNumber.CountryCode,
		obj.PhoneNumber.Number,
		obj.PhoneNumber.Extension,
		obj.PasswordHash,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta UserDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.User) (entity_types.User, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta UserDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.User, error) {
	var elem entity_types.User
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.
	elem.PhoneNumber = &app_types.PhoneNumber{}

	err := rows.Scan(
		&elem.ID,
		&elem.Name.First,
		&elem.Name.MiddleInitial,
		&elem.Name.Last,
		&elem.Email,
		&elem.PhoneNumber.CountryCode,
		&elem.PhoneNumber.Number,
		&elem.PhoneNumber.Extension,
		&elem.PasswordHash,
		&elem.CreatedAt,
		&elem.UpdatedAt,
		&elem.DeletedAt,
	)
	if err != nil {
		return elem, fmt.Errorf("sql.Row scan error: %w", err)
	}

	// If a nested pointer field (optional) if same as an empty struct, make it nil
	{
		var empty app_types.PhoneNumber
		if *elem.PhoneNumber == empty {
			elem.PhoneNumber = nil
		}
	}

	elem = meta.ConvertTimestampColumnsToUTC(elem)
	return elem, nil
}

func (meta UserDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.User) ([]entity_types.User, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta UserDALMeta) GetChangedFieldsAndValues(old, new entity_types.User, allowedFields []entity_types.UserField) ([]entity_types.UserField, []interface{}) {

	var colsWithValueChange []entity_types.UserField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.UserField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_Name_First, allowedFields) {
		if old.Name.First != new.Name.First {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_Name_First)
			vals = append(vals, new.Name.First)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_Name_MiddleInitial, allowedFields) {
		if old.Name.MiddleInitial != new.Name.MiddleInitial {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_Name_MiddleInitial)
			vals = append(vals, new.Name.MiddleInitial)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_Name_Last, allowedFields) {
		if old.Name.Last != new.Name.Last {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_Name_Last)
			vals = append(vals, new.Name.Last)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_Email, allowedFields) {
		if old.Email != new.Email {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_Email)
			vals = append(vals, new.Email)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_PhoneNumber_CountryCode, allowedFields) {
		if old.PhoneNumber.CountryCode != new.PhoneNumber.CountryCode {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_PhoneNumber_CountryCode)
			vals = append(vals, new.PhoneNumber.CountryCode)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_PhoneNumber_Number, allowedFields) {
		if old.PhoneNumber.Number != new.PhoneNumber.Number {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_PhoneNumber_Number)
			vals = append(vals, new.PhoneNumber.Number)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_PhoneNumber_Extension, allowedFields) {
		if old.PhoneNumber.Extension != new.PhoneNumber.Extension {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_PhoneNumber_Extension)
			vals = append(vals, new.PhoneNumber.Extension)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_PasswordHash, allowedFields) {
		if old.PasswordHash != new.PasswordHash {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_PasswordHash)
			vals = append(vals, new.PasswordHash)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.UserField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.UserField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta UserDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.User, entity_types.UserField], allowedFields []entity_types.UserField, elem entity_types.User) (entity_types.User, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
