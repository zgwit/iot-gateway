# IoT通讯协议 之 设备

本节规定了设备的交互内容

## 1、查询

```json5
{
  "module": "device",
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
      "id": "00101",
      "name": "温度计1",
      "product_id": "001",
      "station": {
        "slave": 1, //从站号
      }
    },
  ]
}
```


## 2、创建或更新设备

```json5
{
  "module": "device",
  "command": "replace", //create update
  "data": {
    "id": "00101",
    "name": "温度计1",
    "product_id": "001",
    "station": {
      "slave": 1, //从站号
    }
  }
}
```


##  3、删除设备
```json5
{
  "module": "device",
  "command": "delete",
  "data": {
    "id": "001",
  }
}
```

## 4、设备数据上报

```json5
{
  "module": "device",
  "command": "property",
  "data": {
    "id": "00101",
    "product_id": "001",
    "properties": {
      "temperature": 31.6, //温度
    },
  }
}
```


## 5、设备数据修改

```json5
{
  "module": "device",
  "command": "property",
  "data": {
    "id": "00101",    
    "properties": {
      "temperature": 31.6, //温度
    },
  }
}
```

## 6、设备事件上报

```json5
{
  "module": "device",
  "command": "event",
  "data": {
    "id": "00101",
    "name": "温度异常", //事件名
    "type": "运行", //类型
    "level": 1, //等级
  }
}
```


## 7、设备操作

```json5
{
  "module": "device",
  "command": "action",
  "data": {
    "id": "00101",
    "name": "open", //操作名
    "parameters": {
      //参数
    }
  }
}
```

响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "不支持的操作",  //错误原因
  "data": {
    "return": {
      //返回值
    }
  }
}
```