package rpc

type DeviceId struct {
	Id string `json:"id"`
}

type DeviceItem struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	ProductId string         `json:"product_id"`
	Station   map[string]any `json:"station,omitempty"`
}

type DeviceListRequest struct{}

type DeviceListResponse []DeviceItem

type DeviceCreateRequest DeviceItem

type DeviceCreateResponse struct{}

type DeviceDeleteRequest DeviceId

type DeviceDeleteResponse struct{}

type DevicePropertyRequest struct {
	Id         string         `json:"id"`
	ProductId  string         `json:"product_id,omitempty"`
	Properties map[string]any `json:"properties"`
}

type DevicePropertyResponse struct{}

type DevicePropertyModifyRequest struct {
	Id         string         `json:"id"`
	Properties map[string]any `json:"properties"`
}

type DevicePropertyModifyResponse struct{}

type DeviceEventRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Level int    `json:"level,omitempty"`
}

type DeviceEventResponse struct{}

type DeviceActionRequest struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Parameters map[string]any `json:"parameters"`
}

type DeviceActionResponse struct {
	Return map[string]any `json:"return,omitempty"`
}
