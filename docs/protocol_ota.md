# IoT通讯协议 之 OTA

本节规定了OTA的交互内容

## 1、查询

```json5
{
  "module": "ota",
  "command": "version",
}
```


查询反馈

```json5
{
  "result": "ok",  //结果
  "data": {
      "version": "1.0.0" //版本号
  },
}
```


## 2、上传固件

```json5
{
  "module": "ota",
  "command": "upload",
  "data": {
    "type": "full", //类型
    "length": 15645250, //长度
  }
}
```

响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "空间不足",  //错误原因
  "data": {
    "stream": 14, //数据流ID
  }
}
```


##  3、升级结果
```json5
{
  "module": "ota",
  "command": "result",
  "data": {
    "version": "1.0.0" //版本号
  }
}
```
