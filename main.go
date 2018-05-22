package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server ... at 9999")
	panic(http.ListenAndServe(":9999", http.FileServer(http.Dir("./static"))))
}
