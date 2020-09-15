package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

type timeJSON struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
}

func main() {
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
		log.Println("Retrieving Data")
		now := time.Now()
		date := timeJSON{
			Year:   fmt.Sprintf("%04d", now.Year()),
			Month:  fmt.Sprintf("%02d", now.Month()),
			Day:    fmt.Sprintf("%02d", now.Day()),
			Hour:   fmt.Sprintf("%02d", now.Hour()),
			Minute: fmt.Sprintf("%02d", now.Minute()),
			Second: fmt.Sprintf("%02d", now.Second())}
		w.Header().Set("content-type", "text/json")
		json.NewEncoder(w).Encode(date)
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
