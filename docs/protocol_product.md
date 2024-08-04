# IoT通讯协议 之 产品（物模型）

本节规定了产品的交互内容

## 1、查询

```json5
{
  "module": "product",
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
      "id": "001",
      "name": "温度计",    //名称
      "version": "1.0.0" //版本号
    },
  ]
}
```


## 2、下载产品

下载zip文件

```json5
{
  "module": "product",
  "command": "download",
  "data": {
    "id": "001",
  }
}
```

响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": {
    "stream": 13, //数据流ID
  }
}
```

## 3、上传产品

上传zip文件，用于创建或更新，同名的产品会被替换

```json5
{
  "module": "product",
  "command": "upload",
  "data": {
    "id": "001",
  }
}
```

响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": {
    "stream": 14, //数据流ID
  }
}
```


##  4、删除产品
```json5
{
  "module": "product",
  "command": "delete",
  "data": {
    "id": "001",
  }
}
```
