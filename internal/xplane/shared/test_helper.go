package shared

import "testing"

func GetLoggerForTest(t *testing.T) Logger {
	return Logger{
		Errorf: t.Logf,
		Infof:  t.Logf,
	}
}
