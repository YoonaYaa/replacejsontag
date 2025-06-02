package main

import (
	"fmt"
	"reflect"
)

type temprature int

func goreflect(value interface{}) {
	v := reflect.ValueOf(value)
	fmt.Println(v.Kind())
}

type o struct {
	q string
}

type u struct {
	o interface{}
}

func PrintType(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
}

func main() {
	v := reflect.ValueOf(u{o{}})
	fmt.Println(v.FieldByName("o"), v.FieldByName("o").Elem().Kind())

	// var q unsafe.Pointer
	// qv := reflect.ValueOf(q)
	// fmt.Println(qv, qv.Kind(), qv.Elem())

	var q *int
	qv := reflect.ValueOf(q)
	fmt.Println(qv, qv.Kind(), qv.Elem(), qv.IsValid(), qv.IsZero(), qv.IsNil())
	if q == nil {
		fmt.Println("yes")
	}

	var i interface{} = (*int)(nil)
	qv = reflect.ValueOf(i)
	fmt.Println(qv.IsValid(), qv.Elem().Kind(), qv.IsZero(), qv.IsNil())

	a := []int{12, 31}
	pa := &a
	rpa := reflect.ValueOf(pa)
	fmt.Println(rpa.IsValid(), rpa.Elem().Kind())
}
