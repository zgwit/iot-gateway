package model

//Validator 数据校验
type Validator struct {
	Disabled bool `json:"disabled"`

	//条件
	Compare CompareType `json:"compare"`
	Value   string      `json:"value"`
	Value1  float64     `json:"value1"`
	Value2  float64     `json:"value2"`

	//重复日
	DailyChecker

	//延迟报警
	DelayChecker

	//重复报警
	RepeatChecker

	//产生告警
	AlarmContent
}
