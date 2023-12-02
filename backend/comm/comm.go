package comm

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

var ToTablizer chan []byte

type Document struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Report    []byte
	Metrics   []*Metric `gorm:"foreignKey:DocumentID"`
	Name      string
}

func (d Document) MarshalJSON() ([]byte, error) {

	basicLink := &struct {
		ID        uint      `json:"id"`
		CreatedAt string    `json:"createdat"`
		UpdatedAt string    `json:"updatedat"`
		DeletedAt string    `json:"deletedat"`
		Report    []byte    `json:"report"`
		Metrics   []*Metric `json:"metrics"`
		Name      string    `json:"string"`
	}{
		ID:        d.ID,
		CreatedAt: d.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: d.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: d.DeletedAt.Time.String(),
		Report:    d.Report,
		Metrics:   d.Metrics,
		Name:      d.Name,
	}

	return json.Marshal(basicLink)
}

func OrderSubmetrics(metrics []*Metric) {
	for _, metric := range metrics {
		sort.Sort(ASubmetric(metric.Submetric))
	}
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
	Order       int
}

func (sm Submetric) MarshalJSON() ([]byte, error) {

	basicLink := &struct {
		ID          uint    `json:"id"`
		Label       string  `json:"label"`
		Value       float64 `json:"value"`
		Explanation string  `json:"explanation"`
		Operation   string  `json:"operation"`
	}{
		ID:          sm.ID,
		Label:       sm.Label,
		Value:       sm.Value,
		Explanation: sm.Explanation,
		Operation:   sm.Operation,
	}

	return json.Marshal(basicLink)
}

type ASubmetric []*Submetric

func (sm ASubmetric) Len() int {
	return len(sm)
}

func (sm ASubmetric) Less(i, j int) bool {
	return sm[i].Order < sm[j].Order
}

func (sm ASubmetric) Swap(i, j int) {
	sm[i], sm[j] = sm[j], sm[i]
}

type Metric struct {
	ID          uint `gorm:"primaryKey"`
	DocumentID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Label       string
	Value       float64
	Explanation string
	Operation   string
	Rating      string
	Submetric   []*Submetric `gorm:"foreignKey:MetricID"`
	Tags        []*Tag       `gorm:"foreignKey:MetricID"`
}

func (m Metric) MarshalJSON() ([]byte, error) {

	basicLink := &struct {
		ID          uint         `json:"id"`
		DocumentID  uint         `json:"documentId"`
		Label       string       `json:"label"`
		Value       float64      `json:"value"`
		Explanation string       `json:"explanation"`
		Operation   string       `json:"operation"`
		Rating      string       `json:"rating"`
		Submetric   []*Submetric `json:"submetric"`
	}{
		ID:          m.ID,
		DocumentID:  m.DocumentID,
		Label:       m.Label,
		Value:       m.Value,
		Explanation: m.Explanation,
		Operation:   m.Operation,
		Rating:      m.Rating,
		Submetric:   m.Submetric,
	}

	return json.Marshal(basicLink)
}

type Tag struct {
	gorm.Model
	Name string
	// Metrics  []*Metric `gorm:"many2many:metric_tag;"`
	MetricID uint
}

func (t Tag) MarshalJSON() ([]byte, error) {
	basicLink := &struct {
		ID        uint   `json:"id"`
		CreatedAt string `json:"createdat"`
		UpdatedAt string `json:"updatedat"`
		DeletedAt string `json:"deletedat"`
		Name      string `json:"name"`
	}{
		ID:        t.ID,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: t.DeletedAt.Time.String(),
		Name:      t.Name,
	}

	return json.Marshal(basicLink)
}

func CleanString(cleanB []byte) string {
	clean := string(cleanB)
	clean = strings.ToLower(clean)
	// perform lemmatization here

	clean = strings.ReplaceAll(clean, " ", "")
	clean = strings.ReplaceAll(clean, "\n", "")
	clean = strings.ReplaceAll(clean, "\r", "")
	clean = strings.ReplaceAll(clean, "\t", "")

	var result []rune
	for _, r := range clean {
		if r < 128 {
			result = append(result, r)
		}
	}

	return string(result)
}
