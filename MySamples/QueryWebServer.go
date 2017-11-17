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

//Global map variable to store the individual response
var finalResponse map[string]string

func main() {
	finalResponse = make(map[string]string)
	http.HandleFunc("/", APIcall)
	http.ListenAndServe(":1234", nil)	
}

func APIcall(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "text/html")
	for i:=1;i<=20;i++ {
		//Calling the API concurrently
		go QueryServer(strconv.Itoa(i))	
	}	
	time.Sleep(time.Second * 2)
	w.Write([]byte("<div>"))
	for a:=range finalResponse {
		w.Write([]byte( a + "->" + finalResponse[a] + "</br>"))
	}
	w.Write([]byte("</div>"))
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
	finalResponse[s] = fmt.Sprint(apiS)	//Appending the response in map
}