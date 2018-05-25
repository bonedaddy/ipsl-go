package common

import (
	"errors"
	"fmt"
	"strings"
)

var protocols = [3]string{"ipfs", "https", "ftp"}

// CleanAddress is used to remove the protocol string from the ipfs that is passed in
func CleanAddress(protocol string, address string) (string, error) {
	if address == "" {
		return "", errors.New("invalid address")
	}

	switch protocol {
	case "ipfs":
		// replace /ipfs/ with ''
		address = strings.Replace(fmt.Sprint(address), "/ipfs/", "''", 1)
	case "https":
		// replace https:// with ''
		address = strings.Replace(fmt.Sprint(address), "https://", "''", 1)
	case "ftp":
		// replace ftp:// with ''
		address = strings.Replace(fmt.Sprint(address), "ftp://", "''", 1)
	default:
		return "", errors.New("invalid protocol")
	}

	return address, nil
}
