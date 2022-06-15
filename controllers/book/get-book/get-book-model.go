package getBookController

type InputGetBook struct {
	ID string `json:"id" validate:"required,uuid"`
}