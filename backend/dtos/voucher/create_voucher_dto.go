package voucherdto

type CreateVoucherDto struct {
	BrandId uint    `json:"brand_id" validate:"required,numeric"`
	Point   float64 `json:"point" validate:"required,numeric"`
}
