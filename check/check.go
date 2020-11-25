package check

import (
	"reflect"
)

func IsBool(val interface{}) bool {
	return val == (0 == 0) || val == (0 != 0)
}

func IsNil(i interface{}) (result bool) {
	if i == nil {
		result = true
	}
	ref := reflect.ValueOf(i)
	result = ref.Kind() == reflect.Ptr && ref.IsNil()
	return
}

func IsPtr(in interface{}) bool {
	return reflect.ValueOf(in).Kind() != reflect.Ptr
}

func IsNumber(in interface{}) bool {
	switch in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}
	return false
}
