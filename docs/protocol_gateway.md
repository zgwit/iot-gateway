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
    "os": "linux",
    "platform": "alpine 3.12",
    "kernel": "5.10.1",
    "boot": 1224212, //启动时间
    "cpu": {
      "cores": 4,
      "usage": 71,
      "model": "intel i5 4100"
    },
    "memory": {
      "total": 40404043,
      "free": 3900333,
      "used": 20230032,
      "usage": 50,
    }
  }
}
```