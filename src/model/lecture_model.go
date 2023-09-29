package model

type LecturePayload struct {
	Email        string `json:"email" validate:"required;email"`
	Name         string `json:"name" validate:"required"`
	Title        string `json:"title" validate:"required"`
	SchoolName   string `json:"school_name" validate:"school_name"`
	SupervisorId int64  `json:"uspervisor_id" validate:"required"`
}
