package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Minion struct {
	Kind   string `json:"kind,omitempty"`
	ID     string `json:"id,omitempty"`
	HostIP string `json:"hostIP,omitempty"`
}

func register(endpoint, addr string) error {
	m := &Minion{
		Kind:   "Minion",
		ID:     addr,
		HostIP: addr,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/v1beta1/minions", endpoint)
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 202 || res.StatusCode == 200 {
		log.Printf("registered machine: %s\n", addr)
		return nil
	}
	return errors.New("error registering: " + addr)
}
