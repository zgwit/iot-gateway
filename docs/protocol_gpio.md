# IoT通讯协议 之 GPIO

本节规定了GPIO的交互内容

## 1、查询

```json5
{
  "module": "gpio",
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
      "path": "/dev/gpio0",
      "name": "开关1",
      "type": "io", //io, adc 
      "value": 1,
      "writable": true,
    },
  ]
}
```


## 2、设置GPIO

```json5
{
  "module": "gpio",
  "command": "write",
  "data": {
    "path": "/dev/gpio0",
    "value": 1
  }
}
```


##  3、读GPIO
```json5
{
  "module": "gpio",
  "command": "read",
  "data": {
    "path": "/dev/gpio0",
  }
}
```

响应
```json5
{
  "path": "/dev/gpio0",
  "value": 1,
}
```

## 4、GPIO数据上报

```json5
{
  "module": "gpio",
  "command": "status",
  "data": [
    {
      "path": "/dev/gpio0",
      "value": 1
    },
  ]
}
```
