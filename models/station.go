// models/station.go
package models

type Station struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NAME      string    `json:"name"`
	LINE      string    `json:"line"`
	X		  *int64    `json:"x"`
	Y		  *int64    `json:"y"`
}