# iot-gateway
物联大师网关

[![Go](https://github.com/zgwit/iot-gateway/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-gateway/actions/workflows/go.yml)
[![Node.js](https://github.com/zgwit/iot-gateway/actions/workflows/node.js.yml/badge.svg)](https://github.com/zgwit/iot-gateway/actions/workflows/node.js.yml)
[![CodeQL](https://github.com/zgwit/iot-gateway/actions/workflows/codeql.yml/badge.svg)](https://github.com/zgwit/iot-gateway/actions/workflows/codeql.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zgwit/iot-gateway.svg)](https://pkg.go.dev/github.com/zgwit/iot-gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/zgwit/iot-gateway)](https://goreportcard.com/report/github.com/zgwit/iot-gateway)


## 网关通讯协议

网关与云平台之间使用专用的通讯协议

### 数据包

| 标识 | 长度（位） | 说明 |
|----|----|----|
| 魔术字 | 24 | “iot” |
| 包含数据 | 1 | 0-1 |
| 数据格式 | 3 | 0-7 |
| 包类型 | 4 | 1-15 |
| [事务序号] | 16 | 0-65535 |
| [数据长度] | 16 | 0-64KB |


### 包类型

| 类型 | id | 说明 |
|----|----|----|
| - | 0 | 无效 |
| CONNECT | 1 | 连接 |
| CONNACK | 2 | 连接响应 |
| HEARTBEAT | 3 | 心跳 |
| PING | 4 | ping |
| PONG | 5 | ping响应 |
| REQUEST | 6 | 请求 |
| RESPONSE | 7 | 响应 |
| STREAM | 8 | 数据流 |
| STREAM_END | 9 | 数据流结束 |
| PUBLISH | 10 | 广播 |
| PUBACK | 11 | 广播响应 |
| SUBSCRIBE | 12 | 订阅 |
| SUBACK | 13 | 订阅响应 |
| UNSUBSCRIBE | 14 | 取消订阅 |
| DISCONNECT | 15 | 断开连接 |

### 数据格式

| 类型 | id | 说明 |
|----|----|----|
| Binary | 0 | 未定义 |
| JSON | 1 | |
| YAML | 2 | |
| XML | 3 | |
| CSV | 4 | |
| Protobuf | 5 | |
| MessagePack | 6 | |
| Reserved | 7 | 保留 |

## 交互说明





## 协议支持

- [x] Modbus协议（内置）
- [x] [西门子 S7 PLC](https://github.com/iot-master-contrib/s7)
- [x] [三菱 PLC](https://github.com/iot-master-contrib/melsec)
- [x] [欧姆龙 PLC](https://github.com/iot-gateway-contrib/fins)
- [ ] CJ/T188-2004、2018 户用计量仪表数据传输技术条件
- [x] DL/T645-1997、2007 多功能电表通讯规约
- [ ] IEC 101/103/104 电力系统远程控制和监视的通信协议
- [ ] IEC 61850 电力系统自动化领域全球通用协议
- [ ] SL/T427-2021 水资源监测数据传输规约
- [ ] SL/T651-2014 水文监测数据通信规约
- [ ] SL/T812.1-2021 水利监测数据传输规约
- [ ] SZY206-2016 水资源监测数据传输规约
- [ ] BACnet智能建筑协议
