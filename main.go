package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	var setting string
	var num int

	if len(os.Args) < 2 {
		fmt.Println("Please enter the setting you want (e.g., 'outlook', 'gmail', etc.):")
		fmt.Scan(&setting)
	} else {
		setting = os.Args[1]
	}

	if len(os.Args) < 3 {
		fmt.Println("Please enter the number of emails you want:")
		fmt.Scan(&num)
	} else {
		num, _ = strconv.Atoi(os.Args[2])
	}

	var domain string
	switch setting {
	case "hotmail":
		domain = "hotmail.com"
	case "outlook":
		domain = "outlook.com"
	case "gmail":
		domain = "gmail.com"
	case "yahoo":
		domain = "yahoo.com"
	case "icloud":
		domain = "icloud.com"
	case "aol":
		domain = "aol.com"
	case "live":
		domain = "live.com"
	case "comcast":
		domain = "comcast.net"
	case "att":
		domain = "att.net"
	case "verizon":
		domain = "verizon.net"
	case "bell":
		domain = "bell.net"
	case "sprint":
		domain = "sprint.com"
	case "tmobile":
		domain = "t-mobile.com"
	case "rogers":
		domain = "rogers.com"
	case "disney":
		domain = "disney.com"
	case "charter":
		domain = "charter.net"
	case "bt":
		domain = "btinternet.com"
	case "virgin":
		domain = "virgin.net"
		// Add more cases for other settings if needed
	default:
		fmt.Println("Invalid setting.")
		return
	}

	var str string
	for i := 0; i < num; i++ {
		str += RandStringBytes(25) + "@" + domain + "\n"
	}

	file, errs := os.Create(domain + "Emails.txt")
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file.Close()

	_, errs = file.WriteString(str)
	if errs != nil {
		fmt.Println("Failed to write to file"+setting+"Emails.txt", errs)
		return
	}
	fmt.Println(num, "emails saved to "+setting+"Emails.txt")
}
