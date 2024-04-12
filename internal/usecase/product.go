package usecase

import (
	"context"
	"fmt"

	"go-clean-template/internal/entity"
)

// ProductUseCase -.
type ProductUseCase struct {
	repo ProductRepo
}

// New -.
func New(r ProductRepo) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
	}
}

// AddProduct - adding product to inventory.
func (uc *ProductUseCase) AddProduct(ctx context.Context, p entity.Product) (entity.Product, error) {
	createdProduct, err := uc.repo.AddProduct(ctx, p)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductUseCase - AddProduct - s.repo.AddProduct: %w", err)
	}

	return createdProduct, nil
}

// GetProduct - getting product from inventory.
func (uc *ProductUseCase) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	product, err := uc.repo.GetProduct(ctx, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductUseCase - GetProduct - s.repo.GetProduct: %w", err)
	}

	return product, nil
}

// DeleteProduct - deleting product from inventory.
func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id string) (entity.Product, error) {
	deletedProduct, err := uc.repo.DeleteProduct(ctx, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductUseCase - DeleteProduct - s.Repo.DeleteProduct: %w", err)
	}

	return deletedProduct, nil
}
