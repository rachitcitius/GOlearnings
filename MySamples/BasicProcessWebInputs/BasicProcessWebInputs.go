// BasicProcessWebInputs
//Sample to load a HTML page "BasicProcessWebInputsLogin.html"
//Reading the form  inputs & printing on post page
//How to Use : http://localhost:5111/login

package main

import (
	"fmt"
	"net/http"
//	"html/template"
)

func main() {
	ListenServer()
}

func errChk(e error) () {
	if e != nil {
		panic(e)
	}
}

func ListenServer() {
	http.HandleFunc("/login", LoadPage)
	http.HandleFunc("/post", Login)
	http.ListenAndServe(":5111", nil)
}

func LoadPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r,"BasicProcessWebInputsLogin.html")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserName :", r.PostFormValue("username"))
	output := fmt.Sprintf("You Entered \n UserName : %s \n Password : %s", r.PostFormValue("username"), r.PostFormValue("password"))
	fmt.Fprintf(w, output)
}
