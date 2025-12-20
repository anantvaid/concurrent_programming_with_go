package main

import (
	"encoding/json"
	"fmt"
)

type ManifestGenerator interface {
	Generator() ([]byte, error)
}

type Pod struct {
	Name  string
	Image string
}

func (p Pod) Generator() ([]byte, error) {
	manifest := struct {
		Kind       string `json:"kind"`
		ApiVersion string `json:"apiVersion"`
		Metadata   struct {
			Name string `json:"name"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Name  string `json:"name"`
				Image string `json:"image"`
			} `json:"containers"`
		} `json:"spec"`
	}{
		Kind:       "Pod",
		ApiVersion: "v1",
	}
	manifest.Metadata.Name = p.Name
	manifest.Spec.Containers = []struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	}{{Name: p.Name, Image: p.Image}}

	return json.MarshalIndent(manifest, "", " ")
}

type Service struct {
	Name string
	Port int
}

func (s Service) Generator() ([]byte, error) {
	manifest := struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Name string `json:"name"`
		} `json:"metadata"`
		Spec struct {
			Ports []struct {
				Port     int    `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
		} `json:"spec"`
	}{
		APIVersion: "v1",
		Kind:       "Service",
	}
	manifest.Metadata.Name = s.Name
	manifest.Spec.Ports = []struct {
		Port     int    `json:"port"`
		Protocol string `json:"protocol"`
	}{{Port: s.Port, Protocol: "TCP"}}

	return json.MarshalIndent(manifest, "", " ")
}

func main() {
	resources := []ManifestGenerator{
		Pod{Name: "nginx-pod", Image: "nginx:latest"},
		Service{Name: "nginx-svc", Port: 80},
		Pod{Name: "redis-pod", Image: "redis:alpine"},
	}

	for _, res := range resources {
		output, err := res.Generator()
		if err != nil {
			fmt.Printf("Something is not right! %v", err)
		}

		fmt.Println(string(output))
		fmt.Println("------------")
	}
}
