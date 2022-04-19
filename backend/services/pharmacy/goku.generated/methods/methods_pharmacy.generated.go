package methods_pharmacy

import (
	drug_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/methods"
	medicine_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/methods"
	pharmaceutical_company_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/pharmaceutical_company/goku.generated/methods"
	product_methods "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/methods"
)

// PharmacyService_Server provides all the service methods, including methods from its entities.
type PharmacyService_Server struct {
	*drug_methods.DrugEntity_Server
	*medicine_methods.MedicineEntity_Server
	*pharmaceutical_company_methods.PharmaceuticalCompanyEntity_Server
	*product_methods.ProductEntity_Server
}

func NewServer() *PharmacyService_Server {
	return &PharmacyService_Server{
		DrugEntity_Server:                  drug_methods.NewServer(),
		MedicineEntity_Server:              medicine_methods.NewServer(),
		PharmaceuticalCompanyEntity_Server: pharmaceutical_company_methods.NewServer(),
		ProductEntity_Server:               product_methods.NewServer(),
	}
}
