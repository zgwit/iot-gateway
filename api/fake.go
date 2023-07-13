package api

// 为了生成文档 重新 定义

type ParamSearch struct {
	Skip     int                    `form:"skip" json:"skip"`
	Limit    int                    `form:"limit" json:"limit"`
	Sort     map[string]int         `form:"sort" json:"sort"`
	Filters  map[string]interface{} `form:"filter" json:"filter"`
	Keywords map[string]string      `form:"keyword" json:"keyword"`
}

type ParamList struct {
	Skip  int `form:"skip" json:"skip"`
	Limit int `form:"limit" json:"limit"`
}

type ReplyData[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}

type ReplyList[T any] struct {
	Data  []T    `json:"data"`
	Total int64  `json:"total"`
	Error string `json:"error,omitempty"`
}
