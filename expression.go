package createhtml

type Expression struct {
	firstStatement interface{}
	Condition      interface{}
	Operator       Operator
	Expected       string
	Default        string
}

func (e *Expression) Expression() string {
	fsStr, fsInt := StringOrInt(e.firstStatement)
	ssStr, ssInt := StringOrInt(e.Condition)

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
			return e.Equal(fsInt || ssInt)
		case NotEqual:
			return e.NotEqual(fsInt || ssInt)
		case LessOrEqual:
			return e.LessOrEqual(fsInt || ssInt)
		case Less:
			return e.Less(fsInt || ssInt)
		case Greater:
			return e.Greater(fsInt || ssInt)
		case GreaterOrEqual:
			return e.GreaterOrEqual(fsInt || ssInt)
		}
	}

	return e.Default
}

func (e *Expression) Equal(stringOrInt bool) string {
	if stringOrInt {
		if e.firstStatement.(string) == e.Condition.(string) {
			return e.Expected
		}
		return e.Default
	}

	if e.firstStatement.(int) == e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) NotEqual(stringOrInt bool) string {
	if stringOrInt {
		if e.firstStatement.(string) != e.Condition.(string) {
			return e.Expected
		}
		return e.Default
	}

	if e.firstStatement.(int) == e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) LessOrEqual(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) <= e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) Less(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) < e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) GreaterOrEqual(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) >= e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}

func (e *Expression) Greater(stringOrInt bool) string {
	if stringOrInt {
		return e.Default
	}

	if e.firstStatement.(int) > e.Condition.(int) {
		return e.Expected
	}
	return e.Default
}
