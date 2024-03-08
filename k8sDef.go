package main

type KubeApiResp struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Labels    map[string]interface{}
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Image string `json:"image"`
				// LivenessProbe struct {
				// 	FailureThreshold int `json:"failureThreshold"`
				// 	HTTPGet          struct {
				// 		Path   string `json:"path"`
				// 		Port   int    `json:"port"`
				// 		Scheme string `json:"scheme"`
				// 	} `json:"httpGet"`
				// 	InitialDelaySeconds int `json:"initialDelaySeconds"`
				// 	PeriodSeconds       int `json:"periodSeconds"`
				// 	SuccessThreshold    int `json:"successThreshold"`
				// 	TimeoutSeconds      int `json:"timeoutSeconds"`
				// } `json:"livenessProbe"`
				// ReadinessProbe struct {
				// 	FailureThreshold int `json:"failureThreshold"`
				// 	HTTPGet          struct {
				// 		Path   string `json:"path"`
				// 		Port   int    `json:"port"`
				// 		Scheme string `json:"scheme"`
				// 	} `json:"httpGet"`
				// 	PeriodSeconds    int `json:"periodSeconds"`
				// 	SuccessThreshold int `json:"successThreshold"`
				// 	TimeoutSeconds   int `json:"timeoutSeconds"`
				// } `json:"readinessProbe"`
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
				// VolumeMounts []struct {
				// 	MountPath string `json:"mountPath"`
				// 	Name      string `json:"name"`
				// 	ReadOnly  bool   `json:"readOnly"`
				// } `json:"volumeMounts"`
			} `json:"containers"`
			Volumes []struct {
				// ConfigMap struct {
				// 	DefaultMode int `json:"defaultMode"`
				// 	Items       []struct {
				// 		Key  string `json:"key"`
				// 		Path string `json:"path"`
				// 	} `json:"items"`
				// 	Name string `json:"name"`
				// } `json:"configMap,omitempty"`
				Name string `json:"name"`
			} `json:"volumes"`
		} `json:"spec"`
	} `json:"items"`
}
