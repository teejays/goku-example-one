package pharmaceutical_company

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
	meta_entity "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/meta"
	types_entity "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
)

// PharmaceuticalCompanyEntityDAL encapsulates DAL methods for types that fall under PharmaceuticalCompany
type PharmaceuticalCompanyEntityDAL struct {
	dal_service.PharmacyServiceDAL
}

// AddPharmaceuticalCompany handles insertion of a single PharmaceuticalCompany entity in the database.
func (d PharmaceuticalCompanyEntityDAL) AddPharmaceuticalCompany(ctx context.Context, conn db.Connection, elem types_entity.PharmaceuticalCompany) (types_entity.PharmaceuticalCompany, error) {
	elems, err := d.BatchAddPharmaceuticalCompany(ctx, conn, elem)
	if err != nil {
		return elem, err
	}
	return elems[0], nil
}

// BatchAddPharmaceuticalCompany handles insertion of multiple PharmaceuticalCompanies in the database.
func (d PharmaceuticalCompanyEntityDAL) BatchAddPharmaceuticalCompany(ctx context.Context, conn db.Connection, elems ...types_entity.PharmaceuticalCompany) ([]types_entity.PharmaceuticalCompany, error) {
	clog.Debugf("AddPharmaceuticalCompanies: Adding entities...\n%+v", elems)

	// Meta info
	meta := meta_entity.GetPharmaceuticalCompanyEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return elems, err
	}

	elems, err = dalutil.BatchAddType[types_entity.PharmaceuticalCompany, types_entity.PharmaceuticalCompanyField](ctx, conn, params, meta.BasicTypeDALMeta, elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "PharmaceuticalCompanies", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		return elems, err
	}

	return elems, nil
}

// GetPharmaceuticalCompany fetches the entity PharmaceuticalCompany based on the ID provided.
func (d PharmaceuticalCompanyEntityDAL) GetPharmaceuticalCompany(ctx context.Context, conn db.Connection, req types_entity.GetPharmaceuticalCompanyRequest) (types_entity.PharmaceuticalCompany, error) {
	id := req.ID
	var elem types_entity.PharmaceuticalCompany

	listReq := types_entity.ListPharmaceuticalCompanyRequest{
		Filter: types_entity.PharmaceuticalCompanyFilter{
			ID: &filterlib.UUIDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := d.ListPharmaceuticalCompany(ctx, conn, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list PharmaceuticalCompanies by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No PharmaceuticalCompany found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d PharmaceuticalCompanies found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

func (d PharmaceuticalCompanyEntityDAL) QueryByTextPharmaceuticalCompany(ctx context.Context, conn db.Connection, req types_entity.QueryByTextPharmaceuticalCompanyRequest) (types_entity.ListPharmaceuticalCompanyResponse, error) {

	var resp types_entity.ListPharmaceuticalCompanyResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := d.GetQueryByTextFilterForTypePharmaceuticalCompany(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		clog.Errorf("QueryByTextPharmaceuticalCompany: Text search not available")
		return resp, nil
	}

	listReq := types_entity.ListPharmaceuticalCompanyRequest{
		Filter: *filter,
	}

	listResp, err := d.ListPharmaceuticalCompany(ctx, conn, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query PharmaceuticalCompanies by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// ListPharmaceuticalCompany fetches a list of PharmaceuticalCompany entities based on the given parameters.
func (d PharmaceuticalCompanyEntityDAL) ListPharmaceuticalCompany(ctx context.Context, conn db.Connection, req types_entity.ListPharmaceuticalCompanyRequest) (types_entity.ListPharmaceuticalCompanyResponse, error) {
	var resp types_entity.ListPharmaceuticalCompanyResponse

	// Meta info
	meta := meta_entity.GetPharmaceuticalCompanyEntityDALMeta()

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique PharmaceuticalCompany IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := d.ListPharmaceuticalCompanyIDs(ctx, conn, req, params)
	if err != nil {
		return resp, err
	}

	clog.Debugf("ListPharmaceuticalCompanyIDs:\n%v", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[types_entity.PharmaceuticalCompany, types_entity.PharmaceuticalCompanyField](ctx, conn, listTypeByIDsParams, meta.BasicTypeDALMeta)
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

// ListPharmaceuticalCompanyIDs fetches a list of unique IDs for entity PharmaceuticalCompany that match the filter provided.
func (d PharmaceuticalCompanyEntityDAL) ListPharmaceuticalCompanyIDs(ctx context.Context, conn db.Connection, req types_entity.ListPharmaceuticalCompanyRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := d.ListUniquePharmaceuticalCompanyIDsByFilter(ctx, conn, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// ListUniquePharmaceuticalCompanyIDsByFilter fetches a list of unique IDs for entity PharmaceuticalCompany that match the filter provided.
func (d PharmaceuticalCompanyEntityDAL) ListUniquePharmaceuticalCompanyIDsByFilter(ctx context.Context, conn db.Connection, filter types_entity.PharmaceuticalCompanyFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := d.ListUniquePharmaceuticalCompanyIDsByBaseFilter(ctx, conn, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniquePharmaceuticalCompanyIDsByFilter(ctx, conn, orFilter, params)
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
			orFilterIDs, err := d.ListUniquePharmaceuticalCompanyIDsByFilter(ctx, conn, orFilter, params)
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

// ListUniquePharmaceuticalCompanyIDsByBaseFilter fetches a list of unique IDs for entity PharmaceuticalCompany that match the filter provided, ignoring the OR and AND filters.
func (d PharmaceuticalCompanyEntityDAL) ListUniquePharmaceuticalCompanyIDsByBaseFilter(ctx context.Context, conn db.Connection, filter types_entity.PharmaceuticalCompanyFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := ListTypePharmaceuticalCompanyRequest{Filter: filter}
	ds, err := d.ConstructListQueryBuilderForTypePharmaceuticalCompany(ctx, conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique PharmaceuticalCompany IDs that satisfy the filters: %w", err)
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

// ListTypePharmaceuticalCompanyRequest is the request object for ListTypePharmaceuticalCompany method
type ListTypePharmaceuticalCompanyRequest = dalutil.ListTypeRequest[types_entity.PharmaceuticalCompanyFilter]

// ListTypePharmaceuticalCompanyResponse is the response object for ListTypePharmaceuticalCompany method
type ListTypePharmaceuticalCompanyResponse = dalutil.ListTypeResponse[types_entity.PharmaceuticalCompany]

// ConstructListQueryBuilderForTypePharmaceuticalCompany provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all PharmaceuticalCompany items that belong to `params.TableName` and match the filter.
func (d PharmaceuticalCompanyEntityDAL) ConstructListQueryBuilderForTypePharmaceuticalCompany(ctx context.Context, conn db.Connection, req ListTypePharmaceuticalCompanyRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetPharmaceuticalCompanyIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetPharmaceuticalCompanyMeta()

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

func (d PharmaceuticalCompanyEntityDAL) GetQueryByTextFilterForTypePharmaceuticalCompany(ctx context.Context, queryText string) (*types_entity.PharmaceuticalCompanyFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.PharmaceuticalCompanyFilter
	{
		filter := types_entity.PharmaceuticalCompanyFilter{Name: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.PharmaceuticalCompanyFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// UpdatePharmaceuticalCompany handles the update of PharmaceuticalCompany entity.
func (d PharmaceuticalCompanyEntityDAL) UpdatePharmaceuticalCompany(ctx context.Context, conn db.Connection, req types_entity.UpdatePharmaceuticalCompanyRequest) (types_entity.UpdatePharmaceuticalCompanyResponse, error) {

	var resp types_entity.UpdatePharmaceuticalCompanyResponse

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := meta_entity.GetPharmaceuticalCompanyEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[types_entity.PharmaceuticalCompany, types_entity.PharmaceuticalCompanyField]{
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
		return resp, fmt.Errorf("Updating PharmaceuticalCompany: %w", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	getReq := types_entity.GetPharmaceuticalCompanyRequest{ID: req.Object.ID}
	entity, err := d.GetPharmaceuticalCompany(ctx, conn, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating PharmaceuticalCompany: failed to get PharmaceuticalCompany entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}
