package drug

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect

	"github.com/teejays/clog"
	"github.com/teejays/goku/generator/external/client/db"
	dalutil "github.com/teejays/goku/generator/external/dalutil"
	filterlib "github.com/teejays/goku/generator/external/filter"
	"github.com/teejays/goku/generator/lib/panics"

	scalars "github.com/teejays/goku/generator/external/scalars"

	meta_entity "github.com/teejays/goku/example/backend/services/pharmacy/drug/goku.generated/meta"
	types_entity "github.com/teejays/goku/example/backend/services/pharmacy/drug/goku.generated/types"
	dal_service "github.com/teejays/goku/example/backend/services/pharmacy/goku.generated/dal"
)

// DrugEntityDAL encapsulates DAL methods for types that fall under Drug
type DrugEntityDAL struct {
	dal_service.PharmacyServiceDAL
}

// AddDrug handles insertion of a single Drug entity in the database.
func (d DrugEntityDAL) AddDrug(ctx context.Context, conn db.Connection, elem types_entity.Drug) (types_entity.Drug, error) {
	elems, err := d.BatchAddDrug(ctx, conn, elem)
	if err != nil {
		return elem, err
	}
	return elems[0], nil
}

// BatchAddDrug handles insertion of multiple Drugs in the database.
func (d DrugEntityDAL) BatchAddDrug(ctx context.Context, conn db.Connection, elems ...types_entity.Drug) ([]types_entity.Drug, error) {
	clog.Debugf("AddDrugs: Adding entities...\n%+v", elems)

	// Meta info
	meta := meta_entity.GetDrugEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return elems, err
	}

	elems, err = dalutil.BatchAddType[types_entity.Drug, types_entity.DrugField](ctx, conn, params, meta.BasicTypeDALMeta, elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Drugs", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		return elems, err
	}

	return elems, nil
}

// GetDrug fetches the entity Drug based on the ID provided.
func (d DrugEntityDAL) GetDrug(ctx context.Context, conn db.Connection, req types_entity.GetDrugRequest) (types_entity.Drug, error) {
	id := req.ID
	var elem types_entity.Drug

	listReq := types_entity.ListDrugRequest{
		Filter: types_entity.DrugFilter{
			ID: &filterlib.UUIDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := d.ListDrug(ctx, conn, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Drugs by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No Drug found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Drugs found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

func (d DrugEntityDAL) QueryByTextDrug(ctx context.Context, conn db.Connection, req types_entity.QueryByTextDrugRequest) (types_entity.ListDrugResponse, error) {

	var resp types_entity.ListDrugResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := d.GetQueryByTextFilterForTypeDrug(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		clog.Errorf("QueryByTextDrug: Text search not available")
		return resp, nil
	}

	listReq := types_entity.ListDrugRequest{
		Filter: *filter,
	}

	listResp, err := d.ListDrug(ctx, conn, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Drugs by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// ListDrug fetches a list of Drug entities based on the given parameters.
func (d DrugEntityDAL) ListDrug(ctx context.Context, conn db.Connection, req types_entity.ListDrugRequest) (types_entity.ListDrugResponse, error) {
	var resp types_entity.ListDrugResponse

	// Meta info
	meta := meta_entity.GetDrugEntityDALMeta()

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique Drug IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := d.ListDrugIDs(ctx, conn, req, params)
	if err != nil {
		return resp, err
	}

	clog.Debugf("ListDrugIDs:\n%v", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[types_entity.Drug, types_entity.DrugField](ctx, conn, listTypeByIDsParams, meta.BasicTypeDALMeta)
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

// ListDrugIDs fetches a list of unique IDs for entity Drug that match the filter provided.
func (d DrugEntityDAL) ListDrugIDs(ctx context.Context, conn db.Connection, req types_entity.ListDrugRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := d.ListUniqueDrugIDsByFilter(ctx, conn, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// ListUniqueDrugIDsByFilter fetches a list of unique IDs for entity Drug that match the filter provided.
func (d DrugEntityDAL) ListUniqueDrugIDsByFilter(ctx context.Context, conn db.Connection, filter types_entity.DrugFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := d.ListUniqueDrugIDsByBaseFilter(ctx, conn, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniqueDrugIDsByFilter(ctx, conn, orFilter, params)
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
			orFilterIDs, err := d.ListUniqueDrugIDsByFilter(ctx, conn, orFilter, params)
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

// ListUniqueDrugIDsByBaseFilter fetches a list of unique IDs for entity Drug that match the filter provided, ignoring the OR and AND filters.
func (d DrugEntityDAL) ListUniqueDrugIDsByBaseFilter(ctx context.Context, conn db.Connection, filter types_entity.DrugFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := ListTypeDrugRequest{Filter: filter}
	ds, err := d.ConstructListQueryBuilderForTypeDrug(ctx, conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique Drug IDs that satisfy the filters: %w", err)
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

// ListTypeDrugRequest is the request object for ListTypeDrug method
type ListTypeDrugRequest = dalutil.ListTypeRequest[types_entity.DrugFilter]

// ListTypeDrugResponse is the response object for ListTypeDrug method
type ListTypeDrugResponse = dalutil.ListTypeResponse[types_entity.Drug]

// ConstructListQueryBuilderForTypeDrug provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Drug items that belong to `params.TableName` and match the filter.
func (d DrugEntityDAL) ConstructListQueryBuilderForTypeDrug(ctx context.Context, conn db.Connection, req ListTypeDrugRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetDrugIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetDrugMeta()

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

func (d DrugEntityDAL) GetQueryByTextFilterForTypeDrug(ctx context.Context, queryText string) (*types_entity.DrugFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.DrugFilter
	{
		filter := types_entity.DrugFilter{Name: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.DrugFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// UpdateDrug handles the update of Drug entity.
func (d DrugEntityDAL) UpdateDrug(ctx context.Context, conn db.Connection, req types_entity.UpdateDrugRequest) (types_entity.UpdateDrugResponse, error) {

	var resp types_entity.UpdateDrugResponse

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := meta_entity.GetDrugEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[types_entity.Drug, types_entity.DrugField]{
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
		return resp, fmt.Errorf("Updating Drug: %w", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	getReq := types_entity.GetDrugRequest{ID: req.Object.ID}
	entity, err := d.GetDrug(ctx, conn, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating Drug: failed to get Drug entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}
