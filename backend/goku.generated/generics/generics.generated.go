package example_app_generics

import (
	typeslib "github.com/teejays/goku-util/types"

	example_app_types "github.com/teejays/goku-example-one/backend/goku.generated/types"
	drug_types "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/types"
	medicine_types "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/types"
	pharmaceutical_company_types "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
	product_types "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/types"
	user_types "github.com/teejays/goku-example-one/backend/services/users/user/goku.generated/types"
)

type Entity interface {
	drug_types.Drug | medicine_types.Medicine | pharmaceutical_company_types.PharmaceuticalCompany | product_types.Product | user_types.User
}

type BasicType interface {
	example_app_types.Address | example_app_types.Contact | medicine_types.Ingredient | example_app_types.PersonName | example_app_types.PhoneNumber
}

type BasicTypeWithMeta interface {
	example_app_types.AddressWithMeta | example_app_types.ContactWithMeta | medicine_types.IngredientWithMeta | example_app_types.PersonNameWithMeta | example_app_types.PhoneNumberWithMeta
}

type FilterType interface {
	example_app_types.AddressFilter | example_app_types.ContactFilter | drug_types.DrugFilter | medicine_types.IngredientFilter | medicine_types.MedicineFilter | example_app_types.PersonNameFilter | pharmaceutical_company_types.PharmaceuticalCompanyFilter | example_app_types.PhoneNumberFilter | product_types.ProductFilter | user_types.UserFilter
}

type FieldEnum interface {
	example_app_types.AddressField | example_app_types.ContactField | drug_types.DrugField | medicine_types.IngredientField | medicine_types.MedicineField | example_app_types.PersonNameField | pharmaceutical_company_types.PharmaceuticalCompanyField | example_app_types.PhoneNumberField | product_types.ProductField | user_types.UserField
	typeslib.Field
}
