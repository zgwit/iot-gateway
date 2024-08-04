# IoT通讯协议 之 文件系统

本节规定了文件的交互内容

## 1、查询

```json5
{
  "module": "fs",
  "command": "list",
  "data": {
    "path": "/sd/"
  }
}
```


查询反馈

```json5
{
  "result": "ok/fail",  //结果
  "reason": "无权限",  //错误原因
  "data": [
    {
      "name": "abc",    //名称
      "dir": true,
      "time": 1232132, //时间戳
    },
    {
      "name": "test.txt",    //名称
      "size": 32,
      "time": 1232132, //时间戳
    },
  ]
}
```

## 2、下载文件

下载文件

```json5
{
  "module": "fs",
  "command": "download",
  "data": {
    "path": "/sd/001.txt",
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

## 3、上传文件

上传文件，用于创建或更新，同名的文件会被替换

```json5
{
  "module": "fs",
  "command": "upload",
  "data": {
    "path": "/sd/001.txt",
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


##  4、删除文件（或目录）
```json5
{
  "module": "fs",
  "command": "remove", //rm delete
  "data": {
    "path": "/sd/001.txt",
  }
}
```

##  5、移动文件(支持重命名)
```json5
{
  "module": "fs",
  "command": "move", //mv rename
  "data": {
    "path": "/sd/001.txt",
    "move": "/sd/002.txt",
  }
}
```

##  6、创建目录
```json5
{
  "module": "fs",
  "command": "mkdir",
  "data": {
    "path": "/sd/003",
  }
}
```


##  7、查看磁盘列表
```json5
{
  "module": "fs",
  "command": "disk",
}
```

响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": [
    {
      "disk": "/sd0",
      "total": 1300000,
      "used": 100232,
      "free": 1200030,
    }
  ]
}
```


##  8、格式化
```json5
{
  "module": "fs",
  "command": "format",
  "data": {
    "disk": "/sd0",
    "type": "fat32", //格式
  }
}
```
