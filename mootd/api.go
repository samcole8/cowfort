package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func gen() {
	exec.Command("/bin/sh", "/usr/local/bin/renew.sh").Run()
}

func getTimeData() (time.Time, time.Time) {
	// Get current datetime
	now := time.Now()
	// Get renewal datetime
	renewalTimeStr := os.Getenv("RENEWAL_TIME")
	renewalTime, _ := time.Parse("15:04:05", renewalTimeStr)
	todayRenewalTime := time.Date(
		now.Year(), now.Month(), now.Day(),
		renewalTime.Hour(), renewalTime.Minute(), renewalTime.Second(), 0, now.Location(),
	)
	return now, todayRenewalTime
}

func check() {
	// Create file if it doesn't exist
	if _, err := os.Stat("/srv/mootd"); os.IsNotExist(err) {
		gen()
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
