package finance

import (
	"backend/comm"
	"backend/database"
	"backend/parser"
)

func getMetricsWithTag(tagName string) []*comm.Metric {

	var metrics []*comm.Metric

	database.
		Db.
		Preload("Submetric").
		Preload("Tags").
		Where("document_id = ?", 0).
		Find(&metrics)

	var newMetrics []*comm.Metric

	for _, metric := range metrics {
		for _, tag := range metric.Tags {
			if tag.Name == tagName {
				newMetrics = append(newMetrics, metric)
			}
		}
	}

	return newMetrics
}

func InitFinancer(tables []*parser.Table, tagName string) []*comm.Metric {
	metrics := getMetricsWithTag(tagName)

	if len(metrics) == 0 {
		return metrics
	}

	return processTables(tables, metrics)
}
