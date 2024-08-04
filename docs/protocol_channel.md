# IoT通讯协议 之 通道

本节规定了连接的交互内容

## 1、查询

```json5
{
  "module": "channel",
  "command": "list",
}
```


查询反馈

```json5
{
  "result": "ok/fail",  //结果
  "reason": "密码错误",  //错误原因
  "data": [
    {
      "name": "串口1",      //名称
      "type": "serial",      //类型
      "options": {
        "port_name": "COM1",
        "baud_rate": 9600,
        "data_bits": 7,
        "stop_bits": 1,
        "parity_mode": "N"
      }      
    },
    {
      "name": "S7-Smart200",
      "type": "tcp-client",
      "options": {
        "server": "192.168.31.22",
        "port": 1332,
      }      
    }
  ]
}
```

## 2、打开通道

创建或更新，同名的通道会被替换

```json5
{
  "module": "channel",
  "command": "open",
  "data": {
    "name": "S7-Smart200",
    "type": "tcp-client",
    "options": { //参数
      "server": "192.168.31.22",
      "port": 1332,
    }
  }
}
```

##  3、关闭通道
```json5
{
  "module": "channel",
  "command": "close",
  "data": {
    "name": "S7-Smart200",
  }
}
```

##  4、重启通道
```json5
{
  "module": "channel",
  "command": "restart",
  "data": {
    "name": "S7-Smart200",
  }
}
```

##  5、监听通道数据
```json5
{
  "module": "channel",
  "command": "watch",
  "data": {
    "name": "S7-Smart200",
  }
}
```

监听响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": {
    "stream": 13, //数据流ID
  }
}
```

##  6、透传通道数据
```json5
{
  "module": "channel",
  "command": "pipe",
  "data": {
    "name": "S7-Smart200",
  }
}
```

透传响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": {
    "stream": 14, //数据流ID
  }
}
```