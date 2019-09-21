package spec_model

type Cmd struct {
	Arg         string
	ArgType     string
	ArgKind     string
	FlagMap     map[string]Flag
	TableHeader []string
	Help        string
}

type Flag struct {
	Name      string
	FlagName  string
	ShortName string
	FlagType  string
	FlagKind  string
	CobraType string
	Required  bool
}
