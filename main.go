package main

import "fmt"

var url string = "http://localhost:9091/api/v1/query?query=scrape_duration_seconds"

func main() {
	callPrometheus()
	fmt.Println("end")
}
