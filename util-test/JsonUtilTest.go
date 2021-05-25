package util_test

import (
	jsoniter "github.com/json-iterator/go"
)

func MarshalTest(values []string) (string, error){
	jsonArray, err := jsoniter.Marshal(values)
	if nil == err {
		return string(jsonArray), nil
	}
	return "", err
}
