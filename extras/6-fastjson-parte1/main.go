package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser
	jsonData := `{"foo":"bar", "num":123, "bool": true, "err": [1,2,3]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("foo=%s\n", v.GetArray("err"))
	array := v.GetArray("err")
	for _, value := range array {
		fmt.Println(value)
	}
}
