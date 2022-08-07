package api

type Range struct {
	Date     string `json:"date"`
	End      string `json:"end"`
	Start    string `json:"start"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}
