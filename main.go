package main

import (
	"flag"
	"log"
)

var (
	endpoint string
	metadata string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "fleet endpoint")
	flag.StringVar(&metadata, "metadata", "k8s=kubelet", "comma-delimited key/value pairs")
}

func main() {
	flag.Parse()
	m, err := parseMetadata(metadata)
	if err != nil {
		log.Println(err)
	}
	machines, err := getMachines(endpoint, m)
	if err != nil {
		log.Println(err)
	}
	for _, machine := range machines {
		print(machine)
	}
	// add nodes to api service
}
