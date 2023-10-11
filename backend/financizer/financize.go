package finance

import (
	"backend/comm"
	"backend/parser"
	"fmt"
	"strconv"
	"strings"

	"github.com/apaxa-go/eval"
)

func cleanData(column string) int {
	column = strings.ReplaceAll(column, ",", "")
	column = strings.ReplaceAll(column, " ", "")

	i, err := strconv.Atoi(column)

	if err != nil {
		return -1
	}

	return i
}

func trimLeftChar(s string) (string, string) {

	if len(s) > 1 {
		return string(s[0]), s[1:]
	} else if len(s) == 1 {
		return string(s[0]), ""
	}

	return "", ""
}

func isOperation(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}

	return false
}

func removeSpaces(label string) string {
	return strings.ReplaceAll(label, " ", "")
}

func createMetrics(allMetrics []*comm.Metric, tables []*parser.Table) []*comm.Metric {

	for _, v := range tables {
		for tr, td := range v.Tr {
			if tr == "" {
				continue
			}

			// CLEAN LABEL
			for _, metric := range allMetrics {
				for _, subMetric := range metric.Submetric {
					label := comm.CleanString([]byte(subMetric.Label))
					if label == tr {
						for _, column := range td {
							data := cleanData(column)
							if data > 0 {
								subMetric.Value = float64(data)
								break
							}
						}
					}
				}
			}

		}
	}

	comm.OrderSubmetrics(allMetrics)

	// execute metrics
	for _, metric := range allMetrics {
		executionString := ""

		fullOperation := metric.Operation
		var char string
		char, fullOperation = trimLeftChar(fullOperation)
		currentMetric := 0
		for char != "" {
			subMetric := metric.Submetric[currentMetric]

			if char == "(" {
				executionString += "("
			} else if char == "a" {

				executionString += fmt.Sprintf("%f", subMetric.Value)
			} else if isOperation(char) {
				executionString += subMetric.Operation
				currentMetric += 1
			} else if char == ")" {
				executionString += ")"
			}

			char, fullOperation = trimLeftChar(fullOperation)
		}

		// fmt.Println(executionString)

		expr, err := eval.ParseString(executionString, "")

		if err != nil {
			metric.Value = -1
			continue
		}

		r, err := expr.EvalToInterface(nil)

		if err != nil {
			metric.Value = -1
			continue
		}

		metric.Value = r.(float64)

	}

	return allMetrics
}

func processTables(tables []*parser.Table, metrics []*comm.Metric) []*comm.Metric {
	// var buf bytes.Buffer // execute template on byte buffer

	// metrics := testMetric()

	createMetrics(metrics, tables)

	// statistics := CreateStatistics()
	// statistics.initAllStatistics()

	// for _, v := range tables {
	// 	for tr, td := range v.Tr {
	// 		statistics.checkForMeasureOfWorkingCapital(tr, td)

	// 	}
	// }

	// statistics.calculateAll()

	return metrics
}
