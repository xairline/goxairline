//go:build darwin || freebsd || linux || netbsd || openbsd

package xplane

import (
	"testing"

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
	for _, test := range testCases {
		tmp := NewDataRefExt("test", "test", dataAccess.TypeDouble, test.mockFindDataref, t.Logf)
		if (tmp == nil) != test.expected {
			t.Fatalf("Output %q not equal to expected %v", tmp, test.expected)
		}
	}
}

func TestNewDataRefExt_Getter(t *testing.T) {
	tmp := NewDataRefExt("test", "test", dataAccess.TypeDouble, func(datarefStr string) (dataAccess.DataRef, bool) {
		var res dataAccess.DataRef
		return res, true
	}, t.Logf)
	assert.Equal(t, tmp.GetName(), "test")
	assert.Equal(t, tmp.GetDatarefType(), dataAccess.TypeDouble)
}
