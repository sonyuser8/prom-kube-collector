package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"k8s.io/klog"
)

func execKubectl() {
	// TODO: 如何正規化 cpu / memory 之數值並且相加
	selector := "k8s-app=kube-dns"
	out, errout, err := Shellout(fmt.Sprintf("kubectl get po -A -l %s -o json", selector))

	if err != nil {
		klog.Infof("error: %v\n", err)
	}
	// fmt.Println(out)
	klog.Info(errout)

	var ret KubeApiResp
	json.Unmarshal([]byte(out), &ret)
	// klog.Infof("%+v", ret)
	klog.Infof("%s", ret.Items[0].Spec.Containers[0].Image)

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
