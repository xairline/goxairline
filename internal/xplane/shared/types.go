package shared

type Logger struct {
	Errorf func(format string, a ...interface{})
	Infof  func(format string, a ...interface{})
}
