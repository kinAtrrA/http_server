package main

import (
	"fmt"
	"strconv"

	config "github.com/kinAtrrA/http_server/pkg/config"
)

func main() {
	host := config.GetInstance().Host
	port := config.GetInstance().Port
	fmt.Printf(host + ":" + strconv.FormatInt(int64(port), 10))
	fmt.Println()
}
