package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	storage := NewStorage()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/index.html", 308)
	})

	http.HandleFunc("/index.html", func(w http.ResponseWriter, req *http.Request) {
		page, _ := ioutil.ReadFile("web/index.html")
		w.Write(page)
	})

	http.HandleFunc("/index.js", func(w http.ResponseWriter, req *http.Request) {
		page, _ := ioutil.ReadFile("web/index.js")
		w.Header().Set("content-type", "text/javascript")
		w.Write(page)
	})

	http.HandleFunc("/data.json", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("content-type", "text/json")
		json.NewEncoder(w).Encode(storage)
	})

	log.Println("Start Server...")
	if err := open("http://localhost:1313"); err != nil {
		log.Println(err)
	}
	log.Fatal(http.ListenAndServe(":1313", nil))
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
