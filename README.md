# 简介
一个用golang实现的简单在线聊天室,用于演示最基本的golang socket编程方法。
# 用法
启动服务端
```
git clone  https://github.com/remould92/simplechatroom.git
cd simplechatroom/server
go run chatserver.go
```
启动客户端
```
cd simplechatroom/client
go run  chatclient.go
```
部分代码参考了[Socket编程](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/08.1.md
)