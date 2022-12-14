package api

type ChartEntry struct {
	ActualSeconds     float64 `json:"actual_seconds"`
	ActualSecondsText string  `json:"actual_seconds_text"`
	GoalSeconds       float64 `json:"goal_seconds"`
	GoalSecondsText   string  `json:"goal_seconds_text"`
	Range             Range   `json:"range"`
	RangeStatus       string  `json:"range_status"`
	RangeStatusReason string  `json:"ranges_status_reason"`
}

type Subscriber struct {
	Email          string `json:"email"`
	EmailFrequency string `json:"email_frequency"`
	FullName       string `json:"full_name"`
	UserId         string `json:"user_id"`
	Username       string `json:"username"`
}

type GoalData struct {
	ID               string       `json:"id"`
	AverageStatus    string       `json:"average_status"`
	ChartData        []ChartEntry `json:"chart_data"`
	CumulativeStatus string       `json:"cumulative_status"`
	Delta            string       `json:"delta"`
	IgnoreDays       []string     `json:"ignore_days"`
	ImproveByPercent float64      `json:"improve_by_percent"` // for some reason always null
	IsEnabled        bool         `json:"is_enabled"`
	Languages        []string     `json:"languages"`
	Projects         []string     `json:"projects"`
	RangeText        string       `json:"range_text"`
	Seconds          int64        `json:"seconds"`
	Status           string       `json:"status"`
	Subscribers      []Subscriber `json:"subscribers"`
	Title            string       `json:"title"`
	Type             string       `json:"type"`
}

type Goal struct {
	CachedAt string   `json:"cached_at"`
	Data     GoalData `json:"data"`
}

type Goals struct {
	Data       []GoalData `json:"data"`
	Total      int64      `json:"total"`
	TotalPages int64      `json:"total_pages"`
}
