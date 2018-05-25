package common

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
)

// IPFS is a helper struct
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

// DagGet is not yet implemented
func (i *IPFS) DagGet() {
	// TODO: implement
}

// GetLocalPeerID is used to return the id
// of the local peer
func (i *IPFS) GetLocalPeerID() (string, error) {
	peer, err := i.Shell.ID()
	if err != nil {
		return "", err
	}
	return peer.ID, nil
}
