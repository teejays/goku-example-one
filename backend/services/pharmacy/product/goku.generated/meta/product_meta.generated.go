package product

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

	entity_types "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/types"
)

type ProductEntityDALMeta struct {
	dalutil.EntityDALMetaBase[entity_types.Product, entity_types.ProductField]
}

func GetProductEntityDALMeta() ProductEntityDALMeta {
	meta := dalutil.EntityDALMetaBase[entity_types.Product, entity_types.ProductField]{
		DbTableName:      naam.New("tb_product"),
		BasicTypeDALMeta: GetProductDALMeta(),
	}
	return ProductEntityDALMeta{
		EntityDALMetaBase: meta,
	}
}

type ProductMeta struct {
	types.BasicTypeMetaBase[entity_types.Product, entity_types.ProductField]
}

func GetProductMeta() ProductMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.Product, entity_types.ProductField]{
		Name: naam.New("product"),
		Fields: []entity_types.ProductField{
			entity_types.ProductField_ID,
			entity_types.ProductField_MedicineID,
			entity_types.ProductField_Mass,
			entity_types.ProductField_Count,
			entity_types.ProductField_Name,
			entity_types.ProductField_CreatedAt,
			entity_types.ProductField_UpdatedAt,
			entity_types.ProductField_DeletedAt,
		},
	}
	return ProductMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta ProductMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.Product, entity_types.ProductField] {
	return meta.BasicTypeMetaBase
}

func (meta ProductMeta) SetMetaFieldValues(obj entity_types.Product, now time.Time) entity_types.Product {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Products: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Products: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta ProductMeta) ConvertTimestampColumnsToUTC(obj entity_types.Product) entity_types.Product {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta ProductMeta) SetDefaultFieldValues(obj entity_types.Product) entity_types.Product {
	return obj
}

type ProductDALMeta struct {
	ProductMeta
	dalutil.BasicTypeDALMetaBase[entity_types.Product, entity_types.ProductField]
}

func GetProductDALMeta() ProductDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.Product, entity_types.ProductField]{
		DatabaseColumnFields: []entity_types.ProductField{
			entity_types.ProductField_ID,
			entity_types.ProductField_MedicineID,
			entity_types.ProductField_Mass,
			entity_types.ProductField_Count,
			entity_types.ProductField_Name,
			entity_types.ProductField_CreatedAt,
			entity_types.ProductField_UpdatedAt,
			entity_types.ProductField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.ProductField{},
		MutableOnlyByDALFields: []entity_types.ProductField{
			entity_types.ProductField_UpdatedAt,
			entity_types.ProductField_DeletedAt,
		},
		NonMutableFields: []entity_types.ProductField{
			entity_types.ProductField_ID,
			entity_types.ProductField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.ProductField{
			entity_types.ProductField_CreatedAt,
			entity_types.ProductField_UpdatedAt,
			entity_types.ProductField_DeletedAt,
		},
		UpdatedAtField: entity_types.ProductField_UpdatedAt,
	}
	return ProductDALMeta{
		ProductMeta:          GetProductMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta ProductDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.Product, entity_types.ProductField] {
	return meta.BasicTypeDALMetaBase
}

func (meta ProductDALMeta) GetDirectDBValues(obj entity_types.Product) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.MedicineID,
		obj.Mass,
		obj.Count,
		obj.Name,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta ProductDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.Product) (entity_types.Product, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta ProductDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.Product, error) {
	var elem entity_types.Product
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.MedicineID,
		&elem.Mass,
		&elem.Count,
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

func (meta ProductDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.Product) ([]entity_types.Product, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta ProductDALMeta) GetChangedFieldsAndValues(old, new entity_types.Product, allowedFields []entity_types.ProductField) ([]entity_types.ProductField, []interface{}) {

	var colsWithValueChange []entity_types.ProductField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.ProductField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_MedicineID, allowedFields) {
		if old.MedicineID != new.MedicineID {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_MedicineID)
			vals = append(vals, new.MedicineID)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_Mass, allowedFields) {
		if old.Mass != new.Mass {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_Mass)
			vals = append(vals, new.Mass)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_Count, allowedFields) {
		if old.Count != new.Count {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_Count)
			vals = append(vals, new.Count)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.ProductField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.ProductField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta ProductDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.Product, entity_types.ProductField], allowedFields []entity_types.ProductField, elem entity_types.Product) (entity_types.Product, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}
