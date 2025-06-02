package replacejsontag

import (
	"encoding/json"
	"errors"
	"reflect"
)

func Marshal(obj interface{}, tagName string) ([]byte, error) {
	res, err := process(reflect.ValueOf(obj), tagName)
	if err != nil {
		return nil, err
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func process(value reflect.Value, tagName string) (interface{}, error) {
	// 结构体、数组、指针、其余
	switch value.Kind() {
	case reflect.Invalid:
		return nil, nil

	case reflect.Chan, reflect.Func:
		return nil, errors.New("unsupported type: channels and functions cannot be serialized")

	case reflect.Array, reflect.Slice:
		length := value.Len()
		res := make([]interface{}, length)
		for i := 0; i < length; i++ {
			temp, err := process(value.Index(i), tagName)
			if err != nil {
				return nil, err
			}
			res[i] = temp
		}
		return res, nil

	case reflect.Struct:

	case reflect.Map:

	case reflect.Pointer:
		// 不用区分空指针，因为下一层会判断为invalid
		return process(value.Elem(), tagName)

	case reflect.UnsafePointer:
		// unsafepointer 未赋值

	case reflect.Interface:
		// interface 未赋值
		if value.IsValid() == false {
			return nil, nil
		}
		return process(value.Elem(), tagName)

	default:
		return value.Interface(), nil
	}
}
