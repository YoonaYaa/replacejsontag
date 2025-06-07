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
			temp, err := process(value.Index(i), tag)
			if err != nil {
				return nil, err
			}
			res[i] = temp
		}
		return res, nil

	case reflect.Struct:
		// 重新创建新的map[string]interface{}用于返回
		res := make(map[string]interface{})
		fieldCount := value.NumField()
		for i := 0; i < fieldCount; i++ {
			fieldType := value.Type().Field(i)
			fieldValue := value.Field(i)
			valueInTargetTag, ok := getValueInTargetTag(fieldType, fieldValue, tag)
			if ok {
				temp, err := process(value.Field(i), tag)
				if err != nil {
					return nil, err
				}
				res[valueInTargetTag] = temp
			}
		}
		return res, nil

	case reflect.Map:
		res := make(map[string]interface{})
		for _, key := range value.MapKeys() {
			temp, err := process(value.MapIndex(key), tag)
			if err != nil {
				return nil, err
			}
			res[key.String()] = temp
		}
		return res, nil

	case reflect.Pointer:
		// 不用区分空指针，因为下一层会判断为invalid
		return process(value.Elem(), tag)

	case reflect.UnsafePointer:
		// unsafepointer 未赋值
		return process(value.Elem(), tag)

	case reflect.Interface:
		// interface 未赋值
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
func getValueInTargetTag(field reflect.StructField, fieldValue reflect.Value, tag string) (string, bool) {
	if valueOfTag, ok := field.Tag.Lookup(tag); ok {
		valueOfTagSlice := strings.Split(valueOfTag, ",")
		if len(valueOfTagSlice) >= 1 && (valueOfTagSlice[0] == "-" ||
			isTargetValueContainInSlice("omitempty", valueOfTagSlice) &&
				isValueEqualIsValidAndIsZero(fieldValue)) {
			return "", false
		} else {
			return valueOfTag, true
		}
	} else {
		return "", false
	}
}

func isValueEqualIsValidAndIsZero(fieldValue reflect.Value) bool {
	return fieldValue.IsValid() && fieldValue.IsZero()
}

func isTargetValueContainInSlice(targetValue string, valueOfTagSlice []string) bool {
	for _, value := range valueOfTagSlice {
		if value == targetValue {
			return true
		}
	}
	return false
}
