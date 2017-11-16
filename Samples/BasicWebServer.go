// webserver
//Building a basic web server which will listen at a particular post & send response
package basicwebserver

import (
	"fmt"
	"time"
	"net/http"
)

func basicwebserver() {
	ListenServer()
}

func ListenServer() {
	http.HandleFunc("/", RespondServer)	//function pointer for responding
	http.ListenAndServe(":1234", nil)	//Listener at specified port
}

func RespondServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte( fmt.Sprintf("<font color='red'>Request Received at %s</font>", time.Now().Format(time.RFC3339Nano) )))	//Writting response
	fmt.Println(r.Form)	//Writing at server end
	fmt.Println("Path:", r.URL.Path)
}
