package main

import (
	"fmt"
)

const url = "http://localhost:9090/api/v1/query?query=scrape_duration_seconds"
const ShellToUse = "bash"

func main() {
	// TODO: learn how to use klog v to control log level
	getProductProfileFromImageName()
	// TODO: How to export csv

	// Sample code for kubectl
	// TODO: 如何正規化 cpu / memory 之數值並且相加
	execKubectl()

	// Sample code for prometheus
	callPrometheus()
	fmt.Println("end")
}
