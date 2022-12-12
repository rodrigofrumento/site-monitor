package main

import (
	"fmt"
	"net/http"
	"os"
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
	sites := []string{
		"site1",
		"site2",
		"site3"}

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
