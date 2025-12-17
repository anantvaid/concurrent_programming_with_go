package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type PodManifest struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Containers []struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"containers"`
		Replicas int `json:"replicas"`
	} `json:"spec"`
}

func main() {
	namePtr := flag.String("name", "myapp", "Name of the Pod")
	imagePtr := flag.String("image", "nginx", "Container Image")
	replicasPtr := flag.Int("replicas", 1, "Number of Replicas")

	flag.Parse()

	if *replicasPtr <= 0 {
		fmt.Fprintf(os.Stderr, "Replicas must be greater than 0, you provided %d", *replicasPtr)
		flag.Usage()
		os.Exit(1)
	}

	manifest := PodManifest{
		APIVersion: "v1",
		Kind:       "Pod",
	}
	manifest.Metadata.Name = *namePtr
	manifest.Spec.Replicas = *replicasPtr

	container := struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	}{
		Name:  *namePtr,
		Image: *imagePtr,
	}

	manifest.Spec.Containers = append(manifest.Spec.Containers, container)

	output, err := json.MarshalIndent(manifest, "", " ")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
