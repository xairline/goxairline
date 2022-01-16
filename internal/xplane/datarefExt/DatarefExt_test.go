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
		tmp := NewDataRefExt("test", "test", test.mockFindDataref, nil, &logger)
		if (tmp == nil) != test.expected {
			t.Fatalf("Output %q not equal to expected %v", tmp, test.expected)
		}
	}
}

func TestNewDataRefExt_Getter(t *testing.T) {
	logger := shared.GetLoggerForTest(t)
	tmp := NewDataRefExt("test", "test", func(datarefStr string) (dataAccess.DataRef, bool) {
		var res dataAccess.DataRef
		return res, true
	}, func(dataref dataAccess.DataRef) dataAccess.DataRefType {
		return dataAccess.TypeDouble
	}, &logger)
	assert.Equal(t, tmp.GetName(), "test")
	assert.Equal(t, tmp.GetDatarefType(), dataAccess.TypeDouble)
}
