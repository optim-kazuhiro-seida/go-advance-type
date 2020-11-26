package check

import (
	"encoding/json"
	"io"
	"reflect"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"
)

func AreEqualJson(data1, data2 interface{}) bool {
	if reflect.DeepEqual(data1, data2) {
		return true
	}
	var _data1, _data2 interface{}

	switch v := data1.(type) {
	case io.Reader:
		if json.NewDecoder(v).Decode(&_data1) != nil {
			return false
		}
	case string:
		if convert.UnMarshalJson(convert.Str2Bytes(v), &_data1) != nil {
			return false
		}
	case []byte:
		if convert.UnMarshalJson(v, &_data1) != nil {
			return false
		}
	default:
		if byts, err := convert.MarshalJson(data1); err != nil {
			return false
		} else if convert.UnMarshalJson(byts, &_data1) != nil {
			return false
		}
	}

	switch v := data2.(type) {
	case io.Reader:
		if json.NewDecoder(v).Decode(&_data2) != nil {
			return false
		}
	case string:
		if convert.UnMarshalJson(convert.Str2Bytes(v), &_data2) != nil {
			return false
		}
	case []byte:
		if convert.UnMarshalJson(v, &_data2) != nil {
			return false
		}
	default:
		if byts, err := convert.MarshalJson(data2); err != nil {
			return false
		} else if convert.UnMarshalJson(byts, &_data2) != nil {
			return false
		}
	}
	return reflect.DeepEqual(_data1, _data2)
}
