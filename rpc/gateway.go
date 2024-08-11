package rpc

type GatewayStatusRequest struct {
	Battery int `json:"battery,omitempty"`
	Rssi    int `json:"rssi,omitempty"`
	Cpu     int `json:"cpu,omitempty"`
	Mem     int `json:"mem,omitempty"`
}

type GatewayStatusResponse struct{}

type GatewayEventRequest struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Level int    `json:"level,omitempty"`
}

type GatewaySettingRequest struct {
	Server string `json:"server,omitempty"`
	Port   uint16 `json:"port,omitempty"`
}

type GatewaySettingResponse struct{}

type GatewayMetricsRequest struct{}

type GatewayMetricsResponse struct {
	Modules  []string                `json:"modules,omitempty"`
	Os       string                  `json:"os,omitempty"`
	Platform string                  `json:"platform,omitempty"`
	Kernel   string                  `json:"kernel,omitempty"`
	Boot     int64                   `json:"boot,omitempty"`
	Cpu      GatewayMetricsCpu       `json:"cpu"`
	Memory   GatewayMetricsMemory    `json:"memory"`
	Network  []GatewayMetricsNetwork `json:"network,omitempty"`
	Disk     []GatewayMetricsDisk    `json:"disk,omitempty"`
}

type GatewayMetricsCpu struct {
	Cores int    `json:"cores,omitempty"`
	Usage int    `json:"usage,omitempty"`
	Mhz   int    `json:"mhz,omitempty"`
	Model string `json:"model,omitempty"`
}
type GatewayMetricsMemory struct {
	Total int `json:"total,omitempty"`
	Free  int `json:"free,omitempty"`
	Used  int `json:"used,omitempty"`
	Usage int `json:"usage,omitempty"`
}

type GatewayMetricsNetwork struct {
	Name    string   `json:"name,omitempty"`
	Mac     string   `json:"mac,omitempty"`
	Flags   []string `json:"flags,omitempty"`
	Address []string `json:"address,omitempty"`
	Tx      int      `json:"tx,omitempty"`
	Rx      int      `json:"rx,omitempty"`
}

type GatewayMetricsDisk struct {
	Name  string `json:"name,omitempty"`
	Mount string `json:"mount,omitempty"`
	Type  string `json:"type,omitempty"`
	Total int    `json:"total,omitempty"`
	Free  int    `json:"free,omitempty"`
	Used  int    `json:"used,omitempty"`
	Usage int    `json:"usage,omitempty"`
}
