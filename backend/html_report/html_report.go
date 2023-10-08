package html_report

import "backend/comm"

func InitHtmlReport(metrics []*comm.Metric, htmlReport []byte) uint {
	return createReport(metrics, htmlReport)
}
