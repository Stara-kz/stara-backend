package createCityController

type InputCreateCity struct {
	Name string `json:"city" validate:"required"`
}