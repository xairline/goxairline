package datarefext

import "github.com/xairline/goplane/xplm/dataAccess"

type DataRefExt struct {
	name        string
	dataref     dataAccess.DataRef
	datarefType dataAccess.DataRefType
	value       interface{}
}

type FindDataRef func(dataRefName string) (dataAccess.DataRef, bool)
type GetDatarefType func(dataref dataAccess.DataRef) dataAccess.DataRefType

type DataRefExtStore map[string]DataRefExt
