package common

import (
	"fmt"
	"log"
	"strings"
)

var protocols = [3]string{"ipfs", "https", "ftp"}

func CleanAddress(protocol string, address string) {
	if address == "" {
		return
	}

	switch protocol {
	case "ipfs":
		address = strings.Replace(fmt.Sprint(address), "/ipfs/", "''", 1)
	default:
		log.Fatal("not a valid protocol")
	}
}
