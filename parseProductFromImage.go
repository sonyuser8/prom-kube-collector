package main

import (
	"fmt"
	"regexp"
)

func getProductProfileFromImageName() {
	image := "registry.k8s.io/coredns/coredns:v1.11.1"
	re, _ := regexp.Compile(`.*\/(.*):`)
	// fmt.Println(re, err)
	// matched := re.MatchString(image)
	// re.FindSubmatch(image)
	fmt.Println(re.MatchString(image))
	fmt.Println(re.FindStringSubmatch(image)[1])
}
