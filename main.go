package main

import (
	"fmt"

	"github.com/romatr/interestcalculator/service"
)

func main() {
	fmt.Printf("Starting application\n")
	service.StartWebServer("8080")
}
