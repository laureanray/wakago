package api

type GrandTotal struct {
	Digital      string  `json:"digital"`
	Hours        int64   `json:"hours"`
	Minutes      int64   `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Categories struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Text         string  `json:"text"`
	Percent      float64 `json:"percent"`
	Hours        int64   `json:"hours"`
	Minutes      int64   `json:"minutes"`
}

// TODO: Add projects, editors, os, deps, machines

type Languages struct {
	Name         string  `json:"name"`
	Digital      string  `json:"digital"`
	Hours        int64   `json:"hours"`
	Minutes      int64   `json:"minutes"`
	Seconds      int64   `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
}

type StatusBarData struct {
	GrandTotal GrandTotal   `json:"grand_total"`
	Categories []Categories `json:"categories"`
	Languages  []Languages  `json:"languages"`
}

type StatusBar struct {
	CachedAt string        `json:"cached_at"`
	Data     StatusBarData `json:"data"`
}
