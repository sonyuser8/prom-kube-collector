package main

import (
	"fmt"
)

const url = "http://localhost:9091/api/v1/query?query=scrape_duration_seconds"
const ShellToUse = "bash"

func main() {
	// Sample code for kubectl
	// TODO: output json and parse by golang struct
	// TODO: regex to parse product and release unit in image path
	execCmd()

	// Sample code for prometheus
	callPrometheus()
	fmt.Println("end")
}
