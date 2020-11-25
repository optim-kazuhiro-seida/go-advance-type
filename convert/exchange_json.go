package convert

import (
	"bytes"
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

func IndentJson(object string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(object), "", "  ")
	if err != nil {
		return "", err
	}
	return buf.String(), err
}
func MarshalIndentJson(v interface{}) (string, error) {
	j, err := json.MarshalIndent(v, "", "    ")
	return string(j), err
}
func MarshalJson(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}
func UnMarshalJson(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
