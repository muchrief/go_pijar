package model

type Role string

const (
	Student Role = "student"
	Lecture Role = "lecture"
)

type AddUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     Role   `json:"role" validate:"required"`
}

type CheckEmailPayload struct {
	Email string `json:"email" validate:"required,email"`
}
