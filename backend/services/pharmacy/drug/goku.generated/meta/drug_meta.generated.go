package drug

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

	entity_types "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/types"
)

type DrugEntityDALMeta struct {
	dalutil.EntityDALMetaBase[entity_types.Drug, entity_types.DrugField]
}

func GetDrugEntityDALMeta() DrugEntityDALMeta {
	meta := dalutil.EntityDALMetaBase[entity_types.Drug, entity_types.DrugField]{
		DbTableName:      naam.New("tb_drug"),
		BasicTypeDALMeta: GetDrugDALMeta(),
	}
	return DrugEntityDALMeta{
		EntityDALMetaBase: meta,
	}
}

type DrugMeta struct {
	types.BasicTypeMetaBase[entity_types.Drug, entity_types.DrugField]
}

func GetDrugMeta() DrugMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.Drug, entity_types.DrugField]{
		Name: naam.New("drug"),
		Fields: []entity_types.DrugField{
			entity_types.DrugField_ID,
			entity_types.DrugField_Name,
			entity_types.DrugField_CreatedAt,
			entity_types.DrugField_UpdatedAt,
			entity_types.DrugField_DeletedAt,
		},
	}
	return DrugMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta DrugMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.Drug, entity_types.DrugField] {
	return meta.BasicTypeMetaBase
}

func (meta DrugMeta) SetMetaFieldValues(obj entity_types.Drug, now time.Time) entity_types.Drug {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Drugs: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Drugs: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta DrugMeta) ConvertTimestampColumnsToUTC(obj entity_types.Drug) entity_types.Drug {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta DrugMeta) SetDefaultFieldValues(obj entity_types.Drug) entity_types.Drug {
	return obj
}

type DrugDALMeta struct {
	DrugMeta
	dalutil.BasicTypeDALMetaBase[entity_types.Drug, entity_types.DrugField]
}

func GetDrugDALMeta() DrugDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.Drug, entity_types.DrugField]{
		DatabaseColumnFields: []entity_types.DrugField{
			entity_types.DrugField_ID,
			entity_types.DrugField_Name,
			entity_types.DrugField_CreatedAt,
			entity_types.DrugField_UpdatedAt,
			entity_types.DrugField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.DrugField{},
		MutableOnlyByDALFields: []entity_types.DrugField{
			entity_types.DrugField_UpdatedAt,
			entity_types.DrugField_DeletedAt,
		},
		NonMutableFields: []entity_types.DrugField{
			entity_types.DrugField_ID,
			entity_types.DrugField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.DrugField{
			entity_types.DrugField_CreatedAt,
			entity_types.DrugField_UpdatedAt,
			entity_types.DrugField_DeletedAt,
		},
		UpdatedAtField: entity_types.DrugField_UpdatedAt,
	}
	return DrugDALMeta{
		DrugMeta:             GetDrugMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta DrugDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.Drug, entity_types.DrugField] {
	return meta.BasicTypeDALMetaBase
}

func (meta DrugDALMeta) GetDirectDBValues(obj entity_types.Drug) []interface{} {
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

func (meta DrugDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.Drug) (entity_types.Drug, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta DrugDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.Drug, error) {
	var elem entity_types.Drug
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

func (meta DrugDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.Drug) ([]entity_types.Drug, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta DrugDALMeta) GetChangedFieldsAndValues(old, new entity_types.Drug, allowedFields []entity_types.DrugField) ([]entity_types.DrugField, []interface{}) {

	var colsWithValueChange []entity_types.DrugField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.DrugField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.DrugField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.DrugField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, entity_types.DrugField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(entity_types.DrugField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.DrugField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.DrugField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.DrugField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.DrugField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.DrugField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta DrugDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.Drug, entity_types.DrugField], allowedFields []entity_types.DrugField, elem entity_types.Drug) (entity_types.Drug, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
