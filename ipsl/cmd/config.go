package cmd

// Domains holds domain related info
type Domains struct {
	Patterns []string `json:"patterns"`
	Seen     []string `json:"seen"`
}

// Init is used to generate our config
func Init() map[string]interface{} {
	config := make(map[string]interface{})
	config["domains"] = Domains{
		Patterns: []string{"test"},
		Seen:     []string{"testtest"},
	}
	return config
}
