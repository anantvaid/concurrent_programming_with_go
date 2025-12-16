package main

import (
	"encoding/json"
	"fmt"
)

type PodConfig struct {
	Name       string            `json:"pod_name"`
	Replicas   int               `json:"replicas,omitempty"`
	Labels     map[string]string `json:"labels"`
	InternalID string            `json:"-"`
}

func main() {
	pod := PodConfig{
		Name:     "nginx-frontend",
		Replicas: 3,
		Labels: map[string]string{
			"app": "frontend",
		},
		InternalID: "xyz123",
	}

	jsonStr, _ := json.MarshalIndent(pod, "", " ")

	fmt.Println("---------Full Pod Config---------\n", string(jsonStr))

	// Trying out empty config
	emptyPod := PodConfig{
		Name: "busybox",
	}

	jsonStr2, _ := json.MarshalIndent(emptyPod, "", " ")

	fmt.Println("---------Empty Pod Config---------\n", string(jsonStr2))

	// Trying to parse JSON to Go Struct
	jsonData := `
        {
            "pod_name": "nginx", 
            "labels": {"env": "prod"}, 
            "image": "nginx:1.14.2" 
        }
    `

	var config PodConfig

	err := json.Unmarshal([]byte(jsonData), &config)

	if err != nil {
		fmt.Println("Issue in parsing")
	}

	fmt.Println("------JSON Pod Parsing--------")
	fmt.Printf("Pod Name: %s\nPod Env Label: %s\n", config.Name, config.Labels["env"])
}
