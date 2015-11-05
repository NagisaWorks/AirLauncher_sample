
/*

static xml file server sample

*/


package main

import (
	"net/http"
	"fmt"
	"path"
	"text/template"
)

type PHandler struct {
	fileHandler http.Handler
}

func (f PHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)

	if r.URL.Path == "/" {
		// force redirect / to index.xml
		w.Header().Set("Location", "index.xml")
		w.WriteHeader(302)
		return
	}

	base := path.Base(r.URL.Path)
	if path.Ext(base) == ".xml" {

		param := map[string] interface{}{
			"host" : "http://" +  r.Host,
		}
		t, _ := template.ParseFiles(base)
		t.Execute(w,param)
		return
	}

	f.fileHandler.ServeHTTP(w, r)
}

func main(){

	fmt.Println("Static demos")
	fmt.Println("Input ip:8080 at Air Launcher on your Apple TV")

	hanlder := PHandler{}
	hanlder.fileHandler = http.FileServer(http.Dir("."))
	http.ListenAndServe(":8080", hanlder)
}

