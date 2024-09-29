package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const delay = 5

func insertUrl(url *string) {
	fmt.Print("Insert site you want to check: ")
	fmt.Scanln(url)
}

func menu() {

	fmt.Println("\n1- Start monitoring")
	fmt.Println("2- automatic monitoring")
	fmt.Println("3- View Logs")
	fmt.Println("4- Change url")
	fmt.Println("0- Exit the Program")
}

func startMonitor(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", url, "was loaded with success")
		registerLog(url, true)
	} else {
		fmt.Println("Site:", url, "is with a problem. Status Code:", resp.StatusCode)
		registerLog(url, false)
	}
}

func autoMonitor(monitor *int, url string) {
	fmt.Print("Enter how many monitorings you want to do: ")
	fmt.Scanln(monitor)
	for i := 0; i < *monitor; i++ {
		startMonitor(url)
		time.Sleep(delay + time.Second)
		fmt.Println("")
	}
}

func registerLog(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + url + "- online:" + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(file))

}

func main() {
	var comando int
	var url string
	var monitor int

	insertUrl(&url)
	for {

		menu()
		fmt.Print("Enter the command you want: ")
		fmt.Scanln(&comando)

		switch comando {
		case 1:
			startMonitor(url)
		case 2:
			autoMonitor(&monitor, url)
		case 3:
			printLogs()
		case 4:
			insertUrl(&url)
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("unidentified command")
			os.Exit(-1)
		}
	}
}
