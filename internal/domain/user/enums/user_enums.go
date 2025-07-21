package enums

// UserStatus represents the status of a user
type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusInactive  UserStatus = "inactive"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusDeleted   UserStatus = "deleted"
)

// String returns the string representation of UserStatus
func (s UserStatus) String() string {
	return string(s)
}

// IsValid checks if the user status is valid
func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInactive, UserStatusSuspended, UserStatusDeleted:
		return true
	default:
		return false
	}
}

// GetAllUserStatuses returns all valid user statuses
func GetAllUserStatuses() []UserStatus {
	return []UserStatus{
		UserStatusActive,
		UserStatusInactive,
		UserStatusSuspended,
		UserStatusDeleted,
	}
}

// UserRole represents user roles in the system
type UserRole string

const (
	UserRoleAdmin     UserRole = "admin"
	UserRoleModerator UserRole = "moderator"
	UserRoleUser      UserRole = "user"
	UserRoleGuest     UserRole = "guest"
)

// String returns the string representation of UserRole
func (r UserRole) String() string {
	return string(r)
}

// IsValid checks if the user role is valid
func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleModerator, UserRoleUser, UserRoleGuest:
		return true
	default:
		return false
	}
}
