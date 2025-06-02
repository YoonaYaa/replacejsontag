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

type u struct {
	o string
}

func PrintType(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
}

func main() {
	PrintType(nil)
	PrintType(u{})
}
