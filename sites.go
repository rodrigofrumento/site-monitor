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

const numMonit = 3
const seg = 5

func main() {
	showIntro()

	for {

		showMenu()
		command := readCommand()
		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("This command is not valid!")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "YOUR NAME"
	version := 1.1
	fmt.Println("Hello", name)
	fmt.Println("This program is on version: ", version)
}

func showMenu() {
	fmt.Println("Select one option")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := readSitesFromDoc()

	for i := 0; i < numMonit; i++ {
		fmt.Println("Starting next monitoring")
		for i, site := range sites {
			fmt.Println("monitoring site", i, ":", site)
			siteTesting(site)
		}
		time.Sleep(seg * time.Second)
		fmt.Println("")
		fmt.Println("")
	}

}

func siteTesting(site string) {
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfuly load!")
	} else {
		fmt.Println("Site:", site, "something goes wrong:", res.StatusCode)
	}
}

func readSitesFromDoc() []string {
	var sites []string
	websites, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(websites)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}
	}
	websites.Close()
	return sites
}

func logWriting(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu erro", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu erro", err)
	}
	fmt.Println(string(file))
}
