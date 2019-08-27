package spec

type IssueToken struct {
	User     string `validate:"required;"`
	Password string `validate:"required;"`
}

type IssueTokenData struct {
	Token string
}

type Login struct {
	User     string `validate:"required;"`
	Password string `validate:"required;"`
}

type LoginData struct {
	Token     string
	Name      string
	Authority UserAuthority
}

type UserAuthority struct {
	ServiceMap           map[string]uint
	ProjectServiceMap    map[string]ProjectService
	ActionProjectService ProjectService
}

type ProjectService struct {
	RoleID          uint
	RoleName        string
	ProjectName     string
	ProjectRoleID   uint
	ProjectRoleName string
	ServiceMap      map[string]uint
}
