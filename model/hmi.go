package model

import "time"

type Component struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Group   string    `json:"group"`
	Version string    `json:"version"` //semver
	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	//Deleted time.Time `json:"-" xorm:"deleted"`
}

type Hmi struct {
	Id      string    `json:"id" xorm:"pk"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Updated time.Time `json:"updated" xorm:"updated"`
	Created time.Time `json:"created" xorm:"created"`
	//Deleted time.Time `json:"-" xorm:"deleted"`
}
