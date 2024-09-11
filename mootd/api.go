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

func schedule() {
	for true {
		now, todayRenewalTime := getTimeData()
		// Calculate time until next renewal
		timeUntilRenewal := todayRenewalTime.Sub(now)
		if timeUntilRenewal < 0 {
			timeUntilRenewal += 24 * time.Hour
		}
		// Wait until renewal, then renew
		time.Sleep(timeUntilRenewal)
		gen()
	}
}

func check() {
	// Create file if it doesn't exist
	if _, err := os.Stat("/srv/mootd"); os.IsNotExist(err) {
		gen()
	} else {
		now, todayRenewalTime := getTimeData()
		// Get file modified datetime
		fileInfo, _ := os.Stat("/srv/mootd")
		fileAge := time.Now().Sub(fileInfo.ModTime())
		// Calculate time since last renewal
		timeSinceRenewal := now.Sub(todayRenewalTime)
		if timeSinceRenewal < 0 {
			timeSinceRenewal += 24 * time.Hour
		}
		if fileAge > timeSinceRenewal {
			gen()
		}
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
	go schedule()
	serve()
}
