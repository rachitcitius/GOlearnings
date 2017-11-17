// QueryWebServer
package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"strconv"
)

type apiStruc struct {
	Name string `json:"name"`
	Mass string `json:"mass"`
	Hair_Color string `json:"hair_color"`
}

func main() {
	for i:=1;i<=10;i++ {
		//Calling the API concurrently
		go QueryServer(strconv.Itoa(i))	
	}
	
	time.Sleep(time.Second * 2)
}

func QueryServer(s string) {
	//Calling the free StarWars API
	resp, err := http.Get( fmt.Sprintf("https://swapi.co/api/people/%s/", s) )
	if err!=nil{
		fmt.Println("Error")
	}
	
	defer resp.Body.Close()		//Closing the response once execution of the function done
	var apiS apiStruc
	json.NewDecoder(resp.Body).Decode(&apiS)
	fmt.Println("For ID", s, apiS)
}