package domain

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

func MapToRoleEntity(rolestr string) Role {
	switch rolestr {
	case UserRoleStr:
		return UserRole
	case AdminRoleStr:
		return AdminRole

	}
	return 0
}

func MapFromRoleEntity(role Role) string {
	switch role {
	case UserRole:
		return UserRoleStr
	case AdminRole:
		return AdminRoleStr

	}
	return ""
}
