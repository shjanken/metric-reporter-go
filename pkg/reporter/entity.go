package reporter

// System is opperator system
type System struct {
	Chrome  int
	Windows int
	IPhone  int
	Android int
	Mac     int
}

// Metric 需报告的监控项目
type Metric struct {
	Visitors int
	Requests int
	System
}
