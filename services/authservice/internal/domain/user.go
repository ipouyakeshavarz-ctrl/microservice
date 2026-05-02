package domain

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	Password    string
	Role        Role
}

type Role uint8

const (
	UserRole Role = iota + 1
	AdminRole
)
const (
	UserRoleStr  = "user"
	AdminRoleStr = "admin"
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return UserRoleStr
	case AdminRole:
		return AdminRoleStr
	}
	return ""
}

func MapToRoleEntity(roleStr string) Role {
	switch roleStr {
	case UserRoleStr:
		return UserRole
	case AdminRoleStr:
		return AdminRole

	}
	return 0
}
