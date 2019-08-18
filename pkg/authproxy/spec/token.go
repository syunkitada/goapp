package spec

type IssueToken struct {
	Name     string `validate:"required;"`
	Password string `validate:"required;"`
}
