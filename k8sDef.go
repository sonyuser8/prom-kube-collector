package main

// Should figure out how community handle this part for pod / deployment / statefulset / daemonset / cronjob
type KubeApiResp struct {
	Items []struct {
		Kind     string `json:"kind"`
		Metadata struct {
			Labels    map[string]interface{}
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Image     string `json:"image"`
				Resources struct {
					Limits struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"limits"`
					Requests struct {
						CPU    string `json:"cpu"`
						Memory string `json:"memory"`
					} `json:"requests"`
				} `json:"resources"`
			} `json:"containers"`
			Volumes []struct {
				Name string `json:"name"`
			} `json:"volumes"`
		} `json:"spec"`
	} `json:"items"`
}
