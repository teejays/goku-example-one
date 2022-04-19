package methods_drug

import (
	"context"
	"fmt"

	"github.com/teejays/clog"

	"github.com/teejays/goku-util/client/db"

	drug_dal "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/dal"
	drug_types "github.com/teejays/goku-example-one/backend/services/pharmacy/drug/goku.generated/types"
)

// DrugEntity_Server provides all the methods that fall under this entity
type DrugEntity_Server struct{}

func NewServer() *DrugEntity_Server {
	return &DrugEntity_Server{}
}

func (s DrugEntity_Server) AddDrug(ctx context.Context, req drug_types.Drug) (drug_types.Drug, error) {
	var resp drug_types.Drug
	var err error

	clog.Infof("[Method] AddDrug() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := drug_dal.DrugEntityDAL{}
	resp, err = d.AddDrug(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s DrugEntity_Server) UpdateDrug(ctx context.Context, req drug_types.UpdateDrugRequest) (drug_types.UpdateDrugResponse, error) {
	var resp drug_types.UpdateDrugResponse
	var err error

	clog.Infof("[Method] UpdateDrug() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := drug_dal.DrugEntityDAL{}
	resp, err = d.UpdateDrug(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s DrugEntity_Server) GetDrug(ctx context.Context, req drug_types.GetDrugRequest) (drug_types.Drug, error) {
	var resp drug_types.Drug
	var err error

	clog.Infof("[Method] GetDrug() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := drug_dal.DrugEntityDAL{}
	resp, err = d.GetDrug(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s DrugEntity_Server) ListDrug(ctx context.Context, req drug_types.ListDrugRequest) (drug_types.ListDrugResponse, error) {
	var resp drug_types.ListDrugResponse
	var err error

	clog.Infof("[Method] ListDrug() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := drug_dal.DrugEntityDAL{}
	resp, err = d.ListDrug(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s DrugEntity_Server) QueryByTextDrug(ctx context.Context, req drug_types.QueryByTextDrugRequest) (drug_types.ListDrugResponse, error) {
	var resp drug_types.ListDrugResponse
	var err error

	clog.Infof("[Method] QueryByTextDrug() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := drug_dal.DrugEntityDAL{}
	resp, err = d.QueryByTextDrug(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}
