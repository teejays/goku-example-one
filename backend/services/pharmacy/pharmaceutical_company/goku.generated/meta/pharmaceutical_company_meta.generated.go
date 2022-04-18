package pharmaceutical_company

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

	entity_types "github.com/teejays/goku/example/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
)

type PharmaceuticalCompanyEntityDALMeta struct {
	dalutil.EntityDALMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]
}

func GetPharmaceuticalCompanyEntityDALMeta() PharmaceuticalCompanyEntityDALMeta {
	meta := dalutil.EntityDALMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]{
		DbTableName:      naam.New("tb_pharmaceutical_company"),
		BasicTypeDALMeta: GetPharmaceuticalCompanyDALMeta(),
	}
	return PharmaceuticalCompanyEntityDALMeta{
		EntityDALMetaBase: meta,
	}
}

type PharmaceuticalCompanyMeta struct {
	types.BasicTypeMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]
}

func GetPharmaceuticalCompanyMeta() PharmaceuticalCompanyMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]{
		Name: naam.New("pharmaceutical_company"),
		Fields: []entity_types.PharmaceuticalCompanyField{
			entity_types.PharmaceuticalCompanyField_ID,
			entity_types.PharmaceuticalCompanyField_Name,
			entity_types.PharmaceuticalCompanyField_CreatedAt,
			entity_types.PharmaceuticalCompanyField_UpdatedAt,
			entity_types.PharmaceuticalCompanyField_DeletedAt,
		},
	}
	return PharmaceuticalCompanyMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta PharmaceuticalCompanyMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField] {
	return meta.BasicTypeMetaBase
}

func (meta PharmaceuticalCompanyMeta) SetMetaFieldValues(obj entity_types.PharmaceuticalCompany, now time.Time) entity_types.PharmaceuticalCompany {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting PharmaceuticalCompanies: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting PharmaceuticalCompanies: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta PharmaceuticalCompanyMeta) ConvertTimestampColumnsToUTC(obj entity_types.PharmaceuticalCompany) entity_types.PharmaceuticalCompany {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta PharmaceuticalCompanyMeta) SetDefaultFieldValues(obj entity_types.PharmaceuticalCompany) entity_types.PharmaceuticalCompany {
	return obj
}

type PharmaceuticalCompanyDALMeta struct {
	PharmaceuticalCompanyMeta
	dalutil.BasicTypeDALMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]
}

func GetPharmaceuticalCompanyDALMeta() PharmaceuticalCompanyDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField]{
		DatabaseColumnFields: []entity_types.PharmaceuticalCompanyField{
			entity_types.PharmaceuticalCompanyField_ID,
			entity_types.PharmaceuticalCompanyField_Name,
			entity_types.PharmaceuticalCompanyField_CreatedAt,
			entity_types.PharmaceuticalCompanyField_UpdatedAt,
			entity_types.PharmaceuticalCompanyField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.PharmaceuticalCompanyField{},
		MutableOnlyByDALFields: []entity_types.PharmaceuticalCompanyField{
			entity_types.PharmaceuticalCompanyField_UpdatedAt,
			entity_types.PharmaceuticalCompanyField_DeletedAt,
		},
		NonMutableFields: []entity_types.PharmaceuticalCompanyField{
			entity_types.PharmaceuticalCompanyField_ID,
			entity_types.PharmaceuticalCompanyField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.PharmaceuticalCompanyField{
			entity_types.PharmaceuticalCompanyField_CreatedAt,
			entity_types.PharmaceuticalCompanyField_UpdatedAt,
			entity_types.PharmaceuticalCompanyField_DeletedAt,
		},
		UpdatedAtField: entity_types.PharmaceuticalCompanyField_UpdatedAt,
	}
	return PharmaceuticalCompanyDALMeta{
		PharmaceuticalCompanyMeta: GetPharmaceuticalCompanyMeta(),
		BasicTypeDALMetaBase:      dalMetaBase,
	}
}

func (meta PharmaceuticalCompanyDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField] {
	return meta.BasicTypeDALMetaBase
}

func (meta PharmaceuticalCompanyDALMeta) GetDirectDBValues(obj entity_types.PharmaceuticalCompany) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Name,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta PharmaceuticalCompanyDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.PharmaceuticalCompany) (entity_types.PharmaceuticalCompany, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta PharmaceuticalCompanyDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.PharmaceuticalCompany, error) {
	var elem entity_types.PharmaceuticalCompany
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Name,
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

func (meta PharmaceuticalCompanyDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.PharmaceuticalCompany) ([]entity_types.PharmaceuticalCompany, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta PharmaceuticalCompanyDALMeta) GetChangedFieldsAndValues(old, new entity_types.PharmaceuticalCompany, allowedFields []entity_types.PharmaceuticalCompanyField) ([]entity_types.PharmaceuticalCompanyField, []interface{}) {

	var colsWithValueChange []entity_types.PharmaceuticalCompanyField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.PharmaceuticalCompanyField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.PharmaceuticalCompanyField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.PharmaceuticalCompanyField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, entity_types.PharmaceuticalCompanyField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(entity_types.PharmaceuticalCompanyField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.PharmaceuticalCompanyField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.PharmaceuticalCompanyField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.PharmaceuticalCompanyField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.PharmaceuticalCompanyField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.PharmaceuticalCompanyField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta PharmaceuticalCompanyDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.PharmaceuticalCompany, entity_types.PharmaceuticalCompanyField], allowedFields []entity_types.PharmaceuticalCompanyField, elem entity_types.PharmaceuticalCompany) (entity_types.PharmaceuticalCompany, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
