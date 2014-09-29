package netsh

type Limits interface {
	Check() bool
}

type LimitMinMax struct {
	Min int
	Max int
}

func (l *LimitMinMax) Check(value int) bool {
	if value >= l.Min && value <= l.Max {
		return false
	} else {
		return true
	}
}
