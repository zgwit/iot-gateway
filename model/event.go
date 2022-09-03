package model

import "time"

type Event struct {
	Id       int64     `json:"id"`
	Event    string    `json:"event"`
	Target   string    `json:"target"`
	TargetId int64     `json:"target_id" xorm:"index"`
	Created  time.Time `json:"created" xorm:"created"`
}
