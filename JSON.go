// JSON
package main

import (
	"fmt"
	"encoding/json"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int	`json:"page"`
	Fruits []string	`json:"fruits"`
}


func main() {
	JSONfunctions()
}

func JSONfunctions() {
	boolA, _ := json.Marshal(true)
	fmt.Println(string(boolA))
	
	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "mango", "peach"}	}
	
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	
	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "mango", "peach"}	}
	
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}