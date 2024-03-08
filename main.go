package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AutoGenerated struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]interface{}
			Value  []any `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

var url string = "http://localhost:9091/api/v1/query?query=scrape_duration_seconds"

func main() {
	dat := getUrl(url)
	// dat, _ := os.ReadFile("data.json")

	if dat == nil {
		log.Fatalf("No data from url %s", url)
	}
	log.Default().Println("getUrl pass")
	var ret AutoGenerated
	json.Unmarshal(dat, &ret)
	for _, value := range ret.Data.Result {
		fmt.Println(value.Metric["instance"]) // hardcode instance for now
	}
	log.Println("end")
}

func getUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	ret, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		log.Println(err)
		return nil
	}
	return ret
}
