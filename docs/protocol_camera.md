# IoT通讯协议 之摄像头

本节规定了摄像头的交互内容

## 1、查询

```json5
{
  "module": "camera",
  "command": "list",
}
```


查询反馈

```json5
{
  "result": "ok/fail",  //结果
  "reason": "",  //错误原因
  "data": [
    {
      "id": "/dev/usb-video0",
      "name": "usb camera",
      "type": "camera", //camera, screen
    },
    {
      "id": "ipc0",
      "name": "IP摄像头1",
      "type": "ipc",
      "url": "rtsp://admin:123456@192.168.0.1",
    },
  ]
}
```


## 2、创建或更新IP摄像头

```json5
{
  "module": "camera",
  "command": "replace", //create update
  "data": {
    "id": "ipc0",
    "name": "IP摄像头1",
    "type": "ipc",
    "url": "rtsp://admin:123456@192.168.0.1",
  }
}
```


##  3、删除IP摄像头
```json5
{
  "module": "camera",
  "command": "delete",
  "data": {
    "id": "ipc0",
  }
}
```

## 4、摄像头拍照

```json5
{
  "module": "camera",
  "command": "take", //photo, snap
  "data": {
    "id": "ipc0",
    "quality": 8, //质量 1-10，可选
    "size": "720P", //尺寸，可选 VGA/480P PAL/576P HD/720P FULL/1080P 2K/1440P 4K/2160P
  }
}
```
反馈

```json5
{
  "result": "ok/fail",  //结果
  "reason": "摄像头异常",  //错误原因
  "data": {
    "path": "/usr/local/DCIM/20240807090205.jpg",
  }
}
```

## 5、摄像头录像

```json5
{
  "module": "camera",
  "command": "record", //video
  "data": {
    "id": "ipc0",
    "duration": 3600, //时长 秒
    "frame": 10, //帧率，可选
    "quality": 8, //质量 1-10，可选
    "size": "720P", //尺寸，可选 VGA/480P PAL/576P HD/720P FULL/1080P 2K/1440P 4K/2160P
  }
}
```

反馈

```json5
{
  "result": "ok/fail",  //结果
  "reason": "摄像头异常",  //错误原因
  "data": {
    "path": "/usr/local/DCIM/20240807090206.mp4",
  }
}
```



## 6、摄像头webrtc桥接

```json5
{
  "module": "camera",
  "command": "webrtc",
  "data": {
    "id": "ipc0",    
  }
}
```


响应

```json5
{
  "result": "ok/fail",  //结果
  "reason": "不支持的操作",  //错误原因
  "data": {
    "stream": 89, //流ID，在流中实现webrtc交互
  }
}
```

webrtc握手流程

| 指令         |    | 内容    | 说明               |
|------------|----|-------|------------------|
| offer      | 上行 | sdp   | 客户端 createAnswer |
| answer     | 下行 | sdp   |                  |
| candidate  | 双向 | ice候选 | ice交换            |

连接成功之后，就可以结束流了

