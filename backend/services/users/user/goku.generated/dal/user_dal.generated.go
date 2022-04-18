package user

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

	dal_service "github.com/teejays/goku/example/backend/services/users/goku.generated/dal"
	meta_entity "github.com/teejays/goku/example/backend/services/users/user/goku.generated/meta"
	types_entity "github.com/teejays/goku/example/backend/services/users/user/goku.generated/types"
)

// UserEntityDAL encapsulates DAL methods for types that fall under User
type UserEntityDAL struct {
	dal_service.UsersServiceDAL
}

// AddUser handles insertion of a single User entity in the database.
func (d UserEntityDAL) AddUser(ctx context.Context, conn db.Connection, elem types_entity.User) (types_entity.User, error) {
	elems, err := d.BatchAddUser(ctx, conn, elem)
	if err != nil {
		return elem, err
	}
	return elems[0], nil
}

// BatchAddUser handles insertion of multiple Users in the database.
func (d UserEntityDAL) BatchAddUser(ctx context.Context, conn db.Connection, elems ...types_entity.User) ([]types_entity.User, error) {
	clog.Debugf("AddUsers: Adding entities...\n%+v", elems)

	// Meta info
	meta := meta_entity.GetUserEntityDALMeta()

	params := db.InsertTypeParams{
		TableName: meta.DbTableName.FormatSQL(),
	}

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return elems, err
	}

	elems, err = dalutil.BatchAddType[types_entity.User, types_entity.UserField](ctx, conn, params, meta.BasicTypeDALMeta, elems...)
	if err != nil {
		return nil, fmt.Errorf("Inserting type %s: %w", "Users", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		return elems, err
	}

	return elems, nil
}

// GetUser fetches the entity User based on the ID provided.
func (d UserEntityDAL) GetUser(ctx context.Context, conn db.Connection, req types_entity.GetUserRequest) (types_entity.User, error) {
	id := req.ID
	var elem types_entity.User

	listReq := types_entity.ListUserRequest{
		Filter: types_entity.UserFilter{
			ID: &filterlib.UUIDCondition{Op: filterlib.EQUAL, Values: []scalars.ID{id}},
		},
	}

	listResp, err := d.ListUser(ctx, conn, listReq)
	if err != nil {
		return elem, fmt.Errorf("Could not list Users by ID: %w", err)
	}
	if len(listResp.Items) < 1 {
		return elem, fmt.Errorf("No User found with ID %s", id)
	}
	panics.If(len(listResp.Items) > 1, "Expected 1 item but found %d Users found with ID %s", len(listResp.Items), id)

	return listResp.Items[0], nil

}

func (d UserEntityDAL) QueryByTextUser(ctx context.Context, conn db.Connection, req types_entity.QueryByTextUserRequest) (types_entity.ListUserResponse, error) {

	var resp types_entity.ListUserResponse

	queryText := req.QueryText

	// Prepare filter
	filter, err := d.GetQueryByTextFilterForTypeUser(ctx, queryText)
	if err != nil {
		return resp, err
	}

	// If no empty filter returned, means we have no way to search by text on this entity
	if filter == nil {
		clog.Errorf("QueryByTextUser: Text search not available")
		return resp, nil
	}

	listReq := types_entity.ListUserRequest{
		Filter: *filter,
	}

	listResp, err := d.ListUser(ctx, conn, listReq)
	if err != nil {
		return listResp, fmt.Errorf("Could not query Users by text `%s`: %w", req.QueryText, err)
	}

	return listResp, nil
}

// ListUser fetches a list of User entities based on the given parameters.
func (d UserEntityDAL) ListUser(ctx context.Context, conn db.Connection, req types_entity.ListUserRequest) (types_entity.ListUserResponse, error) {
	var resp types_entity.ListUserResponse

	// Meta info
	meta := meta_entity.GetUserEntityDALMeta()

	// Begin a Transaction
	err := conn.Begin(ctx)
	if err != nil {
		return resp, err
	}

	// Get all the unique User IDs that we should fetch (based on the filters)
	params := db.UniqueIDsQueryBuilderParams{
		TableName:     meta.DbTableName.FormatSQL(),
		SelectColumns: []string{"id"},
	}

	ids, err := d.ListUserIDs(ctx, conn, req, params)
	if err != nil {
		return resp, err
	}

	clog.Debugf("ListUserIDs:\n%v", ids)

	// Use the IDs to fetch the elements
	listTypeByIDsParams := db.ListTypeByIDsParams{
		TableName: meta.DbTableName.FormatSQL(),
		IDColumn:  "id",
		IDs:       ids,
	}
	typeResp, err := dalutil.ListTypeByIDs[types_entity.User, types_entity.UserField](ctx, conn, listTypeByIDsParams, meta.BasicTypeDALMeta)
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

// ListUserIDs fetches a list of unique IDs for entity User that match the filter provided.
func (d UserEntityDAL) ListUserIDs(ctx context.Context, conn db.Connection, req types_entity.ListUserRequest, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	ids, err := d.ListUniqueUserIDsByFilter(ctx, conn, req.Filter, params)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// ListUniqueUserIDsByFilter fetches a list of unique IDs for entity User that match the filter provided.
func (d UserEntityDAL) ListUniqueUserIDsByFilter(ctx context.Context, conn db.Connection, filter types_entity.UserFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// By default, the baseIds, orIDs, and andIDs should be ANDed together i.e. only the IDs that satisfy everything should be returned
	var finalIDs []scalars.ID

	baseIDs, err := d.ListUniqueUserIDsByBaseFilter(ctx, conn, filter, params)
	if err != nil {
		return nil, err
	}
	finalIDs = baseIDs

	var orIDs []scalars.ID
	if len(filter.Or) > 0 {
		var orIDLists = make([][]scalars.ID, len(filter.Or))
		for i, orFilter := range filter.Or {
			orFilterIDs, err := d.ListUniqueUserIDsByFilter(ctx, conn, orFilter, params)
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
			orFilterIDs, err := d.ListUniqueUserIDsByFilter(ctx, conn, orFilter, params)
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

// ListUniqueUserIDsByBaseFilter fetches a list of unique IDs for entity User that match the filter provided, ignoring the OR and AND filters.
func (d UserEntityDAL) ListUniqueUserIDsByBaseFilter(ctx context.Context, conn db.Connection, filter types_entity.UserFilter, params db.UniqueIDsQueryBuilderParams) ([]scalars.ID, error) {

	// Since we do not look at Or/And in this function, nil them out
	filter.Or = nil
	filter.And = nil

	// Fetch the Query builder
	subReq := ListTypeUserRequest{Filter: filter}
	ds, err := d.ConstructListQueryBuilderForTypeUser(ctx, conn, subReq, params)
	if err != nil {
		return nil, fmt.Errorf("Constructing query to fetch unique User IDs that satisfy the filters: %w", err)
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

// ListTypeUserRequest is the request object for ListTypeUser method
type ListTypeUserRequest = dalutil.ListTypeRequest[types_entity.UserFilter]

// ListTypeUserResponse is the response object for ListTypeUser method
type ListTypeUserResponse = dalutil.ListTypeResponse[types_entity.User]

// ConstructListQueryBuilderForTypeUser provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all User items that belong to `params.TableName` and match the filter.
func (d UserEntityDAL) ConstructListQueryBuilderForTypeUser(ctx context.Context, conn db.Connection, req ListTypeUserRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetUserIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_entity.GetUserMeta()

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
		if req.Filter.Name.First != nil {
			// inject a Where filter for Name_First
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Name.First, ds, "name__first")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Name.MiddleInitial != nil {
			// inject a Where filter for Name_MiddleInitial
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Name.MiddleInitial, ds, "name__middle_initial")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Name.Last != nil {
			// inject a Where filter for Name_Last
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Name.Last, ds, "name__last")
			if err != nil {
				return resp, err
			}
		}
	}
	if req.Filter.Email != nil {
		// inject a Where filter for Email
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Email, ds, "email")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.PhoneNumber != nil {
		if req.Filter.PhoneNumber.CountryCode != nil {
			// inject a Where filter for PhoneNumber_CountryCode
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PhoneNumber.CountryCode, ds, "phone_number__country_code")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.PhoneNumber.Number != nil {
			// inject a Where filter for PhoneNumber_Number
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PhoneNumber.Number, ds, "phone_number__number")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.PhoneNumber.Extension != nil {
			// inject a Where filter for PhoneNumber_Extension
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PhoneNumber.Extension, ds, "phone_number__extension")
			if err != nil {
				return resp, err
			}
		}
	}
	if req.Filter.PasswordHash != nil {
		// inject a Where filter for PasswordHash
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PasswordHash, ds, "password_hash")
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

func (d UserEntityDAL) GetQueryByTextFilterForTypeUser(ctx context.Context, queryText string) (*types_entity.UserFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_entity.UserFilter
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypePersonName(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing PersonName query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, types_entity.UserFilter{Name: &nestedFilter.Or[i]})
			}
		}
	}
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypePhoneNumber(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing PhoneNumber query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, types_entity.UserFilter{PhoneNumber: &nestedFilter.Or[i]})
			}
		}
	}
	{
		filter := types_entity.UserFilter{PasswordHash: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_entity.UserFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// UpdateUser handles the update of User entity.
func (d UserEntityDAL) UpdateUser(ctx context.Context, conn db.Connection, req types_entity.UpdateUserRequest) (types_entity.UpdateUserResponse, error) {

	var resp types_entity.UpdateUserResponse

	if req.Object.ID.IsEmpty() {
		return resp, fmt.Errorf("entity to be updated has an empty ID")
	}

	// Meta info
	meta := meta_entity.GetUserEntityDALMeta()

	subReq := dalutil.UpdateTypeRequest[types_entity.User, types_entity.UserField]{
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
		return resp, fmt.Errorf("Updating User: %w", err)
	}

	err = conn.Commit(ctx)
	if err != nil {
		conn.MustRollback(ctx)
		return resp, err
	}

	// Get the entity from the DB again. We cannot use `subResp.Object` since that is based on the req.Object, which may not
	// reflect the actual state of the entity since the req.Object may contain updated/missing fields that are not in the Field mask.
	getReq := types_entity.GetUserRequest{ID: req.Object.ID}
	entity, err := d.GetUser(ctx, conn, getReq)
	if err != nil {
		return resp, fmt.Errorf("Updating User: failed to get User entity after update: %w", err)
	}

	resp.Object = entity

	return resp, nil

}
