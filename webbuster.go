package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joanbono/webbuster/static"
)

var data string

func init() {
	flag.StringVar(&data, "data", "", "JSON file to use")
	flag.Parse()
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	log.Println(" -  Method:", r.Method, " - /")
	var myJSON []static.BustData
	plan, err := ioutil.ReadFile(data)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(plan, &myJSON); err != nil {
		fmt.Println(err)
	}
	if r.Method == "GET" {
		t, _ := template.New("code").Parse(static.HtmlCode)
		//t, _ := template.ParseFiles("static/index.html")
		//u := static.BustData{Url: &myJSON.Url, Method: myJSON.Method}
		//t.Execute(w, u)
		//fmt.Printf("%v\n", myJSON)
		t.Execute(w, myJSON)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	if data == "" {
		fmt.Println("WebBuster USAGE:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	http.HandleFunc("/", mainpage)

	fmt.Println("[+] WeBBuster -  Listening on http://localhost:9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(" -  ListenAndServe: ", err)
	}

}
