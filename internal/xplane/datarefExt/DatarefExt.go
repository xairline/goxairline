package datarefext

import (
	"fmt"
	"math"
	"xairline/goxairline/internal/xplane/shared"

	"github.com/xairline/goplane/extra/logging"
	"github.com/xairline/goplane/xplm/dataAccess"
)

func NewDataRefExt(name, datarefStr string, precision int8, isBytesArray bool, findDataRef FindDataRef, getDataRefType GetDatarefType, logger *shared.Logger) *DataRefExt {
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

	return &DataRefExt{name: name, dataref: myDataref, datarefType: datarefType, isBytesArray: isBytesArray, precision: &precision}
}

func (datarefExt *DataRefExt) GetName() string {
	return datarefExt.name
}

func (datarefExt *DataRefExt) GetDatarefType() dataAccess.DataRefType {
	return datarefExt.datarefType
}

func (datarefExt *DataRefExt) GetCurrentValue() interface{} {
	var currentValue interface{}
	switch datarefExt.datarefType {
	case dataAccess.TypeInt:
		currentValue = dataAccess.GetIntData(datarefExt.dataref)
	case dataAccess.TypeFloat:
		tmp := dataAccess.GetFloatData(datarefExt.dataref)
		if datarefExt.precision != nil {
			currentValue = dataRoundup(float64(tmp), int(*datarefExt.precision))
		} else {
			currentValue = tmp
		}
	case dataAccess.TypeDouble:
		tmp := dataAccess.GetFloatData(datarefExt.dataref)
		if datarefExt.precision != nil {
			currentValue = dataRoundup(float64(tmp), int(*datarefExt.precision))
		} else {
			currentValue = tmp
		}
	case dataAccess.TypeFloatArray:
		tmpValue := dataAccess.GetFloatArrayData(datarefExt.dataref)
		res := make([]float64, len(tmpValue))
		if datarefExt.precision != nil {
			for index, tmp := range tmpValue {
				res[index] = dataRoundup(float64(tmp), int(*datarefExt.precision))
			}
			currentValue = res
		} else {
			currentValue = tmpValue
		}
	case dataAccess.TypeIntArray:
		currentValue = dataAccess.GetIntArrayData(datarefExt.dataref)
	case dataAccess.TypeData: // string??
		tmpValue := dataAccess.GetData(datarefExt.dataref)
		if datarefExt.isBytesArray {
			currentValue = ""
			for _, element := range tmpValue {
				if element == 0 {
					break
				}
				currentValue = fmt.Sprintf("%s", currentValue) + string(byte(element))
			}
		} else {
			currentValue = tmpValue
		}
	default:
		logging.Infof("Unknown dataref type for %+v", datarefExt)
		return nil
	}
	return currentValue
}

func dataRoundup(value float64, precision int) float64 {
	if precision == -1 {
		return value
	}
	precisionFactor := math.Pow10(precision)
	return math.Round(value*precisionFactor) / precisionFactor
}
