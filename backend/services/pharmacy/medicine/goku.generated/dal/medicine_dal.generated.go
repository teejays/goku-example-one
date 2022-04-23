package medicine

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect
	"github.com/teejays/clog"

	"github.com/teejays/goku-util/client/db"
	"github.com/teejays/goku-util/dalutil"
	filterlib "github.com/teejays/goku-util/filter"
	"github.com/teejays/goku-util/panics"
	"github.com/teejays/goku-util/scalars"

	dal_service "github.com/teejays/goku-example-one/backend/services/pharmacy/goku.generated/dal"
	meta_entity "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/meta"
	types_entity "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/types"
)

// MedicineEntityDAL encapsulates DAL methods for types that fall under Medicine
type MedicineEntityDAL struct {
	dal_service.PharmacyServiceDAL
}

// ListTypeIngredientRequest is the request object for ListTypeIngredient method
type ListTypeIngredientRequest = dalutil.ListTypeRequest[types_entity.IngredientFilter]

// ListTypeIngredientResponse is the response object for ListTypeIngredient method
type ListTypeIngredientResponse = dalutil.ListTypeResponse[types_entity.IngredientWithMeta]

// ConstructListQueryBuilderForTypeIngredient provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Ingredient items that belong to `params.TableName` and match the filter.
func (d MedicineEntityDAL) ConstructListQueryBuilderForTypeIngredient(ctx context.Context, conn db.Connection, req ListTypeIngredientRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetIngredientIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetIngredientMeta()

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables

	// Handle Nested objects 1:1

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ParentID != nil {
		// inject a Where filter for ParentID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ParentID, ds, "parent_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DrugID != nil {
		// inject a Where filter for DrugID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DrugID, ds, "drug_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.IsPrimaryIngredient != nil {
		// inject a Where filter for IsPrimaryIngredient
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.IsPrimaryIngredient, ds, "is_primary_ingredient")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d MedicineEntityDAL) GetQueryByTextFilterForTypeIngredient(ctx context.Context, queryText string) (*types_entity.IngredientFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.IngredientFilter

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.IngredientFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// AddMedicine handles insertion of a single Medicine entity in the database.
func (d MedicineEntityDAL) AddMedicine(ctx context.Context, conn db.Connection, elem types_entity.Medicine) (types_entity.Medicine, error) {
	elems, err := d.BatchAddMedicine(ctx, conn, elem)
	if err != nil {
		return elem, err
	}
	return elems[0], nil
}

// BatchAddMedicine handles insertion of multiple Medicines in the database.
func (d MedicineEntityDAL) BatchAddMedicine(ctx context.Context, conn db.Connection, elems ...types_entity.Medicine) ([]types_entity.Medicine, error) {
	clog.Debugf("AddMedicines: Adding entities...\n%+v", elems)

	// Meta info
	meta := meta_entity.GetMedicineEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return elems, err
	}

	elems, err = dalutil.BatchAddType[types_entity.Medicine, types_entity.MedicineField](ctx, conn, params, meta.BasicTypeDALMeta, elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Medicines", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		return elems, err
	}

	return elems, nil
}

// GetMedicine fetches the entity Medicine based on the ID provided.
func (d MedicineEntityDAL) GetMedicine(ctx context.Context, conn db.Connection, req types_entity.GetMedicineRequest) (types_entity.Medicine, error) {
	id := req.ID
	var elem types_entity.Medicine

	listReq := types_entity.ListMedicineRequest{
		Filter: types_entity.MedicineFilter{
			ID: &filterlib.UUIDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := d.ListMedicine(ctx, conn, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Medicines by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No Medicine found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Medicines found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

func (d MedicineEntityDAL) QueryByTextMedicine(ctx context.Context, conn db.Connection, req types_entity.QueryByTextMedicineRequest) (types_entity.ListMedicineResponse, error) {

	var resp types_entity.ListMedicineResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := d.GetQueryByTextFilterForTypeMedicine(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		clog.Errorf("QueryByTextMedicine: Text search not available")
		return resp, nil
	}

	listReq := types_entity.ListMedicineRequest{
		Filter: *filter,
	}

	listResp, err := d.ListMedicine(ctx, conn, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Medicines by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// ListMedicine fetches a list of Medicine entities based on the given parameters.
func (d MedicineEntityDAL) ListMedicine(ctx context.Context, conn db.Connection, req types_entity.ListMedicineRequest) (types_entity.ListMedicineResponse, error) {
	var resp types_entity.ListMedicineResponse

	// Meta info
	meta := meta_entity.GetMedicineEntityDALMeta()

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique Medicine IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := d.ListMedicineIDs(ctx, conn, req, params)
	if err != nil {
		return resp, err
	}

	clog.Debugf("ListMedicineIDs:\n%v", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[types_entity.Medicine, types_entity.MedicineField](ctx, conn, listTypeByIDsParams, meta.BasicTypeDALMeta)
	if err != nil {
		return resp, err
	}

	resp.Items = typeResp.Items
	resp.Count = len(typeResp.Items)

	// Commit the transaction
	err = conn.Commit(ctx)
	if err != nil {
		return resp, err
	}

	return resp, nil

}

// ListMedicineIDs fetches a list of unique IDs for entity Medicine that match the filter provided.
func (d MedicineEntityDAL) ListMedicineIDs(ctx context.Context, conn db.Connection, req types_entity.ListMedicineRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := d.ListUniqueMedicineIDsByFilter(ctx, conn, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// ListUniqueMedicineIDsByFilter fetches a list of unique IDs for entity Medicine that match the filter provided.
func (d MedicineEntityDAL) ListUniqueMedicineIDsByFilter(ctx context.Context, conn db.Connection, filter types_entity.MedicineFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := d.ListUniqueMedicineIDsByBaseFilter(ctx, conn, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniqueMedicineIDsByFilter(ctx, conn, orFilter, params)
			if err != nil {
				return nil, err
			}
			orIDLists[i] = orFilterIDs
		}
		orIDs = dalutil.GetUUIDsUnion(orIDLists...)
		finalIDs = dalutil.GetUUIDsIntersection(finalIDs, orIDs)
	}

	var andIDs []scalars.ID
	if len(filter.And) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniqueMedicineIDsByFilter(ctx, conn, orFilter, params)
			if err != nil {
				return nil, err
			}
			orIDLists[i] = orFilterIDs
		}
		andIDs = dalutil.GetUUIDsIntersection(orIDLists...)
		finalIDs = dalutil.GetUUIDsIntersection(finalIDs, andIDs)
	}

	return finalIDs, nil
}

// ListUniqueMedicineIDsByBaseFilter fetches a list of unique IDs for entity Medicine that match the filter provided, ignoring the OR and AND filters.
func (d MedicineEntityDAL) ListUniqueMedicineIDsByBaseFilter(ctx context.Context, conn db.Connection, filter types_entity.MedicineFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := ListTypeMedicineRequest{Filter: filter}
	ds, err := d.ConstructListQueryBuilderForTypeMedicine(ctx, conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique Medicine IDs that satisfy the filters: %w", err)
	}

	// Construct the query
	query, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}

	// Run the query
	rows, err := conn.QueryRows(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids, err := db.SqlRowsToUUIDs(ctx, rows)
	if err != nil {
		return nil, fmt.Errorf("Parsing sql.Rows into []ID: %w", err)
	}

	return ids, nil
}

// ListTypeMedicineRequest is the request object for ListTypeMedicine method
type ListTypeMedicineRequest = dalutil.ListTypeRequest[types_entity.MedicineFilter]

// ListTypeMedicineResponse is the response object for ListTypeMedicine method
type ListTypeMedicineResponse = dalutil.ListTypeResponse[types_entity.Medicine]

// ConstructListQueryBuilderForTypeMedicine provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Medicine items that belong to `params.TableName` and match the filter.
func (d MedicineEntityDAL) ConstructListQueryBuilderForTypeMedicine(ctx context.Context, conn db.Connection, req ListTypeMedicineRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetMedicineIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetMedicineMeta()

	ds := goqu.Dialect(conn.Dialect).
		From(params.TableName)

	fullColumnNames := []interface{}{}
	for _, col := range params.SelectColumns {
		fullColumnNames = append(fullColumnNames, params.TableName+"."+col)
	}

	ds = ds.Select(fullColumnNames...)

	// Handle Nested 1:Many tables
	if req.Filter.HavingIngredients != nil {
		subTableName := params.TableName + "_" + "ingredients"
		subReq := ListTypeIngredientRequest{
			Filter: *req.Filter.HavingIngredients,
		}
		subParams := db.UniqueIDsQueryBuilderParams{
			TableName:     subTableName,
			SelectColumns: []string{"id", "parent_id"},
		}
		subQuery, err := d.ConstructListQueryBuilderForTypeIngredient(ctx, conn, subReq, subParams)
		if err != nil {
			return resp, err
		}
		// Add With commands
		for _, with := range subQuery.Withs {
			resp.Withs = append(resp.Withs, with)
		}
		// Add the Select command as with too
		resp.Withs = append(resp.Withs, db.With{Alias: "cte_" + subTableName, Select: subQuery.Select})

		// Add Join to the new With Table
		ds = ds.Join(
			goqu.T("cte_"+subTableName), goqu.On(goqu.I("cte_"+subTableName+".parent_id").Eq(goqu.I(params.TableName+".id"))),
		)
	}

	// Handle Nested objects 1:1

	// Where conditions for the main table (Primitives)
	// TODO: Handle Where conditions for direct SQL column of array types

	// Group By the columns so we don't fetch multiple rows (which can happen if the nested 1:many tables had many rows and we join on it)
	ds = ds.GroupBy(fullColumnNames...)
	if req.Filter.ID != nil {
		// inject a Where filter for ID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ID, ds, "id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Name != nil {
		// inject a Where filter for Name
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Name, ds, "name")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CompanyID != nil {
		// inject a Where filter for CompanyID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CompanyID, ds, "company_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.PrimaryIngredient != nil {
		if req.Filter.PrimaryIngredient.DrugID != nil {
			// inject a Where filter for PrimaryIngredient_DrugID
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PrimaryIngredient.DrugID, ds, "primary_ingredient__drug_id")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.PrimaryIngredient.IsPrimaryIngredient != nil {
			// inject a Where filter for PrimaryIngredient_IsPrimaryIngredient
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PrimaryIngredient.IsPrimaryIngredient, ds, "primary_ingredient__is_primary_ingredient")
			if err != nil {
				return resp, err
			}
		}
	}
	if req.Filter.ModeOfDelivery != nil {
		// inject a Where filter for ModeOfDelivery
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.ModeOfDelivery, ds, "mode_of_delivery")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.CreatedAt != nil {
		// inject a Where filter for CreatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CreatedAt, ds, "created_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.UpdatedAt != nil {
		// inject a Where filter for UpdatedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.UpdatedAt, ds, "updated_at")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.DeletedAt != nil {
		// inject a Where filter for DeletedAt
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.DeletedAt, ds, "deleted_at")
		if err != nil {
			return resp, err
		}
	}

	// In the end, return
	resp.Select = ds

	return resp, nil
}

func (d MedicineEntityDAL) GetQueryByTextFilterForTypeMedicine(ctx context.Context, queryText string) (*types_entity.MedicineFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.MedicineFilter
	{
		filter := types_entity.MedicineFilter{Name: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypeIngredient(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing Ingredient query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, types_entity.MedicineFilter{PrimaryIngredient: &nestedFilter.Or[i]})
			}
		}
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.MedicineFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// UpdateMedicine handles the update of Medicine entity.
func (d MedicineEntityDAL) UpdateMedicine(ctx context.Context, conn db.Connection, req types_entity.UpdateMedicineRequest) (types_entity.UpdateMedicineResponse, error) {

	var resp types_entity.UpdateMedicineResponse

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := meta_entity.GetMedicineEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[types_entity.Medicine, types_entity.MedicineField]{
		TableName:     meta.DbTableName.FormatSQLTable(),
		Object:        req.Object,
		Fields:        req.Fields,
		ExcludeFields: req.ExcludeFields,
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	_, err = dalutil.UpdateType(ctx, conn, subReq, meta.BasicTypeDALMeta)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, fmt.Errorf("Updating Medicine: %w", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	getReq := types_entity.GetMedicineRequest{ID: req.Object.ID}
	entity, err := d.GetMedicine(ctx, conn, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating Medicine: failed to get Medicine entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}
