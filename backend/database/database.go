package database

import (
	"backend/comm"
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetDocumentID(metrics []*comm.Metric) {
	for _, metric := range metrics {
		idGenerator := uuid.New()
		metric.ID = uint(idGenerator.ID())
	}
}

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

func AddMetrics(metrics []*comm.Metric) {
	// Set the DocumentID for each Metric
	for _, metric := range metrics {
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
		log.Fatal(err)
	}
	databaseExists := true

	pathd := path.Join(dir, "database", "database.db")
	// if file already existed, don't add metrics
	if _, err := os.Stat(pathd); os.IsNotExist(err) {
		databaseExists = false
		os.Create(pathd)
		log.Info("Database created")
	}

	Db, err = gorm.Open(sqlite.Open(pathd), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
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
	log.Info("Database filled")

	metrics, name := StatusMetrics()
	addMetricsToDatabase(metrics, name)

	metrics, name = BaseMetrics()
	addMetricsToDatabase(metrics, name)
}

func addMetricsToDatabase(metrics []*comm.Metric, name string) {
	for _, metric := range metrics {
		metric.Tags = append(metric.Tags, &comm.Tag{
			Name: "All",
		})
	}

	Db.Save(metrics)
}

func BaseMetrics() ([]*comm.Metric, string) {

	allMetrics := []*comm.Metric{}

	name := "Base"

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Long-term Debt",
		Value:       0,
		Explanation: "This is the long term debt of a company",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "long-term debt",
				Value:       0,
				Explanation: "long term debt of the company",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	})

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Total Revenue",
		Value:       0,
		Explanation: "This is the total amount of money that the company made. (before taxes n all that)",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Total Revenue",
				Value:       0,
				Explanation: "total revenue",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	})

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Total Expenses",
		Value:       0,
		Explanation: "This is the total amount of expenses (before taxes n all that)",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{

			&comm.Submetric{
				Order:       1,
				Label:       "Total Expenses",
				Value:       0,
				Explanation: "total revenue",
				Operation:   "",
			},
		},
		Tags: []*comm.Tag{
			&comm.Tag{
				Name: name,
			},
		},
	})

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Operating income",
		Value:       0,
		Explanation: "Sum total of company's profit after subtracting its regular recurring costs and expenses",
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
	})

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Net Income",
		Value:       0,
		Explanation: "Total amount of income",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{
			&comm.Submetric{
				Order:       1,
				Label:       "Net Income",
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
	})

	allMetrics = append(allMetrics, &comm.Metric{
		Label:       "Total Shareholder's Equity",
		Value:       0,
		Explanation: "Total amount of equity by shareholders",
		Operation:   "a",
		Rating:      "",
		Submetric: []*comm.Submetric{
			&comm.Submetric{
				Order:       1,
				Label:       "Total Shareholder's Equity",
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
	})

	return allMetrics, "Base"
}

func StatusMetrics() ([]*comm.Metric, string) {

	allMetrics := []*comm.Metric{}

	name := "Status"

	Mowc := comm.Metric{
		Label:       "Working Capital",
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

	allMetrics = append(allMetrics, &Mowc)
	allMetrics = append(allMetrics, &Mowc2)

	return allMetrics, "Status"
}
