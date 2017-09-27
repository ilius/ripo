package restpc

type ParamFlag uint16

func (f ParamFlag) Mandatory() bool {
	return f&Optional == 0
}

func (f ParamFlag) FromBody() bool {
	return f&NotFromBody == 0
}

func (f ParamFlag) FromForm() bool {
	return f&NotFromForm == 0
}

const (
	Optional ParamFlag = 1 << iota
	NotFromBody
	NotFromForm
)

func mergeParamFlags(flags ...ParamFlag) ParamFlag {
	var result ParamFlag
	for _, f := range flags {
		result = result | f
	}
	return result
}
