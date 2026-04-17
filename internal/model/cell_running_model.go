package model

import "time"

type CellRunningModel struct {
	ID        int
	CellID    int
	ModelID   int
	StartTime time.Time
}
