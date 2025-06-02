package replacejsontag

import (
	"encoding/json"
	"errors"
	"reflect"
)

func Marshal(obj interface{}, tagName string) ([]byte, error) {
	res, err := process(obj, tagName)
	if err != nil {
		return nil, err
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func process(obj reflect.Value, tagName string) (interface{}, error) {
	// 结构体、数组、指针、其余
	switch obj.Kind() {
	case reflect.Invalid:
		return nil, nil

	case reflect.Chan, reflect.Func:
		return nil, errors.New("unsupported type: channels and functions cannot be serialized")

	case reflect.Array, reflect.Slice:

	case reflect.Struct:

	case reflect.Map:

	case reflect.Pointer:

	case reflect.UnsafePointer:

	case reflect.Interface:

	default:

	}
}
