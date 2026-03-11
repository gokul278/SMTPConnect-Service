package dashboardvalidate

type DailyStats struct {
	Date  string `json:"date"`
	Sent  int64  `json:"sent"`
	Fixed int64  `json:"fixed"` // failed but recorded for chart
}

type DashboardStats struct {
	TotalEmails    int64        `json:"totalEmails"`
	SentEmails     int64        `json:"sentEmails"`
	FailedEmails   int64        `json:"failedEmails"`
	SuccessRate    float64      `json:"successRate"`
	ActiveConfigs  int64        `json:"activeConfigs"`
	RecentActivity []DailyStats `json:"recentActivity"`
}

type DashboardResponse struct {
	Status     bool           `json:"status"`
	Message    string         `json:"message"`
	StatusCode int            `json:"statusCode"`
	Data       DashboardStats `json:"data"`
}
