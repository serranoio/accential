package database

import (
	"backend/comm"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func CloneMetrics(metrics []*comm.Metric) {
	// Set the DocumentID for each Metric
	for _, metric := range metrics {
		idGenerator := uuid.New()
		metric.ID = uint(idGenerator.ID())
		for _, submetric := range metric.Submetric {
			idGenerator := uuid.New()
			submetric.ID = uint(idGenerator.ID())
		}
		for _, tag := range metric.Tags {
			idGenerator := uuid.New()
			tag.ID = uint(idGenerator.ID())
		}
	}
}

func SaveDocument(reportBytes []byte, metrics []*comm.Metric, id uint) uint {

	documentToSave := &comm.Document{
		ID:      id,
		Report:  reportBytes,
		Metrics: metrics,
	}

	CloneMetrics(documentToSave.Metrics)

	Db.Save(documentToSave)

	return documentToSave.ID
}

func InitDatabase() {
	dir, err := os.Getwd()

	if err != nil {
		log.Panicln(err)
	}
	databaseExists := true

	path := path.Join(dir, "database", "database.db")
	// if file already existed, don't add metrics
	if _, err := os.Stat(path); os.IsNotExist(err) {
		databaseExists = false
	}

	Db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	Db.AutoMigrate(&comm.Document{})
	Db.AutoMigrate(&comm.Metric{})
	Db.AutoMigrate(&comm.Submetric{})
	Db.AutoMigrate(&comm.Metric{}, &comm.Tag{})
	Db.AutoMigrate(&comm.Tag{})

	if !databaseExists {
		onRestart()
	}
}

func onRestart() {
	metrics, name := StatusMetrics()
	addMetricsToDatabase(metrics, name)

	metrics, name = EvaluationMetrics()
	addMetricsToDatabase(metrics, name)
}

func addMetricsToDatabase(metrics []*comm.Metric, name string) {
	// tag := comm.Tag{Name: name}
	// for _, metric := range metrics {
	// 	metric.Tags = append(metric.Tags, tag)
	// }

	Db.Save(metrics)
	// Db.Save(tag)
	// Db.Model(&metrics).Association("Tags").Append(tag)
}

func EvaluationMetrics() ([]*comm.Metric, string) {
	allMetrics := []*comm.Metric{}

	name := "Evaluation"

	Mowc := comm.Metric{
		Label:       "My balls",
		Value:       10,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Total Assets",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
			},
			&comm.Submetric{
				Order:       2,
				Label:       "Total Liabilities",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	}

	allMetrics = append(allMetrics, &Mowc)

	return allMetrics, "Evaluation"
}

func StatusMetrics() ([]*comm.Metric, string) {

	allMetrics := []*comm.Metric{}

	name := "Status"

	Mowc := comm.Metric{
		Label:       "Measure of Working Capital",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Total Assets",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
			},
			&comm.Submetric{
				Order:       2,
				Label:       "Total Liabilities",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	}

	Mowc2 := comm.Metric{
		Label:       "Revenue / Expenses",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Total Revenue",
				Value:       0,
				Explanation: "Total assets is...",
				Operation:   "/",
			},
			&comm.Submetric{
				Order:       2,
				Label:       "Total Expenses",
				Value:       0,
				Explanation: "Total liabilities is...",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	}

	Mowc3 := comm.Metric{
		Label:       "Tax Percentage",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "(a/a)",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Taxes payable",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "/",
			},
			&comm.Submetric{
				Order:       2,
				Label:       "Net income",
				Value:       0,
				Explanation: "Taxes...",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	}

	Mowc4 := comm.Metric{
		Label:       "Operating income",
		Value:       0,
		Explanation: "This is the Total Assets / Total Liabilities",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{
			&comm.Submetric{
				Order:       1,
				Label:       "Operating income",
				Value:       0,
				Explanation: "Total income is...",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	}

	allMetrics = append(allMetrics, &Mowc)
	allMetrics = append(allMetrics, &Mowc2)
	allMetrics = append(allMetrics, &Mowc3)
	allMetrics = append(allMetrics, &Mowc4)

	return allMetrics, "Status"
}
