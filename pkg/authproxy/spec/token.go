package spec

type IssueToken struct {
	User     string `validate:"required;"`
	Password string `validate:"required;"`
}

type IssueTokenData struct {
	Token string
}
