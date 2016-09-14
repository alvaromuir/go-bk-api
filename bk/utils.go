package bk

import (
	"fmt"

	"log"
	"strconv"
	"time"
)

// ParseMilisecondTime returns a human readable date from millisecond time string
func ParseMilisecondTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	t := time.Unix(0, msInt*int64(time.Millisecond))
	return t, nil
}

// PrintBanner prints a banner to stdout for testing purposes
func PrintBanner(message string) {
	fmt.Println("")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println(message)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("")
}

// CheckError is a generic error logger
func CheckError(err error, msg string) {
	if err != nil {
		log.Printf(msg, err)
	}
	return
}
