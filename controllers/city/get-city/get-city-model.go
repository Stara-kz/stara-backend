package getCityController

type InputCity struct {
	ID string `json:"id" validate:"required,uuid"`
}