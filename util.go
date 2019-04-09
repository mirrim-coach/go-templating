package template

import (
	"fmt"
	// util "github.com/Pairin/go-util"
	"math"
	"reflect"
	"strconv"
)

var floatType = reflect.TypeOf(float64(0))
var stringType = reflect.TypeOf("")

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		return strconv.ParseFloat(i, 64)
	default:
		v := reflect.ValueOf(unk)
		if !v.IsValid() {
			return math.NaN(), fmt.Errorf("Can't convert %v to float64", unk)
		}

		v = reflect.Indirect(v)
		if v.Type().ConvertibleTo(floatType) {
			fv := v.Convert(floatType)
			return fv.Float(), nil
		} else if v.Type().ConvertibleTo(stringType) {
			sv := v.Convert(stringType)
			s := sv.String()
			return strconv.ParseFloat(s, 64)
		} else {
			return math.NaN(), fmt.Errorf("Can't convert %v to float64", v.Type())
		}
	}
}

func iToiSlice(input interface{}) []interface{} {
	s := reflect.ValueOf(input)

	if s.Kind() != reflect.Slice {
		return []interface{}{}
	}

	arr := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		arr[i] = s.Index(i).Interface()
	}

	return arr
}
func uniqueArray(a interface{}, b interface{}) bool {
	aArr := iToiSlice(a)
	bArr := iToiSlice(b)

	for x, aV := range aArr {
		for y, bV := range bArr {
			if reflect.DeepEqual(aV, bV) {
				aArr[x] = nil
				bArr[y] = nil
				break
			}
		}
	}

	for _, aVL := range aArr {
		if aVL != nil {
			return false
		}
	}
	for _, bVL := range bArr {
		if bVL != nil {
			return false
		}
	}
	return true
}

func interfaceReflection(data []interface{}) []map[string]interface{} {
	newmap := make([]map[string]interface{}, len(data))

	for i, k := range data {
		v := reflect.ValueOf(k)
		if v.Kind() == reflect.Map {
			newmap[i] = map[string]interface{}{}
			for _, key := range v.MapKeys() {
				newmap[i][key.String()] = v.MapIndex(key).Interface()
			}
		}
	}
	return newmap
}
