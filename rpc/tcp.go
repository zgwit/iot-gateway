package rpc

const (
	TCP_LIST uint8 = iota
	TCP_CREATE
	TCP_DELETE
	TCP_RESTART
	TCP_WATCH
	TCP_PIPE
)

type TcpId struct {
	Id string `json:"id"`
}

type TcpItem struct {
	Id      string     `json:"id"`
	Name    string     `json:"name,omitempty"`
	Options TcpOptions `json:"options"`
}

type TcpOptions struct {
	IsServer bool   `json:"is_server,omitempty"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type TcpListRequest struct{}

type TcpListResponse []TcpItem

type TcpCreateRequest TcpItem

type TcpCreateResponse struct{}

type TcpDeleteRequest TcpId

type TcpDeleteResponse struct{}

type TcpRestartRequest TcpId

type TcpRestartResponse struct{}

type TcpWatchRequest TcpId

type TcpWatchResponse StreamId

type TcpPipeRequest TcpId

type TcpPipeResponse StreamId
