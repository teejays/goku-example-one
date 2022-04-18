package methods_pharmaceutical_company

import (
	"context"
	"fmt"

	"github.com/teejays/clog"

	"github.com/teejays/goku/generator/external/client/db"

	pharmaceutical_company_dal "github.com/teejays/goku/example/backend/services/pharmacy/pharmaceutical_company/goku.generated/dal"
	pharmaceutical_company_types "github.com/teejays/goku/example/backend/services/pharmacy/pharmaceutical_company/goku.generated/types"
)

// PharmaceuticalCompanyEntity_Server provides all the methods that fall under this entity
type PharmaceuticalCompanyEntity_Server struct{}

func NewServer() *PharmaceuticalCompanyEntity_Server {
	return &PharmaceuticalCompanyEntity_Server{}
}

func (s PharmaceuticalCompanyEntity_Server) AddPharmaceuticalCompany(ctx context.Context, req pharmaceutical_company_types.PharmaceuticalCompany) (pharmaceutical_company_types.PharmaceuticalCompany, error) {
	var resp pharmaceutical_company_types.PharmaceuticalCompany
	var err error

	clog.Infof("[Method] AddPharmaceuticalCompany() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := pharmaceutical_company_dal.PharmaceuticalCompanyEntityDAL{}
	resp, err = d.AddPharmaceuticalCompany(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s PharmaceuticalCompanyEntity_Server) UpdatePharmaceuticalCompany(ctx context.Context, req pharmaceutical_company_types.UpdatePharmaceuticalCompanyRequest) (pharmaceutical_company_types.UpdatePharmaceuticalCompanyResponse, error) {
	var resp pharmaceutical_company_types.UpdatePharmaceuticalCompanyResponse
	var err error

	clog.Infof("[Method] UpdatePharmaceuticalCompany() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := pharmaceutical_company_dal.PharmaceuticalCompanyEntityDAL{}
	resp, err = d.UpdatePharmaceuticalCompany(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s PharmaceuticalCompanyEntity_Server) GetPharmaceuticalCompany(ctx context.Context, req pharmaceutical_company_types.GetPharmaceuticalCompanyRequest) (pharmaceutical_company_types.PharmaceuticalCompany, error) {
	var resp pharmaceutical_company_types.PharmaceuticalCompany
	var err error

	clog.Infof("[Method] GetPharmaceuticalCompany() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := pharmaceutical_company_dal.PharmaceuticalCompanyEntityDAL{}
	resp, err = d.GetPharmaceuticalCompany(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s PharmaceuticalCompanyEntity_Server) ListPharmaceuticalCompany(ctx context.Context, req pharmaceutical_company_types.ListPharmaceuticalCompanyRequest) (pharmaceutical_company_types.ListPharmaceuticalCompanyResponse, error) {
	var resp pharmaceutical_company_types.ListPharmaceuticalCompanyResponse
	var err error

	clog.Infof("[Method] ListPharmaceuticalCompany() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := pharmaceutical_company_dal.PharmaceuticalCompanyEntityDAL{}
	resp, err = d.ListPharmaceuticalCompany(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s PharmaceuticalCompanyEntity_Server) QueryByTextPharmaceuticalCompany(ctx context.Context, req pharmaceutical_company_types.QueryByTextPharmaceuticalCompanyRequest) (pharmaceutical_company_types.ListPharmaceuticalCompanyResponse, error) {
	var resp pharmaceutical_company_types.ListPharmaceuticalCompanyResponse
	var err error

	clog.Infof("[Method] QueryByTextPharmaceuticalCompany() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := pharmaceutical_company_dal.PharmaceuticalCompanyEntityDAL{}
	resp, err = d.QueryByTextPharmaceuticalCompany(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}
