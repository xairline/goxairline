package datarefext

import "github.com/xairline/goplane/xplm/dataAccess"

type DataRefExt struct {
	name        string
	dataref     dataAccess.DataRef
	datarefType dataAccess.DataRefType
	value       interface{}
}

type FindDataRef func(dataRefName string) (dataAccess.DataRef, bool)
type Logger func(format string, a ...interface{})

type DataRefExtStore map[string]DataRefExt
