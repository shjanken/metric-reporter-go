package domain

// General 包含了主要就监控项目
type General struct {
	Visitors  int `json:"unique_visitors"`
	Requests  int `json:"valid_requests"`
	Bandwidth int `json:"bandwidth"`
}

// Metric 需报告的监控项目
type Metric struct {
	General `json:"general"`
}
