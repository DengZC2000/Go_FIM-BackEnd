# FIM即时通讯BackEnd



##  使用Go语言为后端，使用Docker运行ETCD及Kafka

## 集成QQ第三方登录API

### - 详情可见我的另一个Gin框架项目，有比较详细的介绍：[Gin博客项目](https://github.com/DengZC2000/Go_Project)

## 使用Go-Zero为微服务框架

## 使用ETCD作为服务注册发现中间件

## 使用GORM为ORM(对象关系映射)框架操作Mysql

## 使用Redis中间件做缓存(Redis在wsl中启动，详见：[Windows使用Redis](https://www.blog.dzcs.online/windows%E5%AE%89%E8%A3%85%E5%B9%B6%E4%BD%BF%E7%94%A8redis))

### - 用户注销后JWT的存储

### - 存用户上线状态

### - 存Rpc调用的上线用户基本信息

### - 存用户禁言时间

## 利用Go实现网关

## 使用WebSocket实现聊天通信

## 基于Kafka消息队列实现日志操作

## 基于WebRtc实现视频通话，简单demo如下：

![QQ20240620205701-ezgif.com-video-to-gif-converter (1)](C:\Users\Lenovo\Desktop\FIM项目\QQ20240620205701-ezgif.com-video-to-gif-converter (1).gif)



## 项目大致架构图

![项目大致架构图](C:\Users\Lenovo\Desktop\FIM项目\项目大致架构图.png)

## 数据库表

![数据库表](C:\Users\Lenovo\Desktop\FIM项目\数据库表.png)



## 实现的接口

![实现的接口](C:\Users\Lenovo\Desktop\FIM项目\实现的接口.png)

##  除了网关，ETCD中注册的12个服务(包括5个RPC服务，7个API服务)

![除了网关，etcd里注册的12个服务](C:\Users\Lenovo\Desktop\FIM项目\除了网关，etcd里注册的12个服务.png)

## 使用到的container

![container](C:\Users\Lenovo\Desktop\FIM项目\container.png)

## Kafka的web操作界面——Kafka-map，消费日志：Action或Runtime，如用户auth服务，系统管理员的操作等。

![kafka-map](C:\Users\Lenovo\Desktop\FIM项目\kafka-map.png)
