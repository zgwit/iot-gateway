# IoT通讯协议 之 网关

本节规定了网关状态和配置的交互内容

## 1、状态上报

```json5
{
  "module": "gateway",
  "command": "status",
  "data": {
    //状态 key->value
    "battery": 50,
    "rssi": 90,
    "cpu": 27,
    "mem": 39,
  }
}
```


## 2、事件上报

```json5
{
  "module": "gateway",
  "command": "event",
  "data": {
    "name": "电源异常", //事件名
    "type": "系统", //类型
    "level": 1, //等级
  }
}
```

## 3、下发配置

```json5
{
  "module": "gateway",
  "command": "setting",
  "data": {
    //配置项 key->value
    "server": "192.168.31.33",
    "port": 1843,
  }
}
```

配置响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "不支持的配置项",  //错误原因
}
```


## 4、查询状态

```json5
{
  "module": "gateway",
  "command": "metrics",
}
```

配置响应

```json5
{
  "result": "ok",  //结果
  "data": {
    "modules": ["fs", "channel", "product", "device", "camera", "ota"],
    "os": "linux",
    "platform": "ubuntu 22.04",
    "kernel": "5.10.1",
    "boot": 1224212, //启动时间
    "cpu": {
      "cores": 4,
      "usage": 71,
      "mhz": 3000,
      "model": "intel i5 4100"
    },
    "memory": {
      "total": 40404043,
      "free": 3900333,
      "used": 20230032,
      "usage": 50,
    },
    "net":[
      {
        "name": "eth0", //连接名称
        "mac": "0d:0c:00:0d:0c:00", //mac地址
        "flags": ["up", "loopback"],
        "address": [ //ip地址表
          "192.168.0.12",
          "9d23:234234...",
        ],
        "tx":420112, //发送的数据
        "rx":62001, //接收的数据
      }
    ],
    "disk": [
      {
        "name": "/sd0",
        "mount": "/usr/local",
        "type": "fat32",
        "total": 1300000,
        "used": 100232,
        "free": 1200030,
        "usage": 50,
      }
    ]
  }
}
```
