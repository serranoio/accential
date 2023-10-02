package lexer

import (
	"backend/database"
	"backend/parser"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/apaxa-go/eval"
	"gorm.io/gorm"
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

type Submetric struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	MetricID    uint
	Label       string
	Value       float64
	Explanation string
	Operation   string
}

func (sm Submetric) MarshalJSON() ([]byte, error) {

	basicLink := &struct {
		ID          uint    `json:"id"`
		CreatedAt   string  `json:"createdat"`
		UpdatedAt   string  `json:"updatedat"`
		DeletedAt   string  `json:"deletedat"`
		Label       string  `json:"label"`
		Value       float64 `json:"value"`
		Explanation string  `json:"explanation"`
		Operation   string  `json:"operation"`
	}{
		ID:          sm.ID,
		CreatedAt:   sm.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   sm.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:   sm.DeletedAt.Time.String(),
		Label:       sm.Label,
		Value:       sm.Value,
		Explanation: sm.Explanation,
		Operation:   sm.Operation,
	}

	return json.Marshal(basicLink)
}

type Metric struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Label       string
	Value       float64
	Explanation string
	Operation   string
	Rating      int
	Submetric   []*Submetric `gorm:"foreignKey:MetricID"`
}

func (m Metric) MarshalJSON() ([]byte, error) {

	basicLink := &struct {
		ID          uint         `json:"id"`
		CreatedAt   string       `json:"createdat"`
		UpdatedAt   string       `json:"updatedat"`
		DeletedAt   string       `json:"deletedat"`
		Label       string       `json:"label"`
		Value       float64      `json:"value"`
		Explanation string       `json:"explanation"`
		Operation   string       `json:"operation"`
		Rating      int          `json:"rating"`
		Submetric   []*Submetric `json:"submetric"`
	}{
		ID:          m.ID,
		CreatedAt:   m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:   m.DeletedAt.Time.String(),
		Label:       m.Label,
		Value:       m.Value,
		Explanation: m.Explanation,
		Operation:   m.Operation,
		Rating:      m.Rating,
		Submetric:   m.Submetric,
	}

	return json.Marshal(basicLink)
}

func testMetric() []*Metric {
	allMetrics := []*Metric{}

	Mowc := Metric{
		Label:       "Measure of Working Capital",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      0,
		Submetric: []*Submetric{

			&Submetric{
				Label:       "Total Assets",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
			},
			&Submetric{
				Label:       "Total Liabilities",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
			},
		},
	}

	Mowc2 := Metric{
		Label:       "Revenue / Expenses",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      0,
		Submetric: []*Submetric{

			&Submetric{
				Label:       "Total Revenue",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
			},
			&Submetric{
				Label:       "Total Expenses",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
			},
		},
	}

	Mowc3 := Metric{
		Label:       "Tax Percentage",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      0,
		Submetric: []*Submetric{

			&Submetric{
				Label:       "Taxes payable",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "/",
			},
			&Submetric{
				Label:       "Net income",
				Value:       0,
				Explanation: "Taxes...",
				Operation:   "",
			},
		},
	}

	Mowc4 := Metric{
		Label:       "Operating income",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "a",
		Rating:      0,
		Submetric: []*Submetric{
			&Submetric{
				Label:       "Operating income",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "",
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
				for _, subMetric := range metric.Submetric {
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

func initDB(metrics []*Metric) {
	database.Db.AutoMigrate(&Metric{})
	database.Db.AutoMigrate(&Submetric{})

	ok := database.Db.Create(metrics)

	fmt.Println(ok)
}

func getAll() []*Metric {
	metrics := []*Metric{}

	records := database.Db.Model(&Metric{}).Preload("Submetric").Find(&metrics)

	if records.Error != nil {
		log.Panic(records.Error)
	}

	return metrics
}

func processTables(tables []*parser.Table) []*Metric {
	// var buf bytes.Buffer // execute template on byte buffer

	// metrics := testMetric()
	metrics := getAll()

	createMetrics(metrics, tables)

	// initDB(metrics)
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
