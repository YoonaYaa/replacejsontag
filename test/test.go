package test

import (
	"encoding/json"
	"fmt"
	"replacejsontag/replacejsontag"
)

type User struct {
	Name string `json:"NameJson" openapi:"NameOpenApi"`
	Age  int    `json:"AgeJson" openapi:"AgeOpenApi"`
}

type t struct {
	A string                 `json:"AJson" openapi:"AOpenApi"`
	B bool                   `json:"BJson" openapi:"BOpenApi"`
	C int                    `json:"CJson" openapi:"COpenApi"`
	D float64                `json:"DJson" openapi:"DOpenApi"`
	E []string               `json:"EJson" openapi:"EOpenApi"`
	F []bool                 `json:"FJson" openapi:"FOpenApi"`
	G []int                  `json:"GJson" openapi:"GOpenApi"`
	H []float64              `json:"HJson" openapi:"HOpenApi"`
	I map[string]interface{} `json:"IJson" openapi:"IOpenApi"`
	J User                   `json:"JJson" openapi:"JOpenApi"`
	K []User                 `json:"KJson" openapi:"KOpenApi"`
}

var tag = "openapi"

func Test() {
	temp := t{
		A: "A",
		B: true,
		C: 10,
		D: 10.1,
		E: []string{"ha", "ha"},
		F: []bool{true, false},
		G: []int{1, 2, 3},
		H: []float64{1.1, 2.1, 3.1},
		I: map[string]interface{}{"name": "dog", "age": 100},
		J: User{"dog", 100},
		K: []User{{"dog", 100}, {"Cat", 200}},
	}
	res, err := replacejsontag.Marshal(temp, tag)
	if err != nil {
		fmt.Printf("failed: replacejsontag.Marshal, temp: %+v, tag: %s\n", temp, tag)
		return
	}
	fmt.Println("string(res): ", string(res))

	var tempMap map[string]interface{}
	err = json.Unmarshal(res, &tempMap)
	if err != nil {
		fmt.Printf("failed: replacejsontag.Unmarshal, temp: %+v, tag: %s\n", temp, tag)
		return
	}
	fmt.Println("tempMap: ", tempMap)
}
