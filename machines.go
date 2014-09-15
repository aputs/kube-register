package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/fleet/client"
)

func getMachines(endpoint string, metadata map[string]string) ([]string, error) {
	machineList := make([]string, 0)
	c := &http.Client{}
	fleetClient, err := client.NewHTTPClient(c, endpoint)
	if err != nil {
		return nil, err
	}
	machines, err := fleetClient.Machines()
	if err != nil {
		return nil, err
	}
	for _, m := range machines {
		if hasMetadata(m, metadata) && isHealthy(m.PublicIP) {
			machineList = append(machineList, m.PublicIP)
		}
	}
	return machineList, nil
}

func isHealthy(addr string) bool {
	url := fmt.Sprintf("http://%s:%d/healthz", addr, 10250)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("error health checking %s: %s", addr, err)
		return false
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
		return true
	}
	log.Printf("unhealthy machine: %s will not be registered", addr)
	return false
}
