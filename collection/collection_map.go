package collection

import (
	"reflect"

	ref2 "github.com/optim-kazuhiro-seida/go-advance-type/ref"
)

type _Object struct{}

var Object = _Object{}

func (_Object) GetKeys(m interface{}) (result []string) {
	ref := ref2.Indirect(m)
	if ref.Kind() != reflect.Map {
		return nil
	}
	for _, key := range ref.MapKeys() {
		result = append(result, key.String())
	}
	return nil
}

func (_Object) GetValues(m interface{}) (result []interface{}) {
	ref := ref2.Indirect(m)
	if ref.Kind() != reflect.Map {
		return
	}
	for _, key := range ref.MapKeys() {
		result = append(result, ref.MapIndex(key).Interface())
	}
	return
}
