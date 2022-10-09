package model

type Role string

const (
	Manager    Role = "MANAGER"
	Technician Role = "TECHNICIAN"
)

func (r *Role) IsValid() bool {
	return *r == Manager || *r == Technician
}
