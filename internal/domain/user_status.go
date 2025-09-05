package domain

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
	Banned   UserStatus = "banned"
	Pending  UserStatus = "pending"
)
