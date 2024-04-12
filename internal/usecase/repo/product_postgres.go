package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"

	"go-clean-template/internal/entity"
	"go-clean-template/pkg/postgres"
)

const _defaultEntityCap = 64

// PostgresRepo -.
type PostgresRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}

// AddProduct adds a new product to the database and returns the added product
func (r *PostgresRepo) AddProduct(ctx context.Context, t entity.Product) (entity.Product, error) {
	sql, args, err := r.Builder.
		Insert("products").
		Columns("id", "name", "description", "category_id", "unit_price").
		Values(t.ID, t.Name, t.Description, t.CategoryID, t.UnitPrice).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - AddProduct - r.Builder: %w", err)
	}

	sql += "RETURNING id, name, description, category_id, unit_price"

	var addedProduct entity.Product
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&addedProduct.ID, &addedProduct.Name, &addedProduct.Description, &addedProduct.CategoryID, &addedProduct.UnitPrice)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - AddProduct - r.Pool.QueryRow: %w", err)
	}

	return addedProduct, nil
}

func (r *PostgresRepo) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	sql, args, err := r.Builder.
		Select("id", "name", "description", "category_id", "unit_price", "created_at", "updated_at").
		From("products").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - GetProduct - r.Builder: %w", err)
	}

	var product entity.Product
	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.UnitPrice, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - GetProduct - r.Pool.QueryRow: %w", err)
	}

	return product, nil
}

func (r *PostgresRepo) DeleteProduct(ctx context.Context, id string) (entity.Product, error) {
	sql, args, err := r.Builder.
		Delete("products").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - DeleteProduct - r.Builder: %w", err)
	}

	var deletedProduct entity.Product
	sql = "RETURNING id, name, description, category_id, unit_price, created_at updated_at"
	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(&deletedProduct.ID, &deletedProduct.Name, &deletedProduct.Description, &deletedProduct.CategoryID, &deletedProduct.UnitPrice, &deletedProduct.CreatedAt, &deletedProduct.UpdatedAt); err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - DeleteProduct - r.Pool.QueryRow: %w", err)
	}

	return deletedProduct, nil
}
