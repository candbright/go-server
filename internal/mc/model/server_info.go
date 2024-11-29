package model

type ServerInfo struct {
	Version          string      `json:"version"`
	Exist            bool        `json:"exist"`
	Active           bool        `json:"active"`
	ServerProperties interface{} `json:"server_properties"`
	AllowList        interface{} `json:"allow_list"`
}
