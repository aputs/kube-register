package main

import (
	"flag"
	"log"
)

var (
	apiEndpoint   string
	fleetEndpoint string
	metadata      string
)

func init() {
	log.SetFlags(0)
	flag.StringVar(&apiEndpoint, "api-endpoint", "", "kubernetes API endpoint")
	flag.StringVar(&fleetEndpoint, "fleet-endpoint", "", "fleet endpoint")
	flag.StringVar(&metadata, "metadata", "k8s=kubelet", "comma-delimited key/value pairs")
}

func main() {
	flag.Parse()
	m, err := parseMetadata(metadata)
	if err != nil {
		log.Println(err)
	}
	machines, err := getMachines(fleetEndpoint, m)
	if err != nil {
		log.Println(err)
	}
	for _, machine := range machines {
		if err := register(apiEndpoint, machine); err != nil {
			log.Println(err)
		}
	}
}
