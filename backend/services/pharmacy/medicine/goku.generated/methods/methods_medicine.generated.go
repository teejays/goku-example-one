package methods_medicine

import (
	"context"
	"fmt"

	"github.com/teejays/clog"

	"github.com/teejays/goku-util/client/db"

	medicine_dal "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/dal"
	medicine_types "github.com/teejays/goku-example-one/backend/services/pharmacy/medicine/goku.generated/types"
)

// MedicineEntity_Server provides all the methods that fall under this entity
type MedicineEntity_Server struct{}

func NewServer() *MedicineEntity_Server {
	return &MedicineEntity_Server{}
}

func (s MedicineEntity_Server) AddMedicine(ctx context.Context, req medicine_types.Medicine) (medicine_types.Medicine, error) {
	var resp medicine_types.Medicine
	var err error

	clog.Infof("[Method] AddMedicine() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := medicine_dal.MedicineEntityDAL{}
	resp, err = d.AddMedicine(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s MedicineEntity_Server) UpdateMedicine(ctx context.Context, req medicine_types.UpdateMedicineRequest) (medicine_types.UpdateMedicineResponse, error) {
	var resp medicine_types.UpdateMedicineResponse
	var err error

	clog.Infof("[Method] UpdateMedicine() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := medicine_dal.MedicineEntityDAL{}
	resp, err = d.UpdateMedicine(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s MedicineEntity_Server) GetMedicine(ctx context.Context, req medicine_types.GetMedicineRequest) (medicine_types.Medicine, error) {
	var resp medicine_types.Medicine
	var err error

	clog.Infof("[Method] GetMedicine() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := medicine_dal.MedicineEntityDAL{}
	resp, err = d.GetMedicine(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s MedicineEntity_Server) ListMedicine(ctx context.Context, req medicine_types.ListMedicineRequest) (medicine_types.ListMedicineResponse, error) {
	var resp medicine_types.ListMedicineResponse
	var err error

	clog.Infof("[Method] ListMedicine() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := medicine_dal.MedicineEntityDAL{}
	resp, err = d.ListMedicine(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s MedicineEntity_Server) QueryByTextMedicine(ctx context.Context, req medicine_types.QueryByTextMedicineRequest) (medicine_types.ListMedicineResponse, error) {
	var resp medicine_types.ListMedicineResponse
	var err error

	clog.Infof("[Method] QueryByTextMedicine() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := medicine_dal.MedicineEntityDAL{}
	resp, err = d.QueryByTextMedicine(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}
