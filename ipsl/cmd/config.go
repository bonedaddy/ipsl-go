package cmd

import "github.com/postables/ipls-go/ipsl/common"

// Domains holds domain related info
type Domains struct {
	Patterns []string `json:"patterns"`
	Seen     []string `json:"seen"`
}

// Init is used to generate our config
func Init() map[string]interface{} {

	config := make(map[string]interface{})
	maps := make(map[string]interface{})
	for _, v := range common.Protocols {
		switch v {
		case "ipfs":
			maps["ipfs"] = "place"
		}
	}
	config["domains"] = Domains{
		Patterns: []string{"test"},
		Seen:     []string{"testtest"},
	}
	config["map"] = maps
	return config
}
