package xplane

import (
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/dataAccess"
)

type DataRefExt struct {
	dataref     dataAccess.DataRef
	datarefType dataAccess.DataRefType
	value       interface{}
}

func (datarefExt DataRefExt) GetStoredValue() interface{} {
	return datarefExt.value
}

func (datarefExt DataRefExt) GetCurrentValue() interface{} {
	var currentValue interface{}
	switch datarefExt.datarefType {
	case dataAccess.TypeInt:
		currentValue = dataAccess.GetIntData(datarefExt.dataref)
	case dataAccess.TypeFloat:
		currentValue = dataAccess.GetFloatData(datarefExt.dataref)
	case dataAccess.TypeDouble:
		currentValue = dataAccess.GetDoubleData(datarefExt.dataref)
	case dataAccess.TypeFloatArray:
		currentValue = dataAccess.GetFloatArrayData(datarefExt.dataref)
	case dataAccess.TypeIntArray:
		currentValue = dataAccess.GetIntArrayData(datarefExt.dataref)
	case dataAccess.TypeData: // string??
		currentValue = dataAccess.GetData(datarefExt.dataref)
	default:
		logging.Infof("Unknown dataref type for %+v", datarefExt)
		return nil
	}
	datarefExt.value = currentValue
	return currentValue
}
