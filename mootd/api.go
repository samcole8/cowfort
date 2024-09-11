package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func check() {
	if _, err := os.Stat("/srv/mootd"); os.IsNotExist(err) {
		exec.Command("/bin/sh", "/usr/local/bin/renew.sh").Run()
	}
}

func serve() {
	var get = func(w http.ResponseWriter, _ *http.Request) {
		content, _ := os.ReadFile("/srv/mootd")
		io.WriteString(w, string(content))
	}

	http.HandleFunc("/", get)
	http.ListenAndServe(":80", nil)
}

func main() {
	check()
	serve()
}
