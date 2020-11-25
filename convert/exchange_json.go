package convert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/optim-kazuhiro-seida/go-advance-type/check"

	jsoniter "github.com/json-iterator/go"
)

func MustIndentJson(data interface{}) (result string) {
	result, _ = IndentJson(data)
	return
}
func MustCompactJSON(data interface{}) (result string) {
	result, _ = CompactJSON(data)
	return
}
func MustMarshalJson(v interface{}) (result []byte) {
	result, _ = jsoniter.Marshal(v)
	return
}
func IndentJson(data interface{}) (result string, err error) {
	var buf bytes.Buffer
	switch v := data.(type) {
	case reflect.Value:
		return IndentJson(v.Interface())
	case []byte:
		if err = json.Indent(&buf, v, "", "  "); err == nil {
			result = buf.String()
		}
	case string:
		if err = json.Indent(&buf, Str2Bytes(v), "", "  "); err == nil {
			result = buf.String()
		}
	default:
		if byts, _err := MarshalJson(data); _err != nil {
			err = _err
		} else if err = json.Indent(&buf, byts, "", "  "); err == nil {
			result = buf.String()
		}
	}
	return
}

func CompactJSON(data interface{}) (result string, err error) {
	var buf bytes.Buffer
	switch v := data.(type) {
	case reflect.Value:
		return CompactJSON(v.Interface())
	case []byte:
		if err = json.Compact(&buf, v); err == nil {
			result = buf.String()
		}
	case string:
		if err = json.Compact(&buf, Str2Bytes(v)); err == nil {
			result = buf.String()
		}
	default:
		if byts, _err := MarshalJson(data); _err != nil {
			err = _err
		} else if err = json.Compact(&buf, byts); err == nil {
			result = buf.String()
		}
	}
	return
}
func MarshalJson(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}
func UnMarshalJson(data interface{}, target interface{}) error {
	if !check.IsPtr(target) {
		return errors.New(fmt.Sprintf("not pointer %v", target))
	}
	switch v := data.(type) {
	case reflect.Value:
		return UnMarshalJson(v.Interface(), target)
	case []byte:
		return jsoniter.Unmarshal(v, target)
	case string:
		return jsoniter.Unmarshal(Str2Bytes(v), target)
	default:
		byts, err := MarshalJson(data)
		if err != nil {
			return err
		}
		return UnMarshalJson(byts, target)
	}
}
