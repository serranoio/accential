package lexer

import (
	"backend/parser"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/apaxa-go/eval"
)

type MeasureOfWorkingCapital struct {
	TotalAssets      []int
	TotalLiabilities []int
	WorkingCapital   []float64
	Info             string
}

type Statistics struct {
	Mowc *MeasureOfWorkingCapital
}

func CreateStatistics() *Statistics {
	return &Statistics{
		Mowc: nil,
	}
}

func (s *Statistics) AddMeasureOfWorkingCapital() {
	s.Mowc = &MeasureOfWorkingCapital{
		TotalAssets:      []int{},
		TotalLiabilities: []int{},
		WorkingCapital:   []float64{},
		Info:             "",
	}
}

func (s *Statistics) initAllStatistics() {
	s.AddMeasureOfWorkingCapital()

}

func cleanData(column string) int {
	column = strings.ReplaceAll(column, ",", "")
	column = strings.ReplaceAll(column, " ", "")

	i, err := strconv.Atoi(column)

	if err != nil {
		return -1
	}

	return i
}

func (s *Statistics) checkForMeasureOfWorkingCapital(tr string, td parser.Td) {
	if tr == "Total Assets" {
		for _, column := range td {
			data := cleanData(column)
			if data > 0 {
				s.Mowc.TotalAssets = append(s.Mowc.TotalAssets, data)
			}
		}
	} else if tr == "Total Liabilities" {
		for _, column := range td {
			data := cleanData(column)
			if data > 0 {
				s.Mowc.TotalLiabilities = append(s.Mowc.TotalLiabilities, data)
			}
		}
	}
}

func (s *Statistics) calculateAll() {
	for i := 0; i < len(s.Mowc.TotalAssets); i++ {
		s.Mowc.WorkingCapital = append(
			s.Mowc.WorkingCapital,
			float64(s.Mowc.TotalAssets[i])/float64(s.Mowc.TotalLiabilities[i]))
	}

	s.Mowc.Info = "The Measure of Working Capital is the Ratio of Total Assets / Total Liablities"
}

type Metric struct {
	Label       string
	Value       float64
	Explanation string
	Operation   string
	Metric      []*Metric
}

func testMetric() []*Metric {
	allMetrics := []*Metric{}

	Mowc := Metric{
		Label:       "Measure of Working Capital",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Metric: []*Metric{

			&Metric{
				Label:       "Total Assets",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
				Metric:      nil,
			},
			&Metric{
				Label:       "Total Liabilities",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
				Metric:      nil,
			},
		},
	}

	Mowc2 := Metric{
		Label:       "Revenue / Expenses",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Metric: []*Metric{

			&Metric{
				Label:       "Total Revenue",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
				Metric:      nil,
			},
			&Metric{
				Label:       "Total Expenses",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
				Metric:      nil,
			},
		},
	}

	Mowc3 := Metric{
		Label:       "Tax Percentage",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Metric: []*Metric{

			&Metric{
				Label:       "Taxes payable",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "/",
				Metric:      nil,
			},
			&Metric{
				Label:       "Net income",
				Value:       0,
				Explanation: "Taxes...",
				Operation:   "",
				Metric:      nil,
			},
		},
	}

	Mowc4 := Metric{
		Label:       "Operating income",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "a",
		Metric: []*Metric{

			&Metric{
				Label:       "Operating income",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "",
				Metric:      nil,
			},
		},
	}

	allMetrics = append(allMetrics, &Mowc)
	allMetrics = append(allMetrics, &Mowc2)
	allMetrics = append(allMetrics, &Mowc3)
	allMetrics = append(allMetrics, &Mowc4)

	return allMetrics
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

func createMetrics(allMetrics []*Metric, tables []*parser.Table) []*Metric {

	for _, v := range tables {
		for tr, td := range v.Tr {

			for _, metric := range allMetrics {
				for _, subMetric := range metric.Metric {
					if subMetric.Label == tr {
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
	// execute metrics
	for _, metric := range allMetrics {
		executionString := ""

		fullOperation := metric.Operation
		var char string
		char, fullOperation = trimLeftChar(fullOperation)
		currentMetric := 0
		for char != "" {
			subMetric := metric.Metric[currentMetric]

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
			log.Panic(err)
		}

		r, err := expr.EvalToInterface(nil)

		if err != nil {
			log.Panic(err)
		}

		metric.Value = r.(float64)

	}

	return allMetrics
}

func processTables(tables []*parser.Table) []*Metric {

	allMetrics := createMetrics(testMetric(), tables)

	// statistics := CreateStatistics()
	// statistics.initAllStatistics()

	// for _, v := range tables {
	// 	for tr, td := range v.Tr {
	// 		statistics.checkForMeasureOfWorkingCapital(tr, td)

	// 	}
	// }

	// statistics.calculateAll()

	return allMetrics
}
