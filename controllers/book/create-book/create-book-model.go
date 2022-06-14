package createBookController

type InputCreateBook struct {
	Name string `json:"name" validate:"required"`
	Authors []string `json:"authors" validate:"required"`
	ContentDescription string `json:"content_description"`
	Isbn string `json:"isbn"`
	PhotoUrl string `json:"photo_url"`
	BookConditionDescription string `json:"book_condition_description" validate:"required"`
	Condition uint
	UserID uint `json:"user_id" validate:"required"`
	CityID uint `json:"city_id" validate:"required"`
}

type InputUserBook struct {
	UserID string
	BookID string
}

type InputUserCity struct {
	UserID string
	CityID string
}