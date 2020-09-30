package api

type Configs struct {
	OrgConfigs map[string]OrgConfig
}

type OrgConfig struct {
	Name string `json:"name"`
	Database string `json:"database"`
}
