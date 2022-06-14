package getBookController

type InputGetBook struct {
	UserID string `json:"user_id" validate:"required,uuid"`
	ID string `json:"id" validate:"required,uuid"`
}