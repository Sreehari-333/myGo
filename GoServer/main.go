// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 url not found", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello!..")
// }

// func formHandler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "Post request is successful")
// 	name := r.FormValue("name")

// 	fmt.Fprintf(w, "Name is : %s \n ", name)
// }

// func main() {

// 	fileServer := http.FileServer(http.Dir("./static"))

// 	http.Handle("/", fileServer)
// 	http.HandleFunc("/form", formHandler)
// 	http.HandleFunc("/hello", helloHandler)

// 	http.ListenAndServe(":8080", nil)

// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 URL not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello!..")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	fmt.Fprintln(w, "Post request is successful")
	fmt.Fprintf(w, "Name is: %s\n", name)
}

func main() {
	fileServer := http.FileServer(http.Dir("./GoServer/static"))

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", fileServer)

	fmt.Println("Server is running at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
