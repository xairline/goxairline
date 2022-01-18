package datarefext

import (
	"xairline/goxairline/internal/xplane/shared"

	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/dataAccess"
)

func NewDataRefExt(name, datarefStr string, findDataRef FindDataRef, getDataRefType GetDatarefType, logger *shared.Logger) *DataRefExt {
	myDataref, success := findDataRef(datarefStr)
	if !success {
		logger.Errorf("Failed to FindDataRef: %s", datarefStr)
		return nil
	}

	datarefType := getDataRefType(myDataref)
	// handle multiple data type
	if datarefType == 6 {
		datarefType = dataAccess.TypeDouble
	}

	return &DataRefExt{name: name, dataref: myDataref, datarefType: datarefType}
}

func (datarefExt *DataRefExt) GetName() string {
	return datarefExt.name
}

func (datarefExt *DataRefExt) GetDatarefType() dataAccess.DataRefType {
	return datarefExt.datarefType
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
