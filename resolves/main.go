package main

import (
    "net"
    "bufio"
    "fmt"
    "os"

)

//var urls []string

func main(){
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var urls []string
		urls = append(urls, scanner.Text())
		for _, url := range urls {
        	addr,err := net.LookupIP(url)
           if err != nil {
              fmt.Println("Unknown host")
           } else {
              fmt.Println("Hostname: " + url + " IP :", addr)
           }
		}
}
fmt.Println("Done! (╯°□°)╯︵ ┻━┻")
}
