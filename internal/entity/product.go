// Package entity defines main entities for business logic (services), database mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

// Product -.
type Product struct {
	ID          string    `json:"id" example:"01HV73P3B06PTPXEMYCCBMM03C"`
	Name        string    `json:"name" example:"Laptop"`
	Description string    `json:"description" example:"MacBook Pro 9999"`
	CategoryID  string    `json:"category_id" example:"01F8Z4P8KQ4P7CSHMSZETEATWX"`
	UnitPrice   float64   `json:"unit_price" example:"99.9"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
