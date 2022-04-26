package http_handlers_pharmacy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teejays/clog"
	gopi "github.com/teejays/gopi"

	"github.com/teejays/goku-util/errutil"
	"github.com/teejays/goku-util/httputil"

	drug_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/methods"
	drug_types "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/types"
	medicine_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/methods"
	medicine_types "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/types"
	pharmaceutical_company_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/methods"
	pharmaceutical_company_types "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
	product_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/methods"
	product_types "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/types"
)

// GetRoutes returns all the routes for this namespace
func GetRoutes() []gopi.Route {

	routes := []gopi.Route{
		{
			// API Route for POST pharmacy/drug
			Method:       "POST",
			Version:      1,
			Path:         "pharmacy/drug",
			HandlerFunc:  AddDrugHandler,
			Authenticate: true,
		},
		{
			// API Route for PUT pharmacy/drug
			Method:       "PUT",
			Version:      1,
			Path:         "pharmacy/drug",
			HandlerFunc:  UpdateDrugHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/drug
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/drug",
			HandlerFunc:  GetDrugHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/drug/list
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/drug/list",
			HandlerFunc:  ListDrugHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/drug/query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/drug/query_by_text",
			HandlerFunc:  QueryByTextDrugHandler,
			Authenticate: true,
		},
		{
			// API Route for POST pharmacy/medicine
			Method:       "POST",
			Version:      1,
			Path:         "pharmacy/medicine",
			HandlerFunc:  AddMedicineHandler,
			Authenticate: true,
		},
		{
			// API Route for PUT pharmacy/medicine
			Method:       "PUT",
			Version:      1,
			Path:         "pharmacy/medicine",
			HandlerFunc:  UpdateMedicineHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/medicine
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/medicine",
			HandlerFunc:  GetMedicineHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/medicine/list
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/medicine/list",
			HandlerFunc:  ListMedicineHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/medicine/query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/medicine/query_by_text",
			HandlerFunc:  QueryByTextMedicineHandler,
			Authenticate: true,
		},
		{
			// API Route for POST pharmacy/pharmaceutical_company
			Method:       "POST",
			Version:      1,
			Path:         "pharmacy/pharmaceutical_company",
			HandlerFunc:  AddPharmaceuticalCompanyHandler,
			Authenticate: true,
		},
		{
			// API Route for PUT pharmacy/pharmaceutical_company
			Method:       "PUT",
			Version:      1,
			Path:         "pharmacy/pharmaceutical_company",
			HandlerFunc:  UpdatePharmaceuticalCompanyHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/pharmaceutical_company
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/pharmaceutical_company",
			HandlerFunc:  GetPharmaceuticalCompanyHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/pharmaceutical_company/list
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/pharmaceutical_company/list",
			HandlerFunc:  ListPharmaceuticalCompanyHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/pharmaceutical_company/query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/pharmaceutical_company/query_by_text",
			HandlerFunc:  QueryByTextPharmaceuticalCompanyHandler,
			Authenticate: true,
		},
		{
			// API Route for POST pharmacy/product
			Method:       "POST",
			Version:      1,
			Path:         "pharmacy/product",
			HandlerFunc:  AddProductHandler,
			Authenticate: true,
		},
		{
			// API Route for PUT pharmacy/product
			Method:       "PUT",
			Version:      1,
			Path:         "pharmacy/product",
			HandlerFunc:  UpdateProductHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/product
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/product",
			HandlerFunc:  GetProductHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/product/list
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/product/list",
			HandlerFunc:  ListProductHandler,
			Authenticate: true,
		},
		{
			// API Route for GET pharmacy/product/query_by_text
			Method:       "GET",
			Version:      1,
			Path:         "pharmacy/product/query_by_text",
			HandlerFunc:  QueryByTextProductHandler,
			Authenticate: true,
		},
	}

	return routes
}

// AddDrugHandler is the HTTP handler for the method AddDrug.
// The method's description: Adds a new Drug entity
func AddDrugHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AddDrugHandler starting...")
	// Get the req from HTTP body
	var req drug_types.Drug
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.AddDrug(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// UpdateDrugHandler is the HTTP handler for the method UpdateDrug.
// The method's description: Adds a new Drug entity
func UpdateDrugHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] UpdateDrugHandler starting...")
	// Get the req from HTTP body
	var req drug_types.UpdateDrugRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.UpdateDrug(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// GetDrugHandler is the HTTP handler for the method GetDrug.
// The method's description: Get a Drug entity
func GetDrugHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] GetDrugHandler starting...")
	// Get the req from the HTTP req body
	var req drug_types.GetDrugRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.GetDrug(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// ListDrugHandler is the HTTP handler for the method ListDrug.
// The method's description: List Drug entities
func ListDrugHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] ListDrugHandler starting...")
	// Get the req from the HTTP req body
	var req drug_types.ListDrugRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.ListDrug(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// QueryByTextDrugHandler is the HTTP handler for the method QueryByTextDrug.
// The method's description: List Drugs entities by free text search
func QueryByTextDrugHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] QueryByTextDrugHandler starting...")
	// Get the req from the HTTP req body
	var req drug_types.QueryByTextDrugRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.QueryByTextDrug(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// AddMedicineHandler is the HTTP handler for the method AddMedicine.
// The method's description: Adds a new Medicine entity
func AddMedicineHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AddMedicineHandler starting...")
	// Get the req from HTTP body
	var req medicine_types.Medicine
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.AddMedicine(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// UpdateMedicineHandler is the HTTP handler for the method UpdateMedicine.
// The method's description: Adds a new Medicine entity
func UpdateMedicineHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] UpdateMedicineHandler starting...")
	// Get the req from HTTP body
	var req medicine_types.UpdateMedicineRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.UpdateMedicine(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// GetMedicineHandler is the HTTP handler for the method GetMedicine.
// The method's description: Get a Medicine entity
func GetMedicineHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] GetMedicineHandler starting...")
	// Get the req from the HTTP req body
	var req medicine_types.GetMedicineRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.GetMedicine(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// ListMedicineHandler is the HTTP handler for the method ListMedicine.
// The method's description: List Medicine entities
func ListMedicineHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] ListMedicineHandler starting...")
	// Get the req from the HTTP req body
	var req medicine_types.ListMedicineRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.ListMedicine(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// QueryByTextMedicineHandler is the HTTP handler for the method QueryByTextMedicine.
// The method's description: List Medicines entities by free text search
func QueryByTextMedicineHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] QueryByTextMedicineHandler starting...")
	// Get the req from the HTTP req body
	var req medicine_types.QueryByTextMedicineRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.QueryByTextMedicine(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// AddPharmaceuticalCompanyHandler is the HTTP handler for the method AddPharmaceuticalCompany.
// The method's description: Adds a new PharmaceuticalCompany entity
func AddPharmaceuticalCompanyHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AddPharmaceuticalCompanyHandler starting...")
	// Get the req from HTTP body
	var req pharmaceutical_company_types.PharmaceuticalCompany
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.AddPharmaceuticalCompany(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// UpdatePharmaceuticalCompanyHandler is the HTTP handler for the method UpdatePharmaceuticalCompany.
// The method's description: Adds a new PharmaceuticalCompany entity
func UpdatePharmaceuticalCompanyHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] UpdatePharmaceuticalCompanyHandler starting...")
	// Get the req from HTTP body
	var req pharmaceutical_company_types.UpdatePharmaceuticalCompanyRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.UpdatePharmaceuticalCompany(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// GetPharmaceuticalCompanyHandler is the HTTP handler for the method GetPharmaceuticalCompany.
// The method's description: Get a PharmaceuticalCompany entity
func GetPharmaceuticalCompanyHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] GetPharmaceuticalCompanyHandler starting...")
	// Get the req from the HTTP req body
	var req pharmaceutical_company_types.GetPharmaceuticalCompanyRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.GetPharmaceuticalCompany(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// ListPharmaceuticalCompanyHandler is the HTTP handler for the method ListPharmaceuticalCompany.
// The method's description: List PharmaceuticalCompany entities
func ListPharmaceuticalCompanyHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] ListPharmaceuticalCompanyHandler starting...")
	// Get the req from the HTTP req body
	var req pharmaceutical_company_types.ListPharmaceuticalCompanyRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.ListPharmaceuticalCompany(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// QueryByTextPharmaceuticalCompanyHandler is the HTTP handler for the method QueryByTextPharmaceuticalCompany.
// The method's description: List PharmaceuticalCompanies entities by free text search
func QueryByTextPharmaceuticalCompanyHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] QueryByTextPharmaceuticalCompanyHandler starting...")
	// Get the req from the HTTP req body
	var req pharmaceutical_company_types.QueryByTextPharmaceuticalCompanyRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.QueryByTextPharmaceuticalCompany(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// AddProductHandler is the HTTP handler for the method AddProduct.
// The method's description: Adds a new Product entity
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] AddProductHandler starting...")
	// Get the req from HTTP body
	var req product_types.Product
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.AddProduct(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// UpdateProductHandler is the HTTP handler for the method UpdateProduct.
// The method's description: Adds a new Product entity
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] UpdateProductHandler starting...")
	// Get the req from HTTP body
	var req product_types.UpdateProductRequest
	err := httputil.UnmarshalJSONFromRequest(r, &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.UpdateProduct(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 201, resp)

}

// GetProductHandler is the HTTP handler for the method GetProduct.
// The method's description: Get a Product entity
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] GetProductHandler starting...")
	// Get the req from the HTTP req body
	var req product_types.GetProductRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.GetProduct(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// ListProductHandler is the HTTP handler for the method ListProduct.
// The method's description: List Product entities
func ListProductHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] ListProductHandler starting...")
	// Get the req from the HTTP req body
	var req product_types.ListProductRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.ListProduct(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}

// QueryByTextProductHandler is the HTTP handler for the method QueryByTextProduct.
// The method's description: List Products entities by free text search
func QueryByTextProductHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("[HTTP Handler] QueryByTextProductHandler starting...")
	// Get the req from the HTTP req body
	var req product_types.QueryByTextProductRequest
	reqParam, ok := r.URL.Query()["req"]
	if !ok || len(reqParam) < 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("URL param 'req' is required"), false, nil)
		return
	}
	if len(reqParam) > 1 {
		gopi.WriteError(w, http.StatusBadRequest, fmt.Errorf("multiple URL params with name 'req' found"), false, nil)
		return
	}
	err := json.Unmarshal([]byte(reqParam[0]), &req)
	if err != nil {
		gopi.WriteError(w, http.StatusBadRequest, err, false, nil)
		return
	}

	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.QueryByTextProduct(r.Context(), req)
	if err != nil {
		errutil.HandleHTTPResponseError(w, err)
		return
	}

	gopi.WriteResponse(w, 200, resp)

}
