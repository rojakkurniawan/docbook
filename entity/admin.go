package entity

type CreateDoctorRequest struct {
	User           User           `json:"user" binding:"required"`
	Doctor         Doctor         `json:"doctor" binding:"required"`
	Specialization Specialization `json:"specialization" binding:"required"`
}
