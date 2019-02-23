package eventers

import (
	"time"
)

type Event struct {
	ID       int64     `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Location string    `json:"location,omitempty"`
}
