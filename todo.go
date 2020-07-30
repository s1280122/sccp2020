package main

import(
"fmt"
"net/http"
)

func main() {

    http.HandleFunc("/todo",func(w http.ResponseWriter, r *http.Request){
        switch r.Method {
        case "GET":
	    fmt.Fprint(w, "GET/todo\n")
        case "POST":
	    fmt.Fprint(w, "POST/todo\n")
        }
    })
    http.ListenAndServe(":8080",nil)
}