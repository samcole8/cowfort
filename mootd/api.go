package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func gen() {
	exec.Command("/bin/sh", "/usr/local/bin/renew.sh").Run()
	fmt.Println("MOTD generated")
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
		fmt.Println("Renewal at " + os.Getenv("RENEWAL_TIME"), "in " + timeUntilRenewal.String())
		time.Sleep(timeUntilRenewal)
		gen()
	}
}

func check() {
	// Create file if it doesn't exist
	if _, err := os.Stat("/srv/mootd"); os.IsNotExist(err) {
		fmt.Println("MOTD not found")
		gen()
	} else {
		fmt.Println("MOTD found, testing validity")
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
			fmt.Println("MOTD has expired")
			gen()
		} else {
			fmt.Println("MOTD in-date, " + fileAge.String() + " old")
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
	fmt.Println("Starting checks")
	check()
	fmt.Println("Starting scheduler")
	go schedule()
	fmt.Println("Starting server")
	serve()
}
