package convert

import (
	"errors"
	"reflect"
	"strings"

	"github.com/optim-kazuhiro-seida/go-advance-type/ref"
)

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
	for i, sourceRef := 0, ref.Indirect(source); i < sourceRef.Type().NumField(); i++ {
		if field := sourceRef.Type().Field(i); !field.Anonymous {
			srcField, dstField := sourceRef.FieldByName(field.Name), targetRef.Elem().FieldByName(field.Name)
			if srcField.IsValid() && dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcField.Type() {
				dstField.Set(srcField)
			}
		}
	}
	return
}
func CopyCastFields(source, target interface{}) (err error) {
	targetRef := reflect.ValueOf(target)
	if targetRef.Kind() != reflect.Ptr {
		err = errors.New(" not pointer variable")
		return
	}
	for i, sourceRef := 0, ref.Indirect(source); i < sourceRef.Type().NumField(); i++ {
		if field := sourceRef.Type().Field(i); !field.Anonymous {
			srcField, dstField := sourceRef.FieldByName(field.Name), targetRef.Elem().FieldByName(field.Name)
			if srcField.IsValid() && dstField.IsValid() && dstField.CanSet() {
				if _type := dstField.Type(); _type == srcField.Type() {
					dstField.Set(srcField)
				} else if _type.String() == "string" {
					dstField.Set(reflect.ValueOf(MustStr(srcField)))
				} else if _type.String() == "int" {
					dstField.Set(reflect.ValueOf(MustInt(srcField)))
				} else if _type.String() == "int8" {
					dstField.Set(reflect.ValueOf(MustInt8(srcField)))
				} else if _type.String() == "int16" {
					dstField.Set(reflect.ValueOf(MustInt16(srcField)))
				} else if _type.String() == "int32" {
					dstField.Set(reflect.ValueOf(MustInt32(srcField)))
				} else if _type.String() == "int64" {
					dstField.Set(reflect.ValueOf(MustInt64(srcField)))
				} else if _type.String() == "float32" {
					dstField.Set(reflect.ValueOf(MustFloat32(srcField)))
				} else if _type.String() == "float64" {
					dstField.Set(reflect.ValueOf(MustFloat64(srcField)))
				} else if _type.String() == "bool" {
					dstField.Set(reflect.ValueOf(MustBool(srcField)))
				} else if _type.String() == "*string" {
					tmp := MustStr(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*int" {
					tmp := MustInt(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*int8" {
					tmp := MustInt8(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*int16" {
					tmp := MustInt16(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*int32" {
					tmp := MustInt32(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*int64" {
					tmp := MustInt64(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*float32" {
					tmp := MustFloat32(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*float64" {
					tmp := MustFloat64(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				} else if _type.String() == "*bool" {
					tmp := MustBool(srcField)
					dstField.Set(reflect.ValueOf(&tmp))
				}
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

func Struct2Map(data interface{}) (result map[string]interface{}) {
	result = map[string]interface{}{}
	for i, el := 0, ref.Indirect(data); i < el.NumField(); i++ {
		result[el.Type().Field(i).Name] = el.Field(i).Interface()
	}
	return
}

func StructTag2Map(data interface{}, tag string) (result map[string]interface{}) {
	result = map[string]interface{}{}
	for i, el := 0, ref.Indirect(data); i < el.NumField(); i++ {
		if key, ok := el.Type().Field(i).Tag.Lookup(tag); ok {
			if keys := strings.Split(key, ","); len(keys) > 0 {
				result[keys[0]] = el.Field(i).Interface()
			}
		}
	}
	return
}

func StructJsonTag2Map(data interface{}) map[string]interface{} {
	return StructTag2Map(data, "json")
}
func StructYamlTag2Map(data interface{}) map[string]interface{} {
	return StructTag2Map(data, "yaml")
}
func StructYmlTag2Map(data interface{}) map[string]interface{} {
	return StructTag2Map(data, "yml")
}
func StructDbTag2Map(data interface{}) map[string]interface{} {
	return StructTag2Map(data, "db")
}
