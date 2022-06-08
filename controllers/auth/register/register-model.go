package registerAuthController

type InputUserRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type InputOrganizationRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type InputModeratorRegister struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	OrganizationID uint `json:"organization_id" validate:"required"`
}