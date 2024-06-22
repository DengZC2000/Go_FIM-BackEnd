# FIM即时通讯BackEnd



##  使用Go语言为后端，使用Docker运行ETCD及Kafka、Redis，并使用Dockerfile完成本地部署。

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

## file通过uuid路径映射统一预览，简单demo如下：
![uid-ezgif com-video-to-gif-converter](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/21ff68f8-5f7d-4b55-b82b-98b165ed4673)


## 基于WebRtc实现视频通话，简单demo如下：

![QQ20240620205701-ezgif com-video-to-gif-converter (1)](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/a7cfb5fa-d2b3-4fb1-b502-bb54f046bb63)


## 项目大致架构图

![项目大致架构图](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/c39693e3-800e-4d31-b397-8869f6cead3d)


## 数据库表
![数据库表](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/d554e215-74c6-4a6b-b551-dce9378500f3)


## 实现的接口

![实现的接口](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/f952ff4d-a607-4520-8adb-f237875871b5)


##  除了网关，ETCD中注册的12个服务(5个总的RPC服务，7个总的API服务)

![除了网关，etcd里注册的12个服务](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/faa558b3-6e39-4c1e-a625-f2704b65a9b5)


## 使用到的container

![container](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/fe2bd292-3598-43c7-87f5-e7e8ace000e2)


## Kafka的web操作界面——Kafka-map，消费日志：Action或Runtime，如用户auth服务，系统管理员的操作等。

![kafka-map](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/f51e2fb2-d931-4eb3-a42a-060c12f72803)

## 本地部署相关：
### - 1、etcd、redis、kafka相关的全部用docker运行，开发项目中对应的ip改为内网地址，我这里机器最后一次重启，查看内网ip是172.29.224.1。特别是redis，发开环境中用的wsl中的redis，然后ip就一直写的是127.0.0.1，但是踩坑不行，得改。
### - 2、gateway对应yaml改为0.0.0.0:port，然后测试接口时统一使用127.0.0.1:port/xx/xxx/xx就可以了。
### 本地部署后etcd中注册的12个服务：
![本地部署后etcd中注册的12个服务](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/b52852dd-4c3b-4023-8632-13c4b17bf8db)

### 打包好的镜像：
![所有镜像](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/e5a8b7cb-797d-478c-ad69-49d4ab28c236)

### 启动的所有容器如下：
![启动所有容器](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/6331448e-2f2f-46e9-b1ad-4ea42967ae87)

### 本地部署完成后，就可以一摸一样像开发时一样使用Apifox进行测试

### 可以curl一个不需要JwtToken（也就是不需要先经过网关请求auth的服务，直接curl IP：端口号就行）的Get请求接口看看，这里请求第三方登录信息配置接口，可以看到没问题的：
![curl不需要jwt的get请求](https://github.com/DengZC2000/Go_FIM-BackEnd/assets/153356545/1b279508-7b3d-4160-b87a-078d50882e91)
