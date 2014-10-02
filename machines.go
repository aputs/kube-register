package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/coreos/fleet/client"
)

func getMachines(endpoint string, metadata map[string][]string) ([]string, error) {
	dialFunc := net.Dial
	machineList := make([]string, 0)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "unix" {
		endpoint = "http://domain-sock/"
		dialFunc = func(network, addr string) (net.Conn, error) {
			return net.Dial("unix", u.Path)
		}
	}
	c := &http.Client{
		Transport: &http.Transport{
			Dial: dialFunc,
		},
	}
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
