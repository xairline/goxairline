package xplane

import (
	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/dataAccess"
)

type DataRefExt struct {
	name        string
	dataref     dataAccess.DataRef
	datarefType dataAccess.DataRefType
	value       interface{}
}

type FindDataRef func(dataRefName string) (dataAccess.DataRef, bool)
type Logger func(format string, a ...interface{})

func NewDataRefExt(name, datarefStr string, datarefType dataAccess.DataRefType, findDataRef FindDataRef, logger Logger) *DataRefExt {
	// allow mock
	if findDataRef == nil {
		findDataRef = dataAccess.FindDataRef
	}
	if logger == nil {
		logger = logging.Errorf
	}

	myDataref, success := findDataRef(datarefStr)
	if !success {
		logger("Failed to FindDataRef: %s", datarefStr)
		return nil
	}
	return &DataRefExt{name: name, dataref: myDataref, datarefType: datarefType}
}

func (datarefExt *DataRefExt) GetStoredValue() interface{} {
	return datarefExt.value
}

func (datarefExt *DataRefExt) GetCurrentValue() interface{} {
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
