package createhtml

type Expression struct {
	firstStatement  interface{}
	SecondStatement interface{}
	Operator        Operator
	Expected        string
	Default         string
}

func (e *Expression) Expression() string {
	fsStr, fsInt := StringOrInt(e.firstStatement)
	ssStr, ssInt := StringOrInt(e.SecondStatement)

	if fsStr && ssStr {
		switch e.Operator {
		case Equal:
			return e.Equal(fsStr || ssStr)
		case NotEqual:
			return e.NotEqual(fsStr || ssStr)
		}
	}

	if fsInt && ssInt {
		switch e.Operator {
		case Equal:
			return e.Equal(fsStr || ssStr)
		case NotEqual:
			return e.NotEqual(fsStr || ssStr)
		case LessOrEqual:
			return e.LessOrEqual(fsStr || ssStr)
		case Less:
			return e.Less(fsStr || ssStr)
		case Greater:
			return e.Greater(fsStr || ssStr)
		case GreaterOrEqual:
			return e.GreaterOrEqual(fsStr || ssStr)
		}
	}

	return e.Default
}

func (e *Expression) Equal(stringOrInt bool) string {
	if stringOrInt {
		if e.firstStatement.(string) == e.SecondStatement.(string) {
			return e.Expected
		}
		return e.Default
	}

	if e.firstStatement.(int) == e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) NotEqual(stringOrInt bool) string {
	if stringOrInt {
		if e.firstStatement.(string) != e.SecondStatement.(string) {
			return e.Expected
		}
		return e.Default
	}

	if e.firstStatement.(int) == e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) LessOrEqual(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) <= e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) Less(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) < e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) GreaterOrEqual(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) >= e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) Greater(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) > e.SecondStatement.(int) {
		return e.Expected
	}
	return e.Default
}
