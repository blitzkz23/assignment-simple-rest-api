package entity

import "time"

type Order struct {
	ID           uint
	CustomerName string
	OrderedAt    time.Time
}
