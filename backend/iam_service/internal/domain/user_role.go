package domain

type UserRole int

const (
	Manager UserRole = iota
	Member
)

func (s UserRole) String() string {
	return [...]string{"manager", "member"}[s]
}