package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"k8s.io/klog"
)

type PromResp struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]interface{}
			Value  []any `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func callPrometheus() {
	dat := getUrl(url)
	// dat, _ := os.ReadFile("data.json")

	if dat == nil {
		klog.Warningf("No data from url %s", url)
		return
	}
	klog.Info("getUrl pass")
	var ret PromResp
	json.Unmarshal(dat, &ret)
	for _, value := range ret.Data.Result {
		fmt.Println(value.Metric["instance"]) // hardcode instance for now
	}

}

func getUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		klog.Warning(err)
		return nil
	}
	defer resp.Body.Close()
	ret, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		klog.Warning(err)
		return nil
	}
	return ret
}
