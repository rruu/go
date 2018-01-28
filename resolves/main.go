package main

import (
    "net"
    "bufio"
    "fmt"
    "os"

)

var urls []string

func doGet(url string) {
	addr,err := net.LookupIP(url)
   if err != nil {
      fmt.Println("Unknown host")
   } else {
      fmt.Println("Hostname: " + url + " IP :", addr)
   }
}


func loadUrls() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
}


func main() {
	loadUrls()
	for _, url := range urls {
		doGet(url)
	}
	fmt.Println("Done! (╯°□°)╯︵ ┻━┻")
}


