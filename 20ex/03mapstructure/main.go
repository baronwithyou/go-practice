package main

import (
	"encoding/json"
	"fmt"
)

type s struct {
	Id   int
	Name string
}

func main() {
	m := map[string]string{
		"id":   "123123",
		"name": "baron",
	}

	var res s

	mj, _ := json.Marshal(m)

	if err := json.Unmarshal(mj, &res); err != nil {
		panic(err)
	}

	fmt.Println(res)
}
