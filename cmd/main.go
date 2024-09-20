package main

import "fmt"

func main() {
	host := GetInstance().Host
	port := GetInstance().Port
	fmt.Printf(host + ":" + string(port))
}
