package cmd

import "github.com/postables/ipsl-go/ipsl/common"

// ConfigPath is the config file
const ConfigPath = "/ipsl/config.json"

// TestConfigPath is the config path to use when using test mode
const TestConfigPath = "/ipsl/test_config.json"

// TestMode indicates whether or not we are running in test mode
const TestMode = false

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
