package datarefext

import "github.com/xairline/goplane/xplm/dataAccess"

type DataRefExt struct {
	name         string
	dataref      dataAccess.DataRef
	datarefType  dataAccess.DataRefType
	value        interface{}
	precision    *int8
	isBytesArray bool
}

type FindDataRef func(dataRefName string) (dataAccess.DataRef, bool)
type GetDatarefType func(dataref dataAccess.DataRef) dataAccess.DataRefType
