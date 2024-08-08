package rpc


const (
	DISCONNECT uint8 = iota
	CONNECT       
	CONNECT_ACK   
	HEARTBEAT   
	REQUEST      
	REQUEST_END   
	RESPONSE      
	RESPONSE_END 
	STREAM       
	STREAM_END    
	PUBLISH       
	PUBLISH_END   
	PUBLISH_ACK    
	SUBSCRIBE     
	SUBSCRIBE_ACK 
	UNSUBSCRIBE  
)

const (
	BINARY uint8 = iota
	JSON
	XML
	YAML
	CSV
	MSGPACK
	PROTOBUF
)


type Pack struct {
	Type uint8
	Encoding uint8
	Id uint16
	Length uint16
}