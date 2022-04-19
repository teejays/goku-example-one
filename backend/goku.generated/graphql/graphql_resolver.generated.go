package example_app_graphql_resolver

import (
	"context"
	"net/http"
	"time"

	"github.com/teejays/clog"
	"github.com/teejays/goku-util/ctxutil"

	example_app_types "github.com/teejays/goku-example-one/backend/goku.generated/types"
	drug_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/methods"
	drug_types "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/types"
	medicine_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/methods"
	medicine_types "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/types"
	pharmaceutical_company_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/methods"
	pharmaceutical_company_types "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
	product_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/methods"
	product_types "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/types"
	users_methods "github.com/teejays/goku-example-one/backend/services/users/goku.generated/methods"
	users_types "github.com/teejays/goku-example-one/backend/services/users/goku.generated/types"
	user_methods "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/methods"
	user_types "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/types"
)

type AddressResolver struct {
	example_app_types.Address
}

type AddressWithMetaResolver struct {
	example_app_types.AddressWithMeta
}

type AuthenticateRequestResolver struct {
	users_types.AuthenticateRequest
}

type AuthenticateResponseResolver struct {
	users_types.AuthenticateResponse
}

type ContactResolver struct {
	example_app_types.Contact
}

func (r *ContactResolver) Name() *PersonNameResolver {
	return &PersonNameResolver{PersonName: r.Contact.Name}
}
func (r *ContactResolver) Address() *AddressResolver {
	return &AddressResolver{Address: r.Contact.Address}
}

type ContactWithMetaResolver struct {
	example_app_types.ContactWithMeta
}

func (r *ContactWithMetaResolver) Name() *PersonNameResolver {
	return &PersonNameResolver{PersonName: r.ContactWithMeta.Name}
}
func (r *ContactWithMetaResolver) Address() *AddressResolver {
	return &AddressResolver{Address: r.ContactWithMeta.Address}
}

type DrugResolver struct {
	drug_types.Drug
}

type IngredientResolver struct {
	medicine_types.Ingredient
}

type IngredientWithMetaResolver struct {
	medicine_types.IngredientWithMeta
}

type MedicineResolver struct {
	medicine_types.Medicine
}

func (r *MedicineResolver) PrimaryIngredient() *IngredientResolver {
	return &IngredientResolver{Ingredient: r.Medicine.PrimaryIngredient}
}
func (r *MedicineResolver) Ingredients() []*IngredientWithMetaResolver {
	var ret []*IngredientWithMetaResolver
	for _, elem := range r.Medicine.Ingredients {
		ret = append(ret, &IngredientWithMetaResolver{IngredientWithMeta: elem})
	}
	return ret
}

type PersonNameResolver struct {
	example_app_types.PersonName
}

type PersonNameWithMetaResolver struct {
	example_app_types.PersonNameWithMeta
}

type PharmaceuticalCompanyResolver struct {
	pharmaceutical_company_types.PharmaceuticalCompany
}

type PhoneNumberResolver struct {
	example_app_types.PhoneNumber
}

func (r *PhoneNumberResolver) CountryCode() int32 {
	return int32(r.PhoneNumber.CountryCode)
}

type PhoneNumberWithMetaResolver struct {
	example_app_types.PhoneNumberWithMeta
}

func (r *PhoneNumberWithMetaResolver) CountryCode() int32 {
	return int32(r.PhoneNumberWithMeta.CountryCode)
}

type ProductResolver struct {
	product_types.Product
}

func (r *ProductResolver) Mass() int32 {
	return int32(r.Product.Mass)
}
func (r *ProductResolver) Count() int32 {
	return int32(r.Product.Count)
}

type RegisterUserRequestResolver struct {
	users_types.RegisterUserRequest
}

func (r *RegisterUserRequestResolver) Name() *PersonNameResolver {
	return &PersonNameResolver{PersonName: r.RegisterUserRequest.Name}
}
func (r *RegisterUserRequestResolver) PhoneNumber() *PhoneNumberResolver {
	return &PhoneNumberResolver{PhoneNumber: r.RegisterUserRequest.PhoneNumber}
}

type UserResolver struct {
	user_types.User
}

func (r *UserResolver) Name() *PersonNameResolver {
	return &PersonNameResolver{PersonName: r.User.Name}
}
func (r *UserResolver) PhoneNumber() *PhoneNumberResolver {
	return &PhoneNumberResolver{PhoneNumber: *r.User.PhoneNumber}
}

type AddressFilterResolver struct {
	example_app_types.AddressFilter
}

func (r *AddressFilterResolver) Province() *PakistaniProvinceConditionResolver {
	return &PakistaniProvinceConditionResolver{PakistaniProvinceCondition: *r.AddressFilter.Province}
}
func (r *AddressFilterResolver) Country() *CountryConditionResolver {
	return &CountryConditionResolver{CountryCondition: *r.AddressFilter.Country}
}
func (r *AddressFilterResolver) And() []*AddressFilterResolver {
	var ret []*AddressFilterResolver
	for _, elem := range r.AddressFilter.And {
		ret = append(ret, &AddressFilterResolver{AddressFilter: elem})
	}
	return ret
}
func (r *AddressFilterResolver) Or() []*AddressFilterResolver {
	var ret []*AddressFilterResolver
	for _, elem := range r.AddressFilter.Or {
		ret = append(ret, &AddressFilterResolver{AddressFilter: elem})
	}
	return ret
}

type ContactFilterResolver struct {
	example_app_types.ContactFilter
}

func (r *ContactFilterResolver) Name() *PersonNameFilterResolver {
	return &PersonNameFilterResolver{PersonNameFilter: *r.ContactFilter.Name}
}
func (r *ContactFilterResolver) Address() *AddressFilterResolver {
	return &AddressFilterResolver{AddressFilter: *r.ContactFilter.Address}
}
func (r *ContactFilterResolver) And() []*ContactFilterResolver {
	var ret []*ContactFilterResolver
	for _, elem := range r.ContactFilter.And {
		ret = append(ret, &ContactFilterResolver{ContactFilter: elem})
	}
	return ret
}
func (r *ContactFilterResolver) Or() []*ContactFilterResolver {
	var ret []*ContactFilterResolver
	for _, elem := range r.ContactFilter.Or {
		ret = append(ret, &ContactFilterResolver{ContactFilter: elem})
	}
	return ret
}

type DrugFilterResolver struct {
	drug_types.DrugFilter
}

func (r *DrugFilterResolver) And() []*DrugFilterResolver {
	var ret []*DrugFilterResolver
	for _, elem := range r.DrugFilter.And {
		ret = append(ret, &DrugFilterResolver{DrugFilter: elem})
	}
	return ret
}
func (r *DrugFilterResolver) Or() []*DrugFilterResolver {
	var ret []*DrugFilterResolver
	for _, elem := range r.DrugFilter.Or {
		ret = append(ret, &DrugFilterResolver{DrugFilter: elem})
	}
	return ret
}

type IngredientFilterResolver struct {
	medicine_types.IngredientFilter
}

func (r *IngredientFilterResolver) And() []*IngredientFilterResolver {
	var ret []*IngredientFilterResolver
	for _, elem := range r.IngredientFilter.And {
		ret = append(ret, &IngredientFilterResolver{IngredientFilter: elem})
	}
	return ret
}
func (r *IngredientFilterResolver) Or() []*IngredientFilterResolver {
	var ret []*IngredientFilterResolver
	for _, elem := range r.IngredientFilter.Or {
		ret = append(ret, &IngredientFilterResolver{IngredientFilter: elem})
	}
	return ret
}

type MedicineFilterResolver struct {
	medicine_types.MedicineFilter
}

func (r *MedicineFilterResolver) PrimaryIngredient() *IngredientFilterResolver {
	return &IngredientFilterResolver{IngredientFilter: *r.MedicineFilter.PrimaryIngredient}
}
func (r *MedicineFilterResolver) HavingIngredients() *IngredientFilterResolver {
	return &IngredientFilterResolver{IngredientFilter: *r.MedicineFilter.HavingIngredients}
}
func (r *MedicineFilterResolver) ModeOfDelivery() *ModeOfDeliveryConditionResolver {
	return &ModeOfDeliveryConditionResolver{ModeOfDeliveryCondition: *r.MedicineFilter.ModeOfDelivery}
}
func (r *MedicineFilterResolver) And() []*MedicineFilterResolver {
	var ret []*MedicineFilterResolver
	for _, elem := range r.MedicineFilter.And {
		ret = append(ret, &MedicineFilterResolver{MedicineFilter: elem})
	}
	return ret
}
func (r *MedicineFilterResolver) Or() []*MedicineFilterResolver {
	var ret []*MedicineFilterResolver
	for _, elem := range r.MedicineFilter.Or {
		ret = append(ret, &MedicineFilterResolver{MedicineFilter: elem})
	}
	return ret
}

type PersonNameFilterResolver struct {
	example_app_types.PersonNameFilter
}

func (r *PersonNameFilterResolver) And() []*PersonNameFilterResolver {
	var ret []*PersonNameFilterResolver
	for _, elem := range r.PersonNameFilter.And {
		ret = append(ret, &PersonNameFilterResolver{PersonNameFilter: elem})
	}
	return ret
}
func (r *PersonNameFilterResolver) Or() []*PersonNameFilterResolver {
	var ret []*PersonNameFilterResolver
	for _, elem := range r.PersonNameFilter.Or {
		ret = append(ret, &PersonNameFilterResolver{PersonNameFilter: elem})
	}
	return ret
}

type PharmaceuticalCompanyFilterResolver struct {
	pharmaceutical_company_types.PharmaceuticalCompanyFilter
}

func (r *PharmaceuticalCompanyFilterResolver) And() []*PharmaceuticalCompanyFilterResolver {
	var ret []*PharmaceuticalCompanyFilterResolver
	for _, elem := range r.PharmaceuticalCompanyFilter.And {
		ret = append(ret, &PharmaceuticalCompanyFilterResolver{PharmaceuticalCompanyFilter: elem})
	}
	return ret
}
func (r *PharmaceuticalCompanyFilterResolver) Or() []*PharmaceuticalCompanyFilterResolver {
	var ret []*PharmaceuticalCompanyFilterResolver
	for _, elem := range r.PharmaceuticalCompanyFilter.Or {
		ret = append(ret, &PharmaceuticalCompanyFilterResolver{PharmaceuticalCompanyFilter: elem})
	}
	return ret
}

type PhoneNumberFilterResolver struct {
	example_app_types.PhoneNumberFilter
}

func (r *PhoneNumberFilterResolver) And() []*PhoneNumberFilterResolver {
	var ret []*PhoneNumberFilterResolver
	for _, elem := range r.PhoneNumberFilter.And {
		ret = append(ret, &PhoneNumberFilterResolver{PhoneNumberFilter: elem})
	}
	return ret
}
func (r *PhoneNumberFilterResolver) Or() []*PhoneNumberFilterResolver {
	var ret []*PhoneNumberFilterResolver
	for _, elem := range r.PhoneNumberFilter.Or {
		ret = append(ret, &PhoneNumberFilterResolver{PhoneNumberFilter: elem})
	}
	return ret
}

type ProductFilterResolver struct {
	product_types.ProductFilter
}

func (r *ProductFilterResolver) And() []*ProductFilterResolver {
	var ret []*ProductFilterResolver
	for _, elem := range r.ProductFilter.And {
		ret = append(ret, &ProductFilterResolver{ProductFilter: elem})
	}
	return ret
}
func (r *ProductFilterResolver) Or() []*ProductFilterResolver {
	var ret []*ProductFilterResolver
	for _, elem := range r.ProductFilter.Or {
		ret = append(ret, &ProductFilterResolver{ProductFilter: elem})
	}
	return ret
}

type UserFilterResolver struct {
	user_types.UserFilter
}

func (r *UserFilterResolver) Name() *PersonNameFilterResolver {
	return &PersonNameFilterResolver{PersonNameFilter: *r.UserFilter.Name}
}
func (r *UserFilterResolver) PhoneNumber() *PhoneNumberFilterResolver {
	return &PhoneNumberFilterResolver{PhoneNumberFilter: *r.UserFilter.PhoneNumber}
}
func (r *UserFilterResolver) And() []*UserFilterResolver {
	var ret []*UserFilterResolver
	for _, elem := range r.UserFilter.And {
		ret = append(ret, &UserFilterResolver{UserFilter: elem})
	}
	return ret
}
func (r *UserFilterResolver) Or() []*UserFilterResolver {
	var ret []*UserFilterResolver
	for _, elem := range r.UserFilter.Or {
		ret = append(ret, &UserFilterResolver{UserFilter: elem})
	}
	return ret
}

type CountryConditionResolver struct {
	example_app_types.CountryCondition
}

type ModeOfDeliveryConditionResolver struct {
	medicine_types.ModeOfDeliveryCondition
}

type PakistaniProvinceConditionResolver struct {
	example_app_types.PakistaniProvinceCondition
}

type AddressFieldConditionResolver struct {
	example_app_types.AddressFieldCondition
}

type ContactFieldConditionResolver struct {
	example_app_types.ContactFieldCondition
}

type DrugFieldConditionResolver struct {
	drug_types.DrugFieldCondition
}

type IngredientFieldConditionResolver struct {
	medicine_types.IngredientFieldCondition
}

type MedicineFieldConditionResolver struct {
	medicine_types.MedicineFieldCondition
}

type PersonNameFieldConditionResolver struct {
	example_app_types.PersonNameFieldCondition
}

type PharmaceuticalCompanyFieldConditionResolver struct {
	pharmaceutical_company_types.PharmaceuticalCompanyFieldCondition
}

type PhoneNumberFieldConditionResolver struct {
	example_app_types.PhoneNumberFieldCondition
}

type ProductFieldConditionResolver struct {
	product_types.ProductFieldCondition
}

type UserFieldConditionResolver struct {
	user_types.UserFieldCondition
}

type UpdateDrugRequestResolver struct {
	drug_types.UpdateDrugRequest
}

func (r *UpdateDrugRequestResolver) Object() *DrugResolver {
	return &DrugResolver{Drug: r.UpdateDrugRequest.Object}
}

type UpdateDrugResponseResolver struct {
	drug_types.UpdateDrugResponse
}

func (r *UpdateDrugResponseResolver) Object() *DrugResolver {
	return &DrugResolver{Drug: r.UpdateDrugResponse.Object}
}

type GetDrugRequestResolver struct {
	drug_types.GetDrugRequest
}

type ListDrugRequestResolver struct {
	drug_types.ListDrugRequest
}

func (r *ListDrugRequestResolver) Filter() *DrugFilterResolver {
	return &DrugFilterResolver{DrugFilter: r.ListDrugRequest.Filter}
}

type ListDrugResponseResolver struct {
	drug_types.ListDrugResponse
}

func (r *ListDrugResponseResolver) Items() []*DrugResolver {
	var ret []*DrugResolver
	for _, elem := range r.ListDrugResponse.Items {
		ret = append(ret, &DrugResolver{Drug: elem})
	}
	return ret
}
func (r *ListDrugResponseResolver) Count() int32 {
	return int32(r.ListDrugResponse.Count)
}

type QueryByTextDrugRequestResolver struct {
	drug_types.QueryByTextDrugRequest
}

type UpdateMedicineRequestResolver struct {
	medicine_types.UpdateMedicineRequest
}

func (r *UpdateMedicineRequestResolver) Object() *MedicineResolver {
	return &MedicineResolver{Medicine: r.UpdateMedicineRequest.Object}
}

type UpdateMedicineResponseResolver struct {
	medicine_types.UpdateMedicineResponse
}

func (r *UpdateMedicineResponseResolver) Object() *MedicineResolver {
	return &MedicineResolver{Medicine: r.UpdateMedicineResponse.Object}
}

type GetMedicineRequestResolver struct {
	medicine_types.GetMedicineRequest
}

type ListMedicineRequestResolver struct {
	medicine_types.ListMedicineRequest
}

func (r *ListMedicineRequestResolver) Filter() *MedicineFilterResolver {
	return &MedicineFilterResolver{MedicineFilter: r.ListMedicineRequest.Filter}
}

type ListMedicineResponseResolver struct {
	medicine_types.ListMedicineResponse
}

func (r *ListMedicineResponseResolver) Items() []*MedicineResolver {
	var ret []*MedicineResolver
	for _, elem := range r.ListMedicineResponse.Items {
		ret = append(ret, &MedicineResolver{Medicine: elem})
	}
	return ret
}
func (r *ListMedicineResponseResolver) Count() int32 {
	return int32(r.ListMedicineResponse.Count)
}

type QueryByTextMedicineRequestResolver struct {
	medicine_types.QueryByTextMedicineRequest
}

type UpdatePharmaceuticalCompanyRequestResolver struct {
	pharmaceutical_company_types.UpdatePharmaceuticalCompanyRequest
}

func (r *UpdatePharmaceuticalCompanyRequestResolver) Object() *PharmaceuticalCompanyResolver {
	return &PharmaceuticalCompanyResolver{PharmaceuticalCompany: r.UpdatePharmaceuticalCompanyRequest.Object}
}

type UpdatePharmaceuticalCompanyResponseResolver struct {
	pharmaceutical_company_types.UpdatePharmaceuticalCompanyResponse
}

func (r *UpdatePharmaceuticalCompanyResponseResolver) Object() *PharmaceuticalCompanyResolver {
	return &PharmaceuticalCompanyResolver{PharmaceuticalCompany: r.UpdatePharmaceuticalCompanyResponse.Object}
}

type GetPharmaceuticalCompanyRequestResolver struct {
	pharmaceutical_company_types.GetPharmaceuticalCompanyRequest
}

type ListPharmaceuticalCompanyRequestResolver struct {
	pharmaceutical_company_types.ListPharmaceuticalCompanyRequest
}

func (r *ListPharmaceuticalCompanyRequestResolver) Filter() *PharmaceuticalCompanyFilterResolver {
	return &PharmaceuticalCompanyFilterResolver{PharmaceuticalCompanyFilter: r.ListPharmaceuticalCompanyRequest.Filter}
}

type ListPharmaceuticalCompanyResponseResolver struct {
	pharmaceutical_company_types.ListPharmaceuticalCompanyResponse
}

func (r *ListPharmaceuticalCompanyResponseResolver) Items() []*PharmaceuticalCompanyResolver {
	var ret []*PharmaceuticalCompanyResolver
	for _, elem := range r.ListPharmaceuticalCompanyResponse.Items {
		ret = append(ret, &PharmaceuticalCompanyResolver{PharmaceuticalCompany: elem})
	}
	return ret
}
func (r *ListPharmaceuticalCompanyResponseResolver) Count() int32 {
	return int32(r.ListPharmaceuticalCompanyResponse.Count)
}

type QueryByTextPharmaceuticalCompanyRequestResolver struct {
	pharmaceutical_company_types.QueryByTextPharmaceuticalCompanyRequest
}

type UpdateProductRequestResolver struct {
	product_types.UpdateProductRequest
}

func (r *UpdateProductRequestResolver) Object() *ProductResolver {
	return &ProductResolver{Product: r.UpdateProductRequest.Object}
}

type UpdateProductResponseResolver struct {
	product_types.UpdateProductResponse
}

func (r *UpdateProductResponseResolver) Object() *ProductResolver {
	return &ProductResolver{Product: r.UpdateProductResponse.Object}
}

type GetProductRequestResolver struct {
	product_types.GetProductRequest
}

type ListProductRequestResolver struct {
	product_types.ListProductRequest
}

func (r *ListProductRequestResolver) Filter() *ProductFilterResolver {
	return &ProductFilterResolver{ProductFilter: r.ListProductRequest.Filter}
}

type ListProductResponseResolver struct {
	product_types.ListProductResponse
}

func (r *ListProductResponseResolver) Items() []*ProductResolver {
	var ret []*ProductResolver
	for _, elem := range r.ListProductResponse.Items {
		ret = append(ret, &ProductResolver{Product: elem})
	}
	return ret
}
func (r *ListProductResponseResolver) Count() int32 {
	return int32(r.ListProductResponse.Count)
}

type QueryByTextProductRequestResolver struct {
	product_types.QueryByTextProductRequest
}

type UpdateUserRequestResolver struct {
	user_types.UpdateUserRequest
}

func (r *UpdateUserRequestResolver) Object() *UserResolver {
	return &UserResolver{User: r.UpdateUserRequest.Object}
}

type UpdateUserResponseResolver struct {
	user_types.UpdateUserResponse
}

func (r *UpdateUserResponseResolver) Object() *UserResolver {
	return &UserResolver{User: r.UpdateUserResponse.Object}
}

type GetUserRequestResolver struct {
	user_types.GetUserRequest
}

type ListUserRequestResolver struct {
	user_types.ListUserRequest
}

func (r *ListUserRequestResolver) Filter() *UserFilterResolver {
	return &UserFilterResolver{UserFilter: r.ListUserRequest.Filter}
}

type ListUserResponseResolver struct {
	user_types.ListUserResponse
}

func (r *ListUserResponseResolver) Items() []*UserResolver {
	var ret []*UserResolver
	for _, elem := range r.ListUserResponse.Items {
		ret = append(ret, &UserResolver{User: elem})
	}
	return ret
}
func (r *ListUserResponseResolver) Count() int32 {
	return int32(r.ListUserResponse.Count)
}

type QueryByTextUserRequestResolver struct {
	user_types.QueryByTextUserRequest
}

// Helper Funcs

var gClient *http.Client

func getHTTPClient() *http.Client {
	if gClient != nil {
		return gClient
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	gClient = &http.Client{Transport: tr}
	return gClient
}

func getBaseURL(ctx context.Context) string {
	if ctxutil.GetEnv(ctx) == ctxutil.DEV {
		return "http://127.0.0.1:8080"
	}
	// TODO: If we have a prod env, fix this
	return "http://localhost:8080"
}

// Resolver structs and their methods

type Resolver struct {
	QueryResolver
	MutationResolver
}

type QueryResolver struct{}

type AuthenticateRequestArgs struct {
	Req users_types.AuthenticateRequest
}

func (r *QueryResolver) AuthenticateUser(ctx context.Context, args AuthenticateRequestArgs) (*AuthenticateResponseResolver, error) {
	var ret AuthenticateResponseResolver
	var err error

	clog.Infof("[Graphql Handler] AuthenticateUser starting...")
	// Get the method's server for this service
	s := users_methods.NewServer()

	resp, err := s.AuthenticateUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.AuthenticateResponse = resp

	return &ret, nil
}

type GetDrugRequestArgs struct {
	Req drug_types.GetDrugRequest
}

func (r *QueryResolver) GetDrug(ctx context.Context, args GetDrugRequestArgs) (*DrugResolver, error) {
	var ret DrugResolver
	var err error

	clog.Infof("[Graphql Handler] GetDrug starting...")
	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.GetDrug(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Drug = resp

	return &ret, nil
}

type ListDrugRequestArgs struct {
	Req drug_types.ListDrugRequest
}

func (r *QueryResolver) ListDrug(ctx context.Context, args ListDrugRequestArgs) (*ListDrugResponseResolver, error) {
	var ret ListDrugResponseResolver
	var err error

	clog.Infof("[Graphql Handler] ListDrug starting...")
	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.ListDrug(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListDrugResponse = resp

	return &ret, nil
}

type QueryByTextDrugRequestArgs struct {
	Req drug_types.QueryByTextDrugRequest
}

func (r *QueryResolver) QueryByTextDrug(ctx context.Context, args QueryByTextDrugRequestArgs) (*ListDrugResponseResolver, error) {
	var ret ListDrugResponseResolver
	var err error

	clog.Infof("[Graphql Handler] QueryByTextDrug starting...")
	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.QueryByTextDrug(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListDrugResponse = resp

	return &ret, nil
}

type GetMedicineRequestArgs struct {
	Req medicine_types.GetMedicineRequest
}

func (r *QueryResolver) GetMedicine(ctx context.Context, args GetMedicineRequestArgs) (*MedicineResolver, error) {
	var ret MedicineResolver
	var err error

	clog.Infof("[Graphql Handler] GetMedicine starting...")
	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.GetMedicine(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Medicine = resp

	return &ret, nil
}

type ListMedicineRequestArgs struct {
	Req medicine_types.ListMedicineRequest
}

func (r *QueryResolver) ListMedicine(ctx context.Context, args ListMedicineRequestArgs) (*ListMedicineResponseResolver, error) {
	var ret ListMedicineResponseResolver
	var err error

	clog.Infof("[Graphql Handler] ListMedicine starting...")
	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.ListMedicine(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListMedicineResponse = resp

	return &ret, nil
}

type QueryByTextMedicineRequestArgs struct {
	Req medicine_types.QueryByTextMedicineRequest
}

func (r *QueryResolver) QueryByTextMedicine(ctx context.Context, args QueryByTextMedicineRequestArgs) (*ListMedicineResponseResolver, error) {
	var ret ListMedicineResponseResolver
	var err error

	clog.Infof("[Graphql Handler] QueryByTextMedicine starting...")
	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.QueryByTextMedicine(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListMedicineResponse = resp

	return &ret, nil
}

type GetPharmaceuticalCompanyRequestArgs struct {
	Req pharmaceutical_company_types.GetPharmaceuticalCompanyRequest
}

func (r *QueryResolver) GetPharmaceuticalCompany(ctx context.Context, args GetPharmaceuticalCompanyRequestArgs) (*PharmaceuticalCompanyResolver, error) {
	var ret PharmaceuticalCompanyResolver
	var err error

	clog.Infof("[Graphql Handler] GetPharmaceuticalCompany starting...")
	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.GetPharmaceuticalCompany(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.PharmaceuticalCompany = resp

	return &ret, nil
}

type ListPharmaceuticalCompanyRequestArgs struct {
	Req pharmaceutical_company_types.ListPharmaceuticalCompanyRequest
}

func (r *QueryResolver) ListPharmaceuticalCompany(ctx context.Context, args ListPharmaceuticalCompanyRequestArgs) (*ListPharmaceuticalCompanyResponseResolver, error) {
	var ret ListPharmaceuticalCompanyResponseResolver
	var err error

	clog.Infof("[Graphql Handler] ListPharmaceuticalCompany starting...")
	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.ListPharmaceuticalCompany(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListPharmaceuticalCompanyResponse = resp

	return &ret, nil
}

type QueryByTextPharmaceuticalCompanyRequestArgs struct {
	Req pharmaceutical_company_types.QueryByTextPharmaceuticalCompanyRequest
}

func (r *QueryResolver) QueryByTextPharmaceuticalCompany(ctx context.Context, args QueryByTextPharmaceuticalCompanyRequestArgs) (*ListPharmaceuticalCompanyResponseResolver, error) {
	var ret ListPharmaceuticalCompanyResponseResolver
	var err error

	clog.Infof("[Graphql Handler] QueryByTextPharmaceuticalCompany starting...")
	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.QueryByTextPharmaceuticalCompany(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListPharmaceuticalCompanyResponse = resp

	return &ret, nil
}

type GetProductRequestArgs struct {
	Req product_types.GetProductRequest
}

func (r *QueryResolver) GetProduct(ctx context.Context, args GetProductRequestArgs) (*ProductResolver, error) {
	var ret ProductResolver
	var err error

	clog.Infof("[Graphql Handler] GetProduct starting...")
	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.GetProduct(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Product = resp

	return &ret, nil
}

type ListProductRequestArgs struct {
	Req product_types.ListProductRequest
}

func (r *QueryResolver) ListProduct(ctx context.Context, args ListProductRequestArgs) (*ListProductResponseResolver, error) {
	var ret ListProductResponseResolver
	var err error

	clog.Infof("[Graphql Handler] ListProduct starting...")
	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.ListProduct(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListProductResponse = resp

	return &ret, nil
}

type QueryByTextProductRequestArgs struct {
	Req product_types.QueryByTextProductRequest
}

func (r *QueryResolver) QueryByTextProduct(ctx context.Context, args QueryByTextProductRequestArgs) (*ListProductResponseResolver, error) {
	var ret ListProductResponseResolver
	var err error

	clog.Infof("[Graphql Handler] QueryByTextProduct starting...")
	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.QueryByTextProduct(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListProductResponse = resp

	return &ret, nil
}

type GetUserRequestArgs struct {
	Req user_types.GetUserRequest
}

func (r *QueryResolver) GetUser(ctx context.Context, args GetUserRequestArgs) (*UserResolver, error) {
	var ret UserResolver
	var err error

	clog.Infof("[Graphql Handler] GetUser starting...")
	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.GetUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.User = resp

	return &ret, nil
}

type ListUserRequestArgs struct {
	Req user_types.ListUserRequest
}

func (r *QueryResolver) ListUser(ctx context.Context, args ListUserRequestArgs) (*ListUserResponseResolver, error) {
	var ret ListUserResponseResolver
	var err error

	clog.Infof("[Graphql Handler] ListUser starting...")
	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.ListUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListUserResponse = resp

	return &ret, nil
}

type QueryByTextUserRequestArgs struct {
	Req user_types.QueryByTextUserRequest
}

func (r *QueryResolver) QueryByTextUser(ctx context.Context, args QueryByTextUserRequestArgs) (*ListUserResponseResolver, error) {
	var ret ListUserResponseResolver
	var err error

	clog.Infof("[Graphql Handler] QueryByTextUser starting...")
	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.QueryByTextUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.ListUserResponse = resp

	return &ret, nil
}

type MutationResolver struct{}

type RegisterUserRequestArgs struct {
	Req users_types.RegisterUserRequest
}

func (r *MutationResolver) RegisterUser(ctx context.Context, args RegisterUserRequestArgs) (*AuthenticateResponseResolver, error) {
	var ret AuthenticateResponseResolver
	var err error

	clog.Infof("[Graphql Handler] RegisterUser starting...")
	// Get the method's server for this service
	s := users_methods.NewServer()

	resp, err := s.RegisterUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.AuthenticateResponse = resp

	return &ret, nil
}

type DrugArgs struct {
	Req drug_types.Drug
}

func (r *MutationResolver) AddDrug(ctx context.Context, args DrugArgs) (*DrugResolver, error) {
	var ret DrugResolver
	var err error

	clog.Infof("[Graphql Handler] AddDrug starting...")
	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.AddDrug(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Drug = resp

	return &ret, nil
}

type UpdateDrugRequestArgs struct {
	Req drug_types.UpdateDrugRequest
}

func (r *MutationResolver) UpdateDrug(ctx context.Context, args UpdateDrugRequestArgs) (*UpdateDrugResponseResolver, error) {
	var ret UpdateDrugResponseResolver
	var err error

	clog.Infof("[Graphql Handler] UpdateDrug starting...")
	// Get the method's server for this service
	s := drug_methods.NewServer()

	resp, err := s.UpdateDrug(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.UpdateDrugResponse = resp

	return &ret, nil
}

type MedicineArgs struct {
	Req medicine_types.Medicine
}

func (r *MutationResolver) AddMedicine(ctx context.Context, args MedicineArgs) (*MedicineResolver, error) {
	var ret MedicineResolver
	var err error

	clog.Infof("[Graphql Handler] AddMedicine starting...")
	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.AddMedicine(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Medicine = resp

	return &ret, nil
}

type UpdateMedicineRequestArgs struct {
	Req medicine_types.UpdateMedicineRequest
}

func (r *MutationResolver) UpdateMedicine(ctx context.Context, args UpdateMedicineRequestArgs) (*UpdateMedicineResponseResolver, error) {
	var ret UpdateMedicineResponseResolver
	var err error

	clog.Infof("[Graphql Handler] UpdateMedicine starting...")
	// Get the method's server for this service
	s := medicine_methods.NewServer()

	resp, err := s.UpdateMedicine(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.UpdateMedicineResponse = resp

	return &ret, nil
}

type PharmaceuticalCompanyArgs struct {
	Req pharmaceutical_company_types.PharmaceuticalCompany
}

func (r *MutationResolver) AddPharmaceuticalCompany(ctx context.Context, args PharmaceuticalCompanyArgs) (*PharmaceuticalCompanyResolver, error) {
	var ret PharmaceuticalCompanyResolver
	var err error

	clog.Infof("[Graphql Handler] AddPharmaceuticalCompany starting...")
	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.AddPharmaceuticalCompany(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.PharmaceuticalCompany = resp

	return &ret, nil
}

type UpdatePharmaceuticalCompanyRequestArgs struct {
	Req pharmaceutical_company_types.UpdatePharmaceuticalCompanyRequest
}

func (r *MutationResolver) UpdatePharmaceuticalCompany(ctx context.Context, args UpdatePharmaceuticalCompanyRequestArgs) (*UpdatePharmaceuticalCompanyResponseResolver, error) {
	var ret UpdatePharmaceuticalCompanyResponseResolver
	var err error

	clog.Infof("[Graphql Handler] UpdatePharmaceuticalCompany starting...")
	// Get the method's server for this service
	s := pharmaceutical_company_methods.NewServer()

	resp, err := s.UpdatePharmaceuticalCompany(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.UpdatePharmaceuticalCompanyResponse = resp

	return &ret, nil
}

type ProductArgs struct {
	Req product_types.Product
}

func (r *MutationResolver) AddProduct(ctx context.Context, args ProductArgs) (*ProductResolver, error) {
	var ret ProductResolver
	var err error

	clog.Infof("[Graphql Handler] AddProduct starting...")
	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.AddProduct(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.Product = resp

	return &ret, nil
}

type UpdateProductRequestArgs struct {
	Req product_types.UpdateProductRequest
}

func (r *MutationResolver) UpdateProduct(ctx context.Context, args UpdateProductRequestArgs) (*UpdateProductResponseResolver, error) {
	var ret UpdateProductResponseResolver
	var err error

	clog.Infof("[Graphql Handler] UpdateProduct starting...")
	// Get the method's server for this service
	s := product_methods.NewServer()

	resp, err := s.UpdateProduct(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.UpdateProductResponse = resp

	return &ret, nil
}

type UserArgs struct {
	Req user_types.User
}

func (r *MutationResolver) AddUser(ctx context.Context, args UserArgs) (*UserResolver, error) {
	var ret UserResolver
	var err error

	clog.Infof("[Graphql Handler] AddUser starting...")
	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.AddUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.User = resp

	return &ret, nil
}

type UpdateUserRequestArgs struct {
	Req user_types.UpdateUserRequest
}

func (r *MutationResolver) UpdateUser(ctx context.Context, args UpdateUserRequestArgs) (*UpdateUserResponseResolver, error) {
	var ret UpdateUserResponseResolver
	var err error

	clog.Infof("[Graphql Handler] UpdateUser starting...")
	// Get the method's server for this service
	s := user_methods.NewServer()

	resp, err := s.UpdateUser(ctx, args.Req)
	if err != nil {
		return nil, err
	}
	ret.UpdateUserResponse = resp

	return &ret, nil
}
