
/*

local media file server sample

*/


package main

import (
	"net/http"
	"fmt"
	"strings"
	"net/url"
	"path"
	"io/ioutil"
	"text/template"
	"flag"
	"strconv"
	"net"
	"os"
	"github.com/k0kubun/pp"
)


func main(){


	fmt.Println("Media file server sample for Air Launcher ver 1.0.0")

	var port = 8080
	flag.IntVar(&port, "p", 8080, "serve port")
	flag.Parse()

	mediaPath := flag.Arg(0)
	if mediaPath == "" {
		fmt.Println("Usage: fileserve [-p port(default:8080] [media folder]")
		return
	}

	fmt.Println("Serve local path  " + mediaPath)
	fmt.Println("Serve port : " + strconv.Itoa(port))

	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println("Serve  " + ipnet.IP.String() + ":"+ strconv.Itoa(port))
				}
			}
		}
	}


	if _,err := os.Stat("list.xml") ; err != nil {
		fmt.Println("cannot find list.xml")
		return
	}

	fmt.Println("input address(IP:port) at Air Launcher on your Apple TV.")


	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("/ redicrect to /xml/")
		http.Redirect(w,r, "/xml/",302)
	})

	http.HandleFunc("/xml/",func(w http.ResponseWriter, r *http.Request){
		rpath := strings.Replace(r.URL.Path, "/xml", "", 1)
		fmt.Println("XML " + r.URL.Path)

		pp.Println(r)


		base := "http://" + r.Host

		list := [] map[string]string {}

		ar := strings.Split(rpath,"/")
		ar[0] = mediaPath
		dpath := strings.Join(ar,"/")

		title := path.Base(rpath)

		files, _ := ioutil.ReadDir(dpath)

		for _,f := range files {

			name := f.Name()
			if name[0] == '.' { continue }

			repath := url.QueryEscape(rpath + "/" + name)
			repath = strings.Replace(repath, "+", "%20", -1)

			ename := url.QueryEscape(name)
			ename = strings.Replace(ename, "+", "%20", -1)


			name = strings.Replace(name,"&","and",-1)

			mtype := ""
			mpath := "/media"

			if f.IsDir() {
				mtype = "xml"
				mpath = "/xml/"
			}else {
				switch path.Ext(name) {
				case ".mp4", ".mov", ".m3u8", ".m4v", ".mpg", ".mpeg" :
					mtype = "video"

				case ".mp3", ".m4a" :
					mtype = "audio"
				}
			}

			if mtype != "" {
				p := map[string]string{
					"title": name,
					"url" : base + mpath + repath,
					"type" : mtype,
					"desc" : "",
					"thumb" : "",
					"sthumb":"",
				}

				list = append(list, p)
			}
		}

		param := map[string] interface{}{
			"title" : title,
			"list" : list,
		}
		t, _ := template.ParseFiles("list.xml")
		t.Execute(w,param)
	})

	http.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {

		rpath := strings.Replace(r.URL.Path, "/media", "", 1)

		ar := strings.Split(rpath,"/")
		ar[0] = mediaPath
		dpath := strings.Join(ar,"/")

		fmt.Println("media " + path.Base(dpath) + "  " + r.Header["Range"][0])

		http.ServeFile(w, r, dpath)
	})

	http.HandleFunc("/meta.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("meta")
		http.ServeFile(w, r, mediaPath + "meta.json")
	})


	http.ListenAndServe(":" +  strconv.Itoa(port), nil)
}

