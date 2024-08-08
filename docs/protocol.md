# IoT通讯协议

本文规定了网关与云平台之间的通讯协议

## 一、包头

### 固定头 8字节

| 标识     | 长度（位） | 说明         |
|--------|-------|------------|
| 魔术字    | 24    | “rpc”      |
| 包类型    | 4     | 0-15       |
| 数据格式   | 4     | 0-15       |
| 事务/流ID | 16    | 0-65535，自增 |
| 包体数据长度 | 16    | 0-64KB     |


### 包类型

| id | 类型            | 方向 | 说明    |
|----|---------------|----|-------|
| 0  | DISCONNECT    | 双向 | 关闭连接  |
| 1  | CONNECT       | 上行 | 连接    |
| 2  | CONNECT_ACK   | 下行 | 连接响应  |
| 3  | HEARTBEAT     | 上行 | 心跳    |
| 4  | REQUEST       | 双向 | 请求    |
| 5  | REQUEST_END   | 双向 | 请求结束  |
| 6  | RESPONSE      | 双向 | 响应    |
| 7  | RESPONSE_END  | 双向 | 响应结束  |
| 8  | STREAM        | 双向 | 数据流   |
| 9  | STREAM_END    | 双向 | 数据流结束 |
| 10 | PUBLISH       | 双向 | 发布    |
| 11 | PUBLISH_END   | 双向 | 发布结束  |
| 12 | PUBLISH_ACK   | 双向 | 发布响应  |
| 13 | SUBSCRIBE     | 上行 | 订阅    |
| 14 | SUBSCRIBE_ACK | 下行 | 订阅响应  |
| 15 | UNSUBSCRIBE   | 上行 | 取消订阅  |

### 数据格式

| id   | 类型       | 说明         |
|------|----------|------------|
| 0    | binary   | 默认二进制      |
| 1    | json     | 通用性强       |
| 2    | xml      |            |
| 3    | yaml     | 较JSON省流，清晰 |
| 4    | csv      |            |
| 5    | msgpack  | 省流，高性能     |
| 6    | protobuf |            |
| 7-15 | custom   | 自定义        |

数据格式建议使用默认的json格式，方便开发和调试，对于需要省流的场景，可以改进为msgpack或二进制


## 二、基础交互说明

### CONNECT 连接

首次连接

```json5
{
  "id": "123123123123", //客户端ID，可选
  "username": "user", //用户名
  "password": "123456", //密码
}
```

再次连接，使用token
```json5
{
  "id": "123123123123", //客户端ID
  "token": "xyzxyzxyzxyz", //登录证书
}
```

### CONNECT_ACK 连接响应

```json5
{
  "result": "ok/fail", //结果
  "reason": "密码错误", //错误原因
  "id": "123123123123", //客户端ID
  "token": "abcabcabcabc", //登录证书，客户端需要保存
}
```

### HEARTBEAT 心跳
无内容，无需响应

### STREAM 数据流
内容为二进制

流ID，客户端发起用奇数，服务端发起用偶数

### STREAM_END 数据流结束
内容为二进制

### REQUEST 请求

```json5
{
  "module": "fs", //模块
  "command": "create", //命令
  "data": {}, //数据
}
```

### RESPONSE 响应
```json5
{
  "result": "ok/fail", //结果
  "reason": "密码错误", //错误原因
  "data": {} //自定义数据响应
}
```
详情参见具体交互说明


### DISCONNECT 关闭连接
```json5
{
  "reason": "restart", //关闭原因
}
```

## 三、具体交互说明

- [网关](protocol_gateway.md)
- [连接通道](protocol_channel.md)
- [产品，物模型](protocol_product.md)
- [设备](protocol_device.md)
- [OTA升级](protocol_ota.md)
- [文件系统](protocol_fs.md)
- [摄像头](protocol_camera.md)
