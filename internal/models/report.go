package models

import "time"

// ReportType represents
type ReportType int64

// ImportReportType is the ReportType for importing recipes, either from files or the web.
const ImportReportType ReportType = 1

// NewReport creates a new, initialized and empty Report of the given ReportType.
func NewReport(reportType ReportType) Report {
	return Report{
		CreatedAt: time.Now(),
		Logs:      make([]ReportLog, 0),
		Type:      reportType,
	}
}

// NewReportLog creates a new ReportLog from the title and error.
func NewReportLog(title string, err error) ReportLog {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	return ReportLog{
		Error:     errStr,
		IsSuccess: err == nil,
		Title:     title,
	}
}

// Report holds information on a report.
type Report struct {
	CreatedAt time.Time
	ExecTime  time.Duration
	ID        int64
	Logs      []ReportLog
	Type      ReportType
}

// ReportLog holds information on a report's log.
type ReportLog struct {
	Error     string
	ID        int64
	IsSuccess bool
	Title     string
}
