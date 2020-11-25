package convert

import (
	"errors"
	"reflect"

	"github.com/optim-kazuhiro-seida/go-advance-type/ref"
)

type Map map[string]interface{}

func DeepCopy(source, target interface{}) error {
	if byts, err := MarshalJson(source); err != nil {
		return err
	} else {
		return UnMarshalJson(byts, target)
	}
}

func CopyFields(source, target interface{}) (err error) {
	targetRef := reflect.ValueOf(target)
	if targetRef.Kind() != reflect.Ptr {
		err = errors.New(" not pointer variable")
		return
	}

	resourceRef := ref.Indirect(source)
	for i := 0; i < resourceRef.Type().NumField(); i++ {
		if field := resourceRef.Type().Field(i); !field.Anonymous {
			name := field.Name
			if srcField, dstField := resourceRef.FieldByName(name), targetRef.Elem().FieldByName(name); srcField.IsValid() &&
				dstField.IsValid() &&
				srcField.Type() == dstField.Type() {
				dstField.Set(srcField)
			}
		}
	}
	return
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
