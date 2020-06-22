//print "Hello world"
//if you run, change filename "main.go"

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func main(){
	//fmt.Printf("Hello World\n")
	//run 8080 port
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}