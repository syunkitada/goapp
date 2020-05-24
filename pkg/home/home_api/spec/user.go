package spec

type UpdateUserPassword struct {
	CurrentPassword    string `validate:"required"`
	NewPassword        string `validate:"required"`
	NewPasswordConfirm string `validate:"required"`
}

type UpdateUserPasswordData struct {
}
