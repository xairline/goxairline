//go:build darwin || freebsd || linux || netbsd || openbsd

package datarefext

import (
	"testing"
	"xairline/goxairline/internal/xplane/shared"

	"github.com/stretchr/testify/assert"
	"github.com/xairline/goplane/xplm/dataAccess"
)

type NewDataRefExtTestCases struct {
	mockFindDataref FindDataRef
	expected        bool
}

var testCases = []NewDataRefExtTestCases{
	{mockFindDataref: func(datarefStr string) (dataAccess.DataRef, bool) {
		var res dataAccess.DataRef
		return res, false
	}, expected: true},
	{mockFindDataref: func(datarefStr string) (dataAccess.DataRef, bool) {
		var res dataAccess.DataRef
		return res, true
	}, expected: false},
}

func TestNewDataRefExt(t *testing.T) {
	logger := shared.GetLoggerForTest(t)
	for _, test := range testCases {
		tmp := NewDataRefExt("test", "test", 2, false, test.mockFindDataref, func(dataref dataAccess.DataRef) dataAccess.DataRefType { return dataAccess.TypeDouble }, &logger)
		if (tmp == nil) != test.expected {
			t.Fatalf("Output %v not equal to expected %v", tmp, test.expected)
		}
	}
}

func TestNewDataRefExt_Getter(t *testing.T) {
	logger := shared.GetLoggerForTest(t)
	tmp := NewDataRefExt("test", "test", 2, false, func(datarefStr string) (dataAccess.DataRef, bool) {
		var res dataAccess.DataRef
		return res, true
	}, func(dataref dataAccess.DataRef) dataAccess.DataRefType {
		return dataAccess.TypeDouble
	}, &logger)
	assert.Equal(t, tmp.GetName(), "test")
	assert.Equal(t, tmp.GetDatarefType(), dataAccess.TypeDouble)
}

func TestDataRoundup(t *testing.T) {
	assert.Equal(t, 1.2345, dataRoundup(1.2345, -1))
	assert.Equal(t, float64(0), dataRoundup(float64(0), 0))
	assert.Equal(t, 1.0, dataRoundup(1.2345, 0))
	assert.Equal(t, 1.2, dataRoundup(1.2345, 1))
	assert.Equal(t, 1.23, dataRoundup(1.2345, 2))
	assert.Equal(t, 1.235, dataRoundup(1.2345, 3))
}
