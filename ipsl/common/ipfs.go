package common

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type IPFS struct {
	Shell *ipfsapi.Shell
}

// Connect is used to connect to the given ipfs node
// if no url is supplied, then we connect to the local
// ipfs node
func Connect(url string) *IPFS {
	if url == "" {
		return &IPFS{Shell: ipfsapi.NewLocalShell()}
	}
	return &IPFS{Shell: ipfsapi.NewShell(url)}
}

// DagPutJSON is used to insert json data into the dag
func (i *IPFS) DagPutJSON(data interface{}) (string, error) {
	cid, err := i.Shell.DagPut(data, "json", "cbor")
	if err != nil {
		return "", err
	}
	return cid, nil
}
