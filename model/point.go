package model

// Point 数据点
type Point struct {
	Name    string   `json:"name"`
	Label   string   `json:"label"`
	Unit    string   `json:"unit"`
	Type    DataType `json:"type"`
	LE      bool     `json:"le"` //little endian
	Dot     int      `json:"dot"`
	Area    string   `json:"area"`
	Address string   `json:"address"`
}
