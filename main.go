package main

import (
	"fmt"
)

const url = "http://localhost:9090/api/v1/query?query=scrape_duration_seconds"
const ShellToUse = "bash"

func main() {
	getProductProfileFromImageName()

	// Sample code for kubectl
	execKubectl()

	// Sample code for prometheus
	callPrometheus()
	fmt.Println("end")
}
