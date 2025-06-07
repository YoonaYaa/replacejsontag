package main

import (
	"fmt"
	"reflect"
	"replacejsontag/test"
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

type example struct {
	name   string `json:"name" openapi:nil`
	Age    int    `openapi:""`
	height int
}

func PrintType(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
}

func main() {
	e := example{"ja", 1, 12}
	ex := reflect.ValueOf(e)
	t := reflect.TypeOf(e)
	fmt.Println(t)
	fmt.Println(ex)
	fmt.Println(ex.Type(), ex.Type().Field(0), ex.Field(0))

	temp, ok := ex.Type().Field(0).Tag.Lookup("openapi")
	fmt.Println(temp, ok)

	temp, ok = ex.Type().Field(1).Tag.Lookup("openapi")
	fmt.Println(temp, ok)

	temp, ok = ex.Type().Field(2).Tag.Lookup("openapi")
	fmt.Println(temp, ok)

	fmt.Println(ex.Field(0).Type(), ex.Type().Field(0))

	test.Test()
}
