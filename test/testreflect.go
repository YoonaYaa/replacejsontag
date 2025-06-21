package test

import (
	"fmt"
	"reflect"
)

// 这个文件用于测试reflect机制，因为对于reflect不熟悉
// 对于功能没有作用

func TestReflectElem() {
	var intPointer *int = nil
	valueOfIntPointer := reflect.ValueOf(intPointer)
	fmt.Println(valueOfIntPointer.IsValid()) // true，因为有具体类型
	fmt.Println(valueOfIntPointer.Elem())    // invalid reflect.Value，因为具体值是zero value

	// unsafePointer := unsafe.Pointer(intPointer)
	// valueOfUnsafePointer := reflect.ValueOf(unsafePointer)
	// fmt.Println(valueOfUnsafePointer.IsValid()) // true，因为有具体类型
	// fmt.Println(valueOfUnsafePointer.Elem())    // panic，因为Elem不支持unsafe.Pointer

	// var nilInterface interface{}
	// valueOfNilInterface := reflect.ValueOf(nilInterface)
	// fmt.Println(valueOfNilInterface.IsValid()) // false，因为没有具体类型
	// fmt.Println(valueOfNilInterface.Elem())    // panic，因为没有具体类型，所以panic

	var IntPointerInterface interface{} = intPointer
	valueOfIntPointerInterface := reflect.ValueOf(IntPointerInterface)
	fmt.Println(valueOfIntPointerInterface.IsValid()) // true，因为有具体类型
	fmt.Println(valueOfIntPointerInterface.Elem())    // invalid reflect.Value，因为具体值是zero value
}

func TestReflectIsValidAndIsZero() {
	var intPointer *int = nil
	valueOfIntPointer := reflect.ValueOf(intPointer)
	fmt.Println(valueOfIntPointer.IsValid()) // true，因为有具体类型
	fmt.Println(valueOfIntPointer.IsZero())  // true，因为值是nil

	var intZero int = 0
	valueOfintZero := reflect.ValueOf(intZero)
	fmt.Println(valueOfintZero.IsValid()) // true，因为有具体类型
	fmt.Println(valueOfintZero.IsZero())  // true，因为值是0

	var intOne int = 1
	valueOfintOne := reflect.ValueOf(intOne)
	fmt.Println(valueOfintOne.IsValid()) // true，因为有具体类型
	fmt.Println(valueOfintOne.IsZero())  // false，因为值是1
}
