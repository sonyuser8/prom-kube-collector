package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"k8s.io/klog"
)

type KubeApiResp struct {
	ApiVersion string `json:"apiVersion"`
}

func execCmd() {
	// Sample code for kubectl
	// TODO: output json and parse by golang struct
	// TODO: regex to parse product and release unit in image path
	selector := "k8s-app=kube-dns"
	out, errout, err := Shellout(fmt.Sprintf("kubectl get po -A -l %s -o json", selector))

	if err != nil {
		klog.Infof("error: %v\n", err)
	}
	fmt.Println(out)
	klog.Info(errout)

	var ret KubeApiResp
	json.Unmarshal([]byte(out), &ret)
	klog.Infof("%+v", ret)

}

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
