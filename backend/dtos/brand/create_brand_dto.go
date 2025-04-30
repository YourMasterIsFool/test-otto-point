package branddto

type CreateBrand struct {
	Name string `json:"name" validate:"required"`
}
