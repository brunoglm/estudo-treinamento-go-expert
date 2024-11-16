package main

import (
	"encoding/json"
	"fmt"

	fastJSON "github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	var p fastJSON.Parser
	jsonData := `{"user":{"name": "name1", "age": 20}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	u := v.GetObject("user")

	fmt.Println(u)

	userJSON := v.Get("user").String()

	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		panic(err)
	}

	println(user.Name)
	println(user.Age)
}
