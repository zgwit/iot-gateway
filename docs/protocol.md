# IoT通讯协议

本文规定了网关与云平台之间的通讯协议

## 一、包头

### 固定头 4字节

| 标识   | 长度（位） | 说明    |
|------|-------|-------|
| 魔术字  | 24    | “rpc” |
| 包含数据 | 1     | 0/1   |
| 数据结束 | 1     | 0/1   |
| 数据格式 | 2     | 0-3   |
| 包类型  | 3     | 0-7   |

### 扩展头 4字节

| 标识   | 长度（位） | 说明      |
|------|-------|---------|
| 序号   | 16    | 0-65535 |
| 数据长度 | 16    | 0-64KB  |

### 包类型

| 类型          | id | 扩展头 | 方向 | 说明   |
|-------------|----|-----|----|------|
| -           | 0  |     |    | 无效   |
| CONNECT     | 1  | Y   | 上行 | 连接   |
| CONNECT_ACK | 2  | Y   | 下行 | 连接响应 |
| HEARTBEAT   | 3  |     | 上行 | 心跳   |
| REQUEST     | 4  | Y   | 双向 | 请求   |
| RESPONSE    | 5  | Y   | 双向 | 响应   |
| STREAM      | 6  | Y   | 双向 | 数据流  |
| DISCONNECT  | 7  | Y/N | 双向 | 关闭连接 |

### 数据格式

| 类型          | id | 说明         |
|-------------|----|------------|
| Binary      | 0  | 默认二进制      |
| JSON        | 1  |            |
| YAML        | 2  |            |
| MessagePack | 3  |            |

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

### CONNACK 连接响应

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
