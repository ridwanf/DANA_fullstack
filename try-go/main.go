package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	greet()
	str := generateUUID()
	fmt.Println(str)

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello, temperance"))
	})
	http.ListenAndServe(":80", nil)
}

func greet() {
	fmt.Println("Hello, playground")
}

func generateUUID() string {
	return uuid.NewString()
}
