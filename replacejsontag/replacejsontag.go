package replacejsontag

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

func Marshal(obj interface{}, tag string) ([]byte, error) {
	res, err := process(reflect.ValueOf(obj), tag)
	if err != nil {
		return nil, err
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func process(value reflect.Value, tag string) (interface{}, error) {
	switch value.Kind() {
	case reflect.Invalid:
		return nil, nil

	case reflect.Chan, reflect.Func:
		return nil, errors.New("unsupported type: channels and functions cannot be serialized")

	case reflect.Array, reflect.Slice:
		length := value.Len()
		res := make([]interface{}, length)
		for i := 0; i < length; i++ {
			valueAfterProcess, err := process(value.Index(i), tag)
			if err != nil {
				return nil, err
			}
			res[i] = valueAfterProcess
		}
		return res, nil

	case reflect.Struct:
		res := make(map[string]interface{})
		numField := value.NumField()
		for i := 0; i < numField; i++ {
			fieldType := value.Type().Field(i)
			fieldValue := value.Field(i)

			newFieldName, ok := getValueOfTargetTag(fieldType, fieldValue, tag)
			if ok {
				temp, err := process(value.Field(i), tag)
				if err != nil {
					return nil, err
				}
				res[newFieldName] = temp
			}
		}
		return res, nil

	case reflect.Map:
		res := make(map[string]interface{})
		for _, key := range value.MapKeys() {
			valueAfterProcess, err := process(value.MapIndex(key), tag)
			if err != nil {
				return nil, err
			}
			res[key.String()] = valueAfterProcess
		}
		return res, nil

	case reflect.Pointer:
		if value.IsValid() == false {
			return nil, nil
		}
		return process(value.Elem(), tag)

	case reflect.UnsafePointer:
		return process(value.Elem(), tag)

	case reflect.Interface:
		if value.IsValid() == false {
			return nil, nil
		}
		return process(value.Elem(), tag)

	default:
		return value.Interface(), nil
	}
}

// 获取key
// 如果tag不存在，那么就跳过
// 如果tag存在，但是值为空或不符合格式，那么就跳过
// 如果tag存在，但是值为-，那么就跳过
// 如果tag存在，但是值包含了omitempty且字段的值为默认值，那么就跳过
// 如果tag存在，值包含了多个，那么默认取出第一个
func getValueOfTargetTag(structField reflect.StructField, fieldValue reflect.Value, tag string) (string, bool) {
	if valueOfTag, ok := structField.Tag.Lookup(tag); ok {
		valueSliceOfTag := strings.Split(valueOfTag, ",")
		if len(valueSliceOfTag) >= 1 && (valueSliceOfTag[0] == "-" ||
			isTargetValueContainInSlice("omitempty", valueSliceOfTag) &&
				isValueIsValidAndIsZero(fieldValue)) {
			return "", false
		} else {
			return valueOfTag, true
		}
	} else {
		return "", false
	}
}

func isValueIsValidAndIsZero(fieldValue reflect.Value) bool {
	return fieldValue.IsValid() && fieldValue.IsZero()
}

func isTargetValueContainInSlice(targetValue string, valueSlice []string) bool {
	for _, value := range valueSlice {
		if value == targetValue {
			return true
		}
	}
	return false
}
