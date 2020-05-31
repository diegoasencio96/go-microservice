package status

import "time"

type Status struct {
	Version string    `json:"version"`
	Name    string    `json:"name"`
	Git     string    `json:"git"`
	Date    time.Time `json:"date"`
}