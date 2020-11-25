package convert

import (
	"reflect"
)

type Map map[string]interface{}

func DeepCopy(target, source interface{}) error {
	if byts, err := MarshalJson(source); err != nil {
		return err
	} else {
		return UnMarshalJson(byts, target)
	}
}

func Map2Struct(m map[string]interface{}, val interface{}) error {
	byts, err := MarshalJson(m)
	if err != nil {
		return err
	}
	return UnMarshalJson(byts, val)
}

func Struct2Map(data interface{}) (result map[string]interface{}) {
	for i, el, result := 0, reflect.ValueOf(data).Elem(), map[string]interface{}{}; i < el.NumField(); i++ {
		result[el.Type().Field(i).Name] = el.Field(i).Interface()
	}
	return
}

func Struct2JsonMap(data interface{}) (result map[string]interface{}, _err error) {
	result = map[string]interface{}{}
	if byts, err := MarshalJson(data); err != nil {
		_err = err
		return
	} else {
		_err = UnMarshalJson(byts, &result)
	}
	return
}

func StructTag2Map(data interface{}, tag string) (result map[string]interface{}) {
	for i, el, result := 0, reflect.ValueOf(data).Elem(), map[string]interface{}{}; i < el.NumField(); i++ {
		result[el.Type().Field(i).Tag.Get(tag)] = el.Field(i).Interface()
	}
	return
}
