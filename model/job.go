package model

import (
	"time"
)

//Job 任务
type Job struct {
	Disabled bool `json:"disabled,omitempty"`

	Clock    int            `json:"clock"`
	Weekdays []time.Weekday `json:"weekdays,omitempty"`

	Invokes []Invoke `json:"invokes"`
}
