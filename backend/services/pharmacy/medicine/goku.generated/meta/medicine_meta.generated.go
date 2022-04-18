package medicine

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

	entity_types "github.com/teejays/goku/example/backend/services/pharmacy/medicine/goku.generated/types"
)

type IngredientMeta struct {
	types.BasicTypeMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField]
}

func GetIngredientMeta() IngredientMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField]{
		Name: naam.New("ingredient"),
		Fields: []entity_types.IngredientField{
			entity_types.IngredientField_ParentID,
			entity_types.IngredientField_ID,
			entity_types.IngredientField_DrugID,
			entity_types.IngredientField_IsPrimaryIngredient,
			entity_types.IngredientField_CreatedAt,
			entity_types.IngredientField_UpdatedAt,
			entity_types.IngredientField_DeletedAt,
		},
	}
	return IngredientMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta IngredientMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField] {
	return meta.BasicTypeMetaBase
}

func (meta IngredientMeta) SetMetaFieldValues(obj entity_types.IngredientWithMeta, now time.Time) entity_types.IngredientWithMeta {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Ingredients: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Ingredients: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta IngredientMeta) ConvertTimestampColumnsToUTC(obj entity_types.IngredientWithMeta) entity_types.IngredientWithMeta {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta IngredientMeta) SetDefaultFieldValues(obj entity_types.IngredientWithMeta) entity_types.IngredientWithMeta {
	return obj
}

type IngredientDALMeta struct {
	IngredientMeta
	dalutil.BasicTypeDALMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField]
}

func GetIngredientDALMeta() IngredientDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField]{
		DatabaseColumnFields: []entity_types.IngredientField{
			entity_types.IngredientField_ParentID,
			entity_types.IngredientField_ID,
			entity_types.IngredientField_DrugID,
			entity_types.IngredientField_IsPrimaryIngredient,
			entity_types.IngredientField_CreatedAt,
			entity_types.IngredientField_UpdatedAt,
			entity_types.IngredientField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.IngredientField{},
		MutableOnlyByDALFields: []entity_types.IngredientField{
			entity_types.IngredientField_UpdatedAt,
			entity_types.IngredientField_DeletedAt,
		},
		NonMutableFields: []entity_types.IngredientField{
			entity_types.IngredientField_ParentID,
			entity_types.IngredientField_ID,
			entity_types.IngredientField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.IngredientField{
			entity_types.IngredientField_CreatedAt,
			entity_types.IngredientField_UpdatedAt,
			entity_types.IngredientField_DeletedAt,
		},
		UpdatedAtField: entity_types.IngredientField_UpdatedAt,
	}
	return IngredientDALMeta{
		IngredientMeta:       GetIngredientMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta IngredientDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.IngredientWithMeta, entity_types.IngredientField] {
	return meta.BasicTypeDALMetaBase
}

func (meta IngredientDALMeta) GetDirectDBValues(obj entity_types.IngredientWithMeta) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ParentID,
		obj.ID,
		obj.DrugID,
		obj.IsPrimaryIngredient,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta IngredientDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.IngredientWithMeta) (entity_types.IngredientWithMeta, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables

	return obj, nil
}

func (meta IngredientDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.IngredientWithMeta, error) {
	var elem entity_types.IngredientWithMeta
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ParentID,
		&elem.ID,
		&elem.DrugID,
		&elem.IsPrimaryIngredient,
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

func (meta IngredientDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.IngredientWithMeta) ([]entity_types.IngredientWithMeta, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields

	return elems, nil
}

func (meta IngredientDALMeta) GetChangedFieldsAndValues(old, new entity_types.IngredientWithMeta, allowedFields []entity_types.IngredientField) ([]entity_types.IngredientField, []interface{}) {

	var colsWithValueChange []entity_types.IngredientField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.IngredientField_ParentID, allowedFields) {
		if old.ParentID != new.ParentID {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_ParentID)
			vals = append(vals, new.ParentID)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_DrugID, allowedFields) {
		if old.DrugID != new.DrugID {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_DrugID)
			vals = append(vals, new.DrugID)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_IsPrimaryIngredient, allowedFields) {
		if old.IsPrimaryIngredient != new.IsPrimaryIngredient {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_IsPrimaryIngredient)
			vals = append(vals, new.IsPrimaryIngredient)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.IngredientField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.IngredientField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta IngredientDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.IngredientWithMeta, entity_types.IngredientField], allowedFields []entity_types.IngredientField, elem entity_types.IngredientWithMeta) (entity_types.IngredientWithMeta, error) {
	// Update Nested (1:1 & 1:Many)

	return elem, nil
}

type MedicineEntityDALMeta struct {
	dalutil.EntityDALMetaBase[entity_types.Medicine, entity_types.MedicineField]
}

func GetMedicineEntityDALMeta() MedicineEntityDALMeta {
	meta := dalutil.EntityDALMetaBase[entity_types.Medicine, entity_types.MedicineField]{
		DbTableName:      naam.New("tb_medicine"),
		BasicTypeDALMeta: GetMedicineDALMeta(),
	}
	return MedicineEntityDALMeta{
		EntityDALMetaBase: meta,
	}
}

type MedicineMeta struct {
	types.BasicTypeMetaBase[entity_types.Medicine, entity_types.MedicineField]
}

func GetMedicineMeta() MedicineMeta {
	objMeta := types.BasicTypeMetaBase[entity_types.Medicine, entity_types.MedicineField]{
		Name: naam.New("medicine"),
		Fields: []entity_types.MedicineField{
			entity_types.MedicineField_ID,
			entity_types.MedicineField_Name,
			entity_types.MedicineField_CompanyID,
			entity_types.MedicineField_PrimaryIngredient,
			entity_types.MedicineField_PrimaryIngredient_DrugID,
			entity_types.MedicineField_PrimaryIngredient_IsPrimaryIngredient,
			entity_types.MedicineField_Ingredients,
			entity_types.MedicineField_Ingredients_ParentID,
			entity_types.MedicineField_Ingredients_ID,
			entity_types.MedicineField_Ingredients_DrugID,
			entity_types.MedicineField_Ingredients_IsPrimaryIngredient,
			entity_types.MedicineField_Ingredients_CreatedAt,
			entity_types.MedicineField_Ingredients_UpdatedAt,
			entity_types.MedicineField_Ingredients_DeletedAt,
			entity_types.MedicineField_ModeOfDelivery,
			entity_types.MedicineField_CreatedAt,
			entity_types.MedicineField_UpdatedAt,
			entity_types.MedicineField_DeletedAt,
		},
	}
	return MedicineMeta{
		BasicTypeMetaBase: objMeta,
	}
}

func (meta MedicineMeta) GetBasicTypeMetaBase() types.BasicTypeMetaBase[entity_types.Medicine, entity_types.MedicineField] {
	return meta.BasicTypeMetaBase
}

func (meta MedicineMeta) SetMetaFieldValues(obj entity_types.Medicine, now time.Time) entity_types.Medicine {
	nowScalar := scalars.NewTime(now)
	if obj.ID.IsEmpty() {
		obj.ID = scalars.NewID()
	} else {
		clog.Warnf("Inserting Medicines: entity already has ID set: %s", obj.ID)
	}
	if obj.CreatedAt.IsZero() {
		obj.CreatedAt = nowScalar
	} else {
		clog.Warnf("Inserting Medicines: entity already has CreatedAt set: %s", obj.CreatedAt)
	}
	obj.UpdatedAt = nowScalar
	return obj
}

func (meta MedicineMeta) ConvertTimestampColumnsToUTC(obj entity_types.Medicine) entity_types.Medicine {
	obj.CreatedAt = scalars.NewTime(obj.CreatedAt.UTC())
	obj.UpdatedAt = scalars.NewTime(obj.UpdatedAt.UTC())

	if obj.DeletedAt != nil {
		*obj.DeletedAt = scalars.NewTime(obj.DeletedAt.UTC())
	}
	return obj
}

func (meta MedicineMeta) SetDefaultFieldValues(obj entity_types.Medicine) entity_types.Medicine {
	return obj
}

type MedicineDALMeta struct {
	MedicineMeta
	dalutil.BasicTypeDALMetaBase[entity_types.Medicine, entity_types.MedicineField]
}

func GetMedicineDALMeta() MedicineDALMeta {
	dalMetaBase := dalutil.BasicTypeDALMetaBase[entity_types.Medicine, entity_types.MedicineField]{
		DatabaseColumnFields: []entity_types.MedicineField{
			entity_types.MedicineField_ID,
			entity_types.MedicineField_Name,
			entity_types.MedicineField_CompanyID,
			entity_types.MedicineField_PrimaryIngredient_DrugID,
			entity_types.MedicineField_PrimaryIngredient_IsPrimaryIngredient,
			entity_types.MedicineField_ModeOfDelivery,
			entity_types.MedicineField_CreatedAt,
			entity_types.MedicineField_UpdatedAt,
			entity_types.MedicineField_DeletedAt,
		},
		DatabaseSubTableFields: []entity_types.MedicineField{
			entity_types.MedicineField_Ingredients,
		},
		MutableOnlyByDALFields: []entity_types.MedicineField{
			entity_types.MedicineField_Ingredients_UpdatedAt,
			entity_types.MedicineField_Ingredients_DeletedAt,
			entity_types.MedicineField_UpdatedAt,
			entity_types.MedicineField_DeletedAt,
		},
		NonMutableFields: []entity_types.MedicineField{
			entity_types.MedicineField_ID,
			entity_types.MedicineField_Ingredients_ParentID,
			entity_types.MedicineField_Ingredients_ID,
			entity_types.MedicineField_Ingredients_CreatedAt,
			entity_types.MedicineField_CreatedAt,
		},
		DatabaseColumnTimestampFields: []entity_types.MedicineField{
			entity_types.MedicineField_CreatedAt,
			entity_types.MedicineField_UpdatedAt,
			entity_types.MedicineField_DeletedAt,
		},
		UpdatedAtField: entity_types.MedicineField_UpdatedAt,
	}
	return MedicineDALMeta{
		MedicineMeta:         GetMedicineMeta(),
		BasicTypeDALMetaBase: dalMetaBase,
	}
}

func (meta MedicineDALMeta) GetDALMetaBase() dalutil.BasicTypeDALMetaBase[entity_types.Medicine, entity_types.MedicineField] {
	return meta.BasicTypeDALMetaBase
}

func (meta MedicineDALMeta) GetDirectDBValues(obj entity_types.Medicine) []interface{} {
	// If a nested field (in same DB table) is nil e.g. Foo.Bar, we'll hit a nil pointer panic if accessing the underling values e.g. Foo.Bar.Baz. Hence, replace nil with empty values

	var vals = []interface{}{
		obj.ID,
		obj.Name,
		obj.CompanyID,
		obj.PrimaryIngredient.DrugID,
		obj.PrimaryIngredient.IsPrimaryIngredient,
		obj.ModeOfDelivery,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.DeletedAt,
	}
	return vals
}

func (meta MedicineDALMeta) AddSubTableFieldsToDB(ctx context.Context, conn db.Connection, params db.InsertTypeParams, obj entity_types.Medicine) (entity_types.Medicine, error) {

	// Insert 1:1 sub-tables

	// Insert 1:Many sub-tables
	// Insert Ingredients
	if len(obj.Ingredients) > 0 {
		for j := range obj.Ingredients {
			obj.Ingredients[j].ParentID = obj.ID
		}
		req := db.InsertTypeParams{
			TableName: params.TableName + "_" + entity_types.MedicineField_Ingredients.Name().FormatSQLTable(),
		}
		subMeta := GetIngredientDALMeta()
		subElems, err := dalutil.BatchAddType[entity_types.IngredientWithMeta, entity_types.IngredientField](ctx, conn, req, subMeta, obj.Ingredients...)
		if err != nil {
			return obj, fmt.Errorf("Inserting Ingredients: %w", err)
		}
		obj.Ingredients = subElems
	}

	return obj, nil
}

func (meta MedicineDALMeta) ScanDBNextRow(rows *sql.Rows) (entity_types.Medicine, error) {
	var elem entity_types.Medicine
	// For any pointer (optional) nested field e.g. Foo.Nested.FieldA, create a new instance of struct to prevent nil pointers when Nested is nil.

	err := rows.Scan(
		&elem.ID,
		&elem.Name,
		&elem.CompanyID,
		&elem.PrimaryIngredient.DrugID,
		&elem.PrimaryIngredient.IsPrimaryIngredient,
		&elem.ModeOfDelivery,
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

func (meta MedicineDALMeta) FetchSubTableFields(ctx context.Context, conn db.Connection, params db.ListTypeByIDsParams, elems []entity_types.Medicine) ([]entity_types.Medicine, error) {
	// Unique Primary IDs of the fetched type
	var ids []scalars.ID
	for _, elem := range elems {
		ids = append(ids, elem.ID)
	}

	// Step 1: Get the Nested (1:1) fields

	// Step 2: Get the Nested (1:Many) fields
	{
		// Fetch Ingredients (type IngredientWithMeta)
		subParams := db.ListTypeByIDsParams{
			TableName: params.TableName + "_" + "ingredients",
			IDColumn:  "parent_id",
			IDs:       ids,
		}
		subMeta := GetIngredientDALMeta()
		subResp, err := dalutil.ListTypeByIDs[entity_types.IngredientWithMeta, entity_types.IngredientField](ctx, conn, subParams, subMeta)
		if err != nil {
			return nil, fmt.Errorf("fetching Ingredient: %w", err)
		}

		// assign the right Ingredient to the right elem
		var subElemMap = make(map[scalars.ID][]entity_types.IngredientWithMeta)
		for _, subElem := range subResp.Items {
			subElemMap[subElem.ParentID] = append(subElemMap[subElem.ParentID], subElem)
		}

		for i := range elems {
			subElems := subElemMap[elems[i].ID]
			elems[i].Ingredients = subElems
		}
	}

	return elems, nil
}

func (meta MedicineDALMeta) GetChangedFieldsAndValues(old, new entity_types.Medicine, allowedFields []entity_types.MedicineField) ([]entity_types.MedicineField, []interface{}) {

	var colsWithValueChange []entity_types.MedicineField // columns that actually have a value change so required in update statement
	var vals []interface{}

	// Get Values (and check if there is a change)
	if types.IsFieldInFields(entity_types.MedicineField_ID, allowedFields) {
		if old.ID != new.ID {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_ID)
			vals = append(vals, new.ID)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_Name, allowedFields) {
		if old.Name != new.Name {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_Name)
			vals = append(vals, new.Name)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_CompanyID, allowedFields) {
		if old.CompanyID != new.CompanyID {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_CompanyID)
			vals = append(vals, new.CompanyID)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_PrimaryIngredient_DrugID, allowedFields) {
		if old.PrimaryIngredient.DrugID != new.PrimaryIngredient.DrugID {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_PrimaryIngredient_DrugID)
			vals = append(vals, new.PrimaryIngredient.DrugID)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_PrimaryIngredient_IsPrimaryIngredient, allowedFields) {
		if old.PrimaryIngredient.IsPrimaryIngredient != new.PrimaryIngredient.IsPrimaryIngredient {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_PrimaryIngredient_IsPrimaryIngredient)
			vals = append(vals, new.PrimaryIngredient.IsPrimaryIngredient)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_ModeOfDelivery, allowedFields) {
		if old.ModeOfDelivery != new.ModeOfDelivery {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_ModeOfDelivery)
			vals = append(vals, new.ModeOfDelivery)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_CreatedAt, allowedFields) {
		if old.CreatedAt != new.CreatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_CreatedAt)
			vals = append(vals, new.CreatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_UpdatedAt, allowedFields) {
		if old.UpdatedAt != new.UpdatedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_UpdatedAt)
			vals = append(vals, new.UpdatedAt)
		}
	}
	if types.IsFieldInFields(entity_types.MedicineField_DeletedAt, allowedFields) {
		if old.DeletedAt != new.DeletedAt {
			colsWithValueChange = append(colsWithValueChange, entity_types.MedicineField_DeletedAt)
			vals = append(vals, new.DeletedAt)
		}
	}

	return colsWithValueChange, vals
}

func (meta MedicineDALMeta) UpdateSubTableFields(ctx context.Context, conn db.Connection, req dalutil.UpdateTypeRequest[entity_types.Medicine, entity_types.MedicineField], allowedFields []entity_types.MedicineField, elem entity_types.Medicine) (entity_types.Medicine, error) {
	// Update Nested (1:1 & 1:Many)
	// Update Ingredients
	{
		for i := range elem.Ingredients {
			subElem := elem.Ingredients[i]
			// Populate subElem.ParentID in case it is missing
			if !subElem.ParentID.IsEmpty() && subElem.ParentID != elem.ID {
				return elem, fmt.Errorf("Updating Medicine: Nested object Ingredients's ParentID (%s) is different than the parent Medicine's ID (%s)", subElem.ParentID, elem.ID)
			}
			subElem.ParentID = elem.ID
			// TODO: If the 1:1 subobject does not have an ID, that's okay since we don't always expect users to provide
			// IDs for subobjects, so fetch it.
			subElemTableName := req.TableName + "_" + entity_types.MedicineField_Ingredients.Name().FormatSQLTable()
			subReq := dalutil.UpdateTypeRequest[entity_types.IngredientWithMeta, entity_types.IngredientField]{
				TableName: subElemTableName,
				Object:    subElem,
				// Fields: <empty> : include all fields as we can't yet handle nested fields
				// ExcludeFields: <empty> : include all fields as we can't yet handle nested fields
			}
			subMeta := GetIngredientDALMeta()
			subResp, err := dalutil.UpdateType[entity_types.IngredientWithMeta, entity_types.IngredientField](ctx, conn, subReq, subMeta)
			if err != nil {
				return elem, fmt.Errorf("Updating Ingredients: %w", err)
			}
			elem.Ingredients[i] = subResp.Object

		}
	}

	return elem, nil
}
