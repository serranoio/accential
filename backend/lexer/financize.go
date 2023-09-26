package lexer

import (
	"backend/parser"
	"strconv"
	"strings"
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

func processTables(tables []*parser.Table) *Statistics {

	statistics := CreateStatistics()
	statistics.initAllStatistics()

	for _, v := range tables {
		for tr, td := range v.Tr {
			statistics.checkForMeasureOfWorkingCapital(tr, td)

		}
	}

	statistics.calculateAll()

	return statistics
}
