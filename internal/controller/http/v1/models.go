package v1

type addProductRequest struct {
	Name        string  `json:"name" example:"Laptop"`
	Description string  `json:"description" example:"MacBook Pro 9999"`
	CategoryID  string  `json:"category_id" example:"01F8Z4P8KQ4P7CSHMSZETEATWX"`
	UnitPrice   float64 `json:"unit_price" example:"99.9"`
}

type getProductRequest struct {
	ID string `json:"id" example:"01HV73P3B06PTPXEMYCCBMM03C"`
}
