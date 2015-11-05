
/*

tvjs serve sample

*/


package main

import (
	"fmt"
	"net/http"
)

func main(){

	fmt.Println("JS demos")
	fmt.Println("input ip:8080/test.js at Air Launcher on your Apple TV")

	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}

