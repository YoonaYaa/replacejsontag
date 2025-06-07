package test

import (
	"fmt"
	"replacejsontag/replacejsontag"
)

type t struct {
	A string `json:"AJson" openapi:"AOpenApi"`
	B bool   `json:"BJson" openapi:"BOpenApi"`
	C int    `json:"CJson" openapi:"COpenApi"`
}

var tag = "openapi"

func Test() {
	temp := t{"A", true, 10}
	res, err := replacejsontag.Marshal(temp, tag)
	if err != nil {
		fmt.Printf("failed: replacejsontag.Marshal, temp: %+v, tag: %s\n", temp, tag)
		return
	}
	fmt.Println(string(res))
}
