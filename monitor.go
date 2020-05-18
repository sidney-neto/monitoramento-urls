package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 1
const delay = 2

func main() {

	for {
		menu()

		switch command() {
		case 1:
			monitoring()
		case 2:
			readLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command Invalid")
			os.Exit(-1)
		}
	}
}

func menu() {
	fmt.Println("\n1. Monitoring Url")
	fmt.Println("2. Viewing Logs")
	fmt.Println("0. Exit")
}

func command() int {
	fmt.Print("Options: ")
	var command int
	fmt.Scanln(&command)
	return command
}

func monitoring() {
	// urlsList := []string{
	// 	"https://alura.com.br",
	// 	"https://google.com.br",
	// }

	urlsList := readFile()

	for i := 0; i < monitoringTimes; i++ {
		for _, url := range urlsList {
			fmt.Println("\nMonitoring", url)
			response(url)
		}
		time.Sleep(delay * time.Second)
	}
}

func response(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Status Code:", response.StatusCode)
		fmt.Println("Url ok")
		createLogs(url, int64(response.StatusCode))
	} else {
		fmt.Println("Status Code:", response.StatusCode)
		fmt.Println("Url not ok")
		createLogs(url, int64(response.StatusCode))
	}
}

func readFile() []string {
	var urls []string

	// file, err := ioutil.ReadFile("urls.txt")
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		url, err := reader.ReadString('\n')
		url = strings.TrimSpace(url)
		urls = append(urls, url)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return urls
}

func createLogs(url string, status int64) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("error:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ; ") + "Url: " + url + " ; Status Code: " + strconv.FormatInt(status, 10) + "\n")

	file.Close()
}

func readLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\nShowing Logs...")
	fmt.Println(string(file))
}
