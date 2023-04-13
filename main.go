package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return 
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	/* Fprintf syntax:-
		w =  the output destination where the formatted text will be written.
		format string: a format string that specifies how the values will be printed.
		a ...interface{}: a variadic argument that contains the values to be printed.
	*/
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)

}


//This line declares the function helloHandler which takes in two arguments, a ResponseWriter and a Request object, both of which are from the net/http package.
func helloHandler(w http.ResponseWriter, r *http.Request){

	//This line checks if the requested URL path is not equal to "/hello". If it is not, it sets the error message "404 not found" and returns an HTTP 404 status error code.
	err := r.URL.Path
	if  err != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}

	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w, "hello!")
}

func main(){
	// Taking all the files present in "static folder" and put them on internet.
	fileserver:= http.FileServer(http.Dir("./static"))

	// Telling webserver that whenever our website URL is visited put this server to action.
	http.Handle("/",fileserver)

	// Whenever "/form" will be written in front of our website's URL formHandler function will come into action.
	http.HandleFunc("/form", formHandler)

	// Whenever "/hello" will be written in front of our website's URL helloHandler function wil come into action.
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	// Starts an HTTP server and listens for incoming HTTP requests on a specified network address.
	err:= http.ListenAndServe(":8080",nil)
	
	if err != nil{ 
		log.Fatal(err)
	}




}
/*An HTTP status code is a three-digit numerical code
 that is returned by a web server to indicate the status of a client's request 
 made over the HTTP (Hypertext Transfer Protocol) protocol.*/

//404 Not Found - Requested resource could not be found