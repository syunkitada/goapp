package spec

type UpdateUserPassword struct {
	CurrentPassword    string `validate:"required"`
	NewPassword        string `validate:"required"`
	NewPasswordConfirm string `validate:"required"`
}

type UpdateUserPasswordData struct {
}

type User struct {
	Name     string
	RoleName string
}

type GetProjectUsers struct {
}

type GetProjectUsersData struct {
	Users []User
}
