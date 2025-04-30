package transactiondto

type CreateTransactionDto struct {
	BrandId    int    `json:"brand_id" validate:"required,numeric"`
	VocuhersId []uint `json:"vouchers_id" validate:"required,dive,required,numeric"`
}
