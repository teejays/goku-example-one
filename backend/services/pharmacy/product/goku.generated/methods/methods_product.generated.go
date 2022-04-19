package methods_product

import (
	"context"
	"fmt"

	"github.com/teejays/clog"

	"github.com/teejays/goku-util/client/db"

	product_dal "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/dal"
	product_types "github.com/teejays/goku-example-one/backend/services/pharmacy/product/goku.generated/types"
)

// ProductEntity_Server provides all the methods that fall under this entity
type ProductEntity_Server struct{}

func NewServer() *ProductEntity_Server {
	return &ProductEntity_Server{}
}

func (s ProductEntity_Server) AddProduct(ctx context.Context, req product_types.Product) (product_types.Product, error) {
	var resp product_types.Product
	var err error

	clog.Infof("[Method] AddProduct() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := product_dal.ProductEntityDAL{}
	resp, err = d.AddProduct(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s ProductEntity_Server) UpdateProduct(ctx context.Context, req product_types.UpdateProductRequest) (product_types.UpdateProductResponse, error) {
	var resp product_types.UpdateProductResponse
	var err error

	clog.Infof("[Method] UpdateProduct() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := product_dal.ProductEntityDAL{}
	resp, err = d.UpdateProduct(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s ProductEntity_Server) GetProduct(ctx context.Context, req product_types.GetProductRequest) (product_types.Product, error) {
	var resp product_types.Product
	var err error

	clog.Infof("[Method] GetProduct() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := product_dal.ProductEntityDAL{}
	resp, err = d.GetProduct(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s ProductEntity_Server) ListProduct(ctx context.Context, req product_types.ListProductRequest) (product_types.ListProductResponse, error) {
	var resp product_types.ListProductResponse
	var err error

	clog.Infof("[Method] ListProduct() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := product_dal.ProductEntityDAL{}
	resp, err = d.ListProduct(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}

func (s ProductEntity_Server) QueryByTextProduct(ctx context.Context, req product_types.QueryByTextProductRequest) (product_types.ListProductResponse, error) {
	var resp product_types.ListProductResponse
	var err error

	clog.Infof("[Method] QueryByTextProduct() starting with Request\n%+v", req)
	conn, err := db.NewConnection("pharmacy")
	if err != nil {
		return resp, fmt.Errorf("Fetching DB Connection for '%s': %w", "pharmacy", err)
	}
	// Get the DAL wrapper
	d := product_dal.ProductEntityDAL{}
	resp, err = d.QueryByTextProduct(ctx, conn, req)
	if err != nil {
		return resp, fmt.Errorf("DAL method: %w", err)
	}

	return resp, err
}
