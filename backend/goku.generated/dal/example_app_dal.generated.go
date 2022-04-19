package example_app

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // required for 'postgres' dialect

	"github.com/teejays/clog"
	"github.com/teejays/goku-util/client/db"
	dalutil "github.com/teejays/goku-util/dalutil"
	filterlib "github.com/teejays/goku-util/filter"

	types_global "github.com/teejays/goku-example-one/backend/goku.generated/types"
)

// ExampleAppAppDAL encapsulates DAL methods for types that fall under ExampleApp
type ExampleAppAppDAL struct {
}

// ListTypeAddressRequest is the request object for ListTypeAddress method
type ListTypeAddressRequest = dalutil.ListTypeRequest[types_global.AddressFilter]

// ListTypeAddressResponse is the response object for ListTypeAddress method
type ListTypeAddressResponse = dalutil.ListTypeResponse[types_global.AddressWithMeta]

// ConstructListQueryBuilderForTypeAddress provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Address items that belong to `params.TableName` and match the filter.
func (d ExampleAppAppDAL) ConstructListQueryBuilderForTypeAddress(ctx context.Context, conn db.Connection, req ListTypeAddressRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetAddressIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_global.GetAddressMeta()

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
	if req.Filter.Line1 != nil {
		// inject a Where filter for Line1
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Line1, ds, "line_1")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Line2 != nil {
		// inject a Where filter for Line2
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Line2, ds, "line_2")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.City != nil {
		// inject a Where filter for City
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.City, ds, "city")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Province != nil {
		// inject a Where filter for Province
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Province, ds, "province")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.PostalCode != nil {
		// inject a Where filter for PostalCode
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.PostalCode, ds, "postal_code")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Country != nil {
		// inject a Where filter for Country
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Country, ds, "country")
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

func (d ExampleAppAppDAL) GetQueryByTextFilterForTypeAddress(ctx context.Context, queryText string) (*types_global.AddressFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_global.AddressFilter
	{
		filter := types_global.AddressFilter{Line1: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.AddressFilter{Line2: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.AddressFilter{City: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.AddressFilter{PostalCode: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_global.AddressFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ListTypeContactRequest is the request object for ListTypeContact method
type ListTypeContactRequest = dalutil.ListTypeRequest[types_global.ContactFilter]

// ListTypeContactResponse is the response object for ListTypeContact method
type ListTypeContactResponse = dalutil.ListTypeResponse[types_global.ContactWithMeta]

// ConstructListQueryBuilderForTypeContact provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all Contact items that belong to `params.TableName` and match the filter.
func (d ExampleAppAppDAL) ConstructListQueryBuilderForTypeContact(ctx context.Context, conn db.Connection, req ListTypeContactRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetContactIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_global.GetContactMeta()

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
	if req.Filter.Address != nil {
		if req.Filter.Address.Line1 != nil {
			// inject a Where filter for Address_Line1
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.Line1, ds, "address__line_1")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Address.Line2 != nil {
			// inject a Where filter for Address_Line2
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.Line2, ds, "address__line_2")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Address.City != nil {
			// inject a Where filter for Address_City
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.City, ds, "address__city")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Address.Province != nil {
			// inject a Where filter for Address_Province
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.Province, ds, "address__province")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Address.PostalCode != nil {
			// inject a Where filter for Address_PostalCode
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.PostalCode, ds, "address__postal_code")
			if err != nil {
				return resp, err
			}
		}
		if req.Filter.Address.Country != nil {
			// inject a Where filter for Address_Country
			ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Address.Country, ds, "address__country")
			if err != nil {
				return resp, err
			}
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

func (d ExampleAppAppDAL) GetQueryByTextFilterForTypeContact(ctx context.Context, queryText string) (*types_global.ContactFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_global.ContactFilter
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypePersonName(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing PersonName query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, types_global.ContactFilter{Name: &nestedFilter.Or[i]})
			}
		}
	}
	{
		nestedFilter, err := d.GetQueryByTextFilterForTypeAddress(ctx, queryText)
		if err != nil {
			return nil, fmt.Errorf("constructing Address query text filter: %w", err)
		}
		if nestedFilter != nil {
			for i := range nestedFilter.Or {
				orFilters = append(orFilters, types_global.ContactFilter{Address: &nestedFilter.Or[i]})
			}
		}
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_global.ContactFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ListTypePersonNameRequest is the request object for ListTypePersonName method
type ListTypePersonNameRequest = dalutil.ListTypeRequest[types_global.PersonNameFilter]

// ListTypePersonNameResponse is the response object for ListTypePersonName method
type ListTypePersonNameResponse = dalutil.ListTypeResponse[types_global.PersonNameWithMeta]

// ConstructListQueryBuilderForTypePersonName provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all PersonName items that belong to `params.TableName` and match the filter.
func (d ExampleAppAppDAL) ConstructListQueryBuilderForTypePersonName(ctx context.Context, conn db.Connection, req ListTypePersonNameRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetPersonNameIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_global.GetPersonNameMeta()

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
	if req.Filter.First != nil {
		// inject a Where filter for First
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.First, ds, "first")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.MiddleInitial != nil {
		// inject a Where filter for MiddleInitial
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.MiddleInitial, ds, "middle_initial")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Last != nil {
		// inject a Where filter for Last
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Last, ds, "last")
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

func (d ExampleAppAppDAL) GetQueryByTextFilterForTypePersonName(ctx context.Context, queryText string) (*types_global.PersonNameFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_global.PersonNameFilter
	{
		filter := types_global.PersonNameFilter{First: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.PersonNameFilter{MiddleInitial: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.PersonNameFilter{Last: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_global.PersonNameFilter{
		Or: orFilters,
	}

	return &filter, nil
}

// ListTypePhoneNumberRequest is the request object for ListTypePhoneNumber method
type ListTypePhoneNumberRequest = dalutil.ListTypeRequest[types_global.PhoneNumberFilter]

// ListTypePhoneNumberResponse is the response object for ListTypePhoneNumber method
type ListTypePhoneNumberResponse = dalutil.ListTypeResponse[types_global.PhoneNumberWithMeta]

// ConstructListQueryBuilderForTypePhoneNumber provides a query builder (goqu.SelectDataset) representing a query that can
// be used to get the IDs of all PhoneNumber items that belong to `params.TableName` and match the filter.
func (d ExampleAppAppDAL) ConstructListQueryBuilderForTypePhoneNumber(ctx context.Context, conn db.Connection, req ListTypePhoneNumberRequest, params db.UniqueIDsQueryBuilderParams) (db.SelectQueryBuilder, error) {
	clog.Debugf("GetPhoneNumberIDsSelectQueryBuilder, with Request\n%+v", req)
	var resp db.SelectQueryBuilder
	var err error

	// Get Meta
	// meta := meta_global.GetPhoneNumberMeta()

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
	if req.Filter.CountryCode != nil {
		// inject a Where filter for CountryCode
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.CountryCode, ds, "country_code")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Number != nil {
		// inject a Where filter for Number
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Number, ds, "number")
		if err != nil {
			return resp, err
		}
	}
	if req.Filter.Extension != nil {
		// inject a Where filter for Extension
		ds, err = filterlib.InjectConditionIntoSqlBuilder(req.Filter.Extension, ds, "extension")
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

func (d ExampleAppAppDAL) GetQueryByTextFilterForTypePhoneNumber(ctx context.Context, queryText string) (*types_global.PhoneNumberFilter, error) {

	// Add OR filters for each field that can be searched over text
	var orFilters []types_global.PhoneNumberFilter
	{
		filter := types_global.PhoneNumberFilter{Number: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}
	{
		filter := types_global.PhoneNumberFilter{Extension: &filterlib.StringCondition{Op: filterlib.ILIKE, Values: []string{queryText}}}
		orFilters = append(orFilters, filter)
	}

	if len(orFilters) < 1 {
		return nil, nil
	}

	filter := types_global.PhoneNumberFilter{
		Or: orFilters,
	}

	return &filter, nil
}
