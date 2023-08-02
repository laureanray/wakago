package api

type Project struct {
	ID                           int      `json:"id"`
	Name                         string   `json:"name"`
	Repository                   string   `json:"repository"`
	Badge                        string   `json:"badge"`
	Color                        string   `json:"color"`
	Clients                      []string `json:"clients"`
	HasPublicUrl                 bool     `json:"has_public_url"`
	HumanReadableLastHeartbeatAt string   `json:"human_readable_last_heartbeat_at"`
	LastHeartbeatAt              string   `json:"last_heartbeat_at"`
	Url                          string   `json:"url"`
	UrlEncodedName               string   `json:"urlencoded_name"`
	CreatedAt                    string   `json:"created_at"`
}

type Projects struct {
	CachedAt string  `json:"cached_at"`
	Data     Project `json:"data"`
}
