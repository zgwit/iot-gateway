package model

type Request struct {
	Id string `json:"id"`
}

type RequestList struct {
	Skip     int                    `form:"skip" json:"skip"`
	Limit    int                    `form:"limit" json:"limit"`
	Sort     map[string]int         `form:"sort" json:"sort"`
	Filters  map[string]interface{} `form:"filter" json:"filter"`
	Keywords map[string]string      `form:"keyword" json:"keyword"`
}
