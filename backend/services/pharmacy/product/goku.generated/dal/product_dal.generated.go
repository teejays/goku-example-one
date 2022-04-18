package product

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

	dal_service "github.com/teejays/goku/example/backend/services/pharmacy/goku.generated/dal"
	meta_entity "github.com/teejays/goku/example/backend/services/pharmacy/product/goku.generated/meta"
	types_entity "github.com/teejays/goku/example/backend/services/pharmacy/product/goku.generated/types"
)

// ProductEntityDAL encapsulates DAL methods for types that fall under Product
type ProductEntityDAL struct {
	dal_service.PharmacyServiceDAL
}

// AddProduct handles insertion of a single Product entity in the database.
func (d ProductEntityDAL) AddProduct(ctx context.Context, conn db.Connection, elem types_entity.Product) (types_entity.Product, error) {
	elems, err := d.BatchAddProduct(ctx, conn, elem)
	if err != nil {
		return elem, err
	}
	return elems[0], nil
}

// BatchAddProduct handles insertion of multiple Products in the database.
func (d ProductEntityDAL) BatchAddProduct(ctx context.Context, conn db.Connection, elems ...types_entity.Product) ([]types_entity.Product, error) {
	clog.Debugf("AddProducts: Adding entities...\n%+v", elems)

	// Meta info
	meta := meta_entity.GetProductEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return elems, err
	}

	elems, err = dalutil.BatchAddType[types_entity.Product, types_entity.ProductField](ctx, conn, params, meta.BasicTypeDALMeta, elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Products", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		return elems, err
	}

	return elems, nil
}

// GetProduct fetches the entity Product based on the ID provided.
func (d ProductEntityDAL) GetProduct(ctx context.Context, conn db.Connection, req types_entity.GetProductRequest) (types_entity.Product, error) {
	id := req.ID
	var elem types_entity.Product

	listReq := types_entity.ListProductRequest{
		Filter: types_entity.ProductFilter{
			ID: &filterlib.UUIDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := d.ListProduct(ctx, conn, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Products by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No Product found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Products found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

func (d ProductEntityDAL) QueryByTextProduct(ctx context.Context, conn db.Connection, req types_entity.QueryByTextProductRequest) (types_entity.ListProductResponse, error) {

	var resp types_entity.ListProductResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := d.GetQueryByTextFilterForTypeProduct(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		clog.Errorf("QueryByTextProduct: Text search not available")
		return resp, nil
	}

	listReq := types_entity.ListProductRequest{
		Filter: *filter,
	}

	listResp, err := d.ListProduct(ctx, conn, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Products by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// ListProduct fetches a list of Product entities based on the given parameters.
func (d ProductEntityDAL) ListProduct(ctx context.Context, conn db.Connection, req types_entity.ListProductRequest) (types_entity.ListProductResponse, error) {
	var resp types_entity.ListProductResponse

	// Meta info
	meta := meta_entity.GetProductEntityDALMeta()

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique Product IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := d.ListProductIDs(ctx, conn, req, params)
	if err != nil {
		return resp, err
	}

	clog.Debugf("ListProductIDs:\n%v", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[types_entity.Product, types_entity.ProductField](ctx, conn, listTypeByIDsParams, meta.BasicTypeDALMeta)
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

// ListProductIDs fetches a list of unique IDs for entity Product that match the filter provided.
func (d ProductEntityDAL) ListProductIDs(ctx context.Context, conn db.Connection, req types_entity.ListProductRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := d.ListUniqueProductIDsByFilter(ctx, conn, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// ListUniqueProductIDsByFilter fetches a list of unique IDs for entity Product that match the filter provided.
func (d ProductEntityDAL) ListUniqueProductIDsByFilter(ctx context.Context, conn db.Connection, filter types_entity.ProductFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := d.ListUniqueProductIDsByBaseFilter(ctx, conn, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniqueProductIDsByFilter(ctx, conn, orFilter, params)
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
			orFilterIDs, err := d.ListUniqueProductIDsByFilter(ctx, conn, orFilter, params)
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

// ListUniqueProductIDsByBaseFilter fetches a list of unique IDs for entity Product that match the filter provided, ignoring the OR and AND filters.
func (d ProductEntityDAL) ListUniqueProductIDsByBaseFilter(ctx context.Context, conn db.Connection, filter types_entity.ProductFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := ListTypeProductRequest{Filter: filter}
	ds, err := d.ConstructListQueryBuilderForTypeProduct(ctx, conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique Product IDs that satisfy the filters: %w", err)
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

// ListTypeProductRequest is the request object for ListTypeProduct method
type ListTypeProductRequest = dalutil.ListTypeRequest[types_entity.ProductFilter]

// ListTypeProductResponse is the response object for ListTypeProduct method
type ListTypeProductResponse = dalutil.ListTypeResponse[types_entity.Product]

// ConstructListQueryBuilderForTypeProduct provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Product items that belong to `params.TableName` and match the filter.
func (d ProductEntityDAL) ConstructListQueryBuilderForTypeProduct(ctx context.Context, conn db.Connection, req ListTypeProductRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetProductIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetProductMeta()

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
	if req.Filter.MedicineID != nil {
		// inject a Where filter for MedicineID
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.MedicineID, ds, "medicine_id")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Mass != nil {
		// inject a Where filter for Mass
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Mass, ds, "mass")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Count != nil {
		// inject a Where filter for Count
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Count, ds, "count")
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

func (d ProductEntityDAL) GetQueryByTextFilterForTypeProduct(ctx context.Context, queryText string) (*types_entity.ProductFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.ProductFilter
	{
		filter := types_entity.ProductFilter{Name: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.ProductFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// UpdateProduct handles the update of Product entity.
func (d ProductEntityDAL) UpdateProduct(ctx context.Context, conn db.Connection, req types_entity.UpdateProductRequest) (types_entity.UpdateProductResponse, error) {

	var resp types_entity.UpdateProductResponse

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := meta_entity.GetProductEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[types_entity.Product, types_entity.ProductField]{
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
		return resp, fmt.Errorf("Updating Product: %w", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	getReq := types_entity.GetProductRequest{ID: req.Object.ID}
	entity, err := d.GetProduct(ctx, conn, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating Product: failed to get Product entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}
