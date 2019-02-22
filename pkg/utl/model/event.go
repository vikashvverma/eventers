package eventers

import (
	"time"
)

type Event struct {
	ID       string    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
}
