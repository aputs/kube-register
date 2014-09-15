package main

import (
	"errors"
	"strings"

	"github.com/coreos/fleet/machine"
)

func hasMetadata(ms machine.MachineState, metadata map[string]string) bool {
	for k, v := range metadata {
		if ms.Metadata[k] != v {
			return false
		}
	}
	return true
}

func parseMetadata(rawMetadata string) (map[string]string, error) {
	metadataList := strings.Split(metadata, ",")
	metadata := make(map[string]string)
	for _, kv := range metadataList {
		i := strings.Index(kv, "=")
		if i > 0 {
			metadata[kv[:i]] = kv[i+1:]
		} else {
			return nil, errors.New("invalid key/value pair " + kv)
		}
	}
	return metadata, nil
}
