// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Product -.
	Product interface {
		AddProduct(context.Context, entity.Product) (entity.Product, error)
		GetProduct(context.Context, string) (entity.Product, error)
		//UpdateProductByID(context.Context, int, entity.Product) (entity.Product, error)
		DeleteProduct(context.Context, string) (entity.Product, error)
	}

	// ProductRepo -.
	ProductRepo interface {
		AddProduct(ctx context.Context, t entity.Product) (entity.Product, error)
		GetProduct(context.Context, string) (entity.Product, error)
		//UpdateProductByID(context.Context, int, entity.Product) (entity.Product, error)
		DeleteProduct(context.Context, string) (entity.Product, error)
	}
)
