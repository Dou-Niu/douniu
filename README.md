# 斗牛短视频

## 一、项目简介

> 项目地址：https://github.com/Dou-Niu/douniu

斗牛短视频是基于go-zero框架开发的一款微服务Web端短视频应用，本项目使用七牛云对象存储(Kodo)
对视频，图片进行存储，智能多媒体服务(Dora)
对视频进行处理，Redis作为缓存，Kafka进行异步处理和流量削峰，Etcd作为服务发现与注册中心。同时项目还具有良好的可观测性，使用Jaeger进行全链路追踪，Promethues+Grafana对服务进行性能监控，filebeat+go-stash+elasticsearch+kibana采集日志并且进行可视化呈现，使用docker-compose、Github
Action进行CI/CD流程，最后使用traefik作为网关对服务进行负载均衡和转发

## 二、团队分工

| 刘显-后端 | user、video服务的开发，CI/CD流程，OpenTelemetry |
|-------|---------------------------------------|
|       |                                       |
|       |                                       |

## 三、项目实现

### 3.1 项目演示

### 3.2 需求分析

#### 3.2.1 用户需求分析

1. 手机号+验证码一键注册/登录，手机号+密码登录
2. 能查看自己和其他用户的基本信息
3. 能修改自己的昵称、头像、背景图
4. 忘记密码了能重置，登录后能修改自己的密码

#### 3.2.2 视频需求分析

1. 视频feed流响应速度要快
2. 用户能自己上传视频和删除视频
3. 能根据热度或者时间对视频排序
4. 关键词搜索视频
5. 对视频进行分区展示

#### 3.2.3 点赞需求分析

1. 能对视频点赞/取消赞
2. 点赞/取消赞不能卡顿

#### 3.2.4 评论需求分析

1. 用户能对视频评论和嵌套评论其他用户的评论
2. 评论的内容要做管控

#### 3.2.5 关注需求分析

1. 用户能关注/取关其他用户
2. 用户能查看自己的关注列表

#### 3.2.6 聊天需求分析

1. 用户能与好友（互相关注）进行简易聊天

### 3.3 技术选型

| **功能** | **实现**                                 |
|--------|----------------------------------------|
| http框架 | go-zero                                |
| rpc框架  | go-zero                                |
| 数据库    | MySQL，Redis                            |
| 搜索引擎   | elasticsearch                          |
| 注册中心   | etcd                                   |
| 对象存储   | 七牛云Kodo                                |
| 视频处理   | 七牛云Dora                                |
| 链路追踪   | jaeger                                 |
| 服务监控   | prometheus，grafana                     |
| 消息队列   | kafka                                  |
| 日志搜集   | filebeat,go-stash,elasticsearch,kibana |
| 网关     | traefik                                |
| 部署     | Docker,docer-compose                   |
| CI/CD  | Github Action                          |

### 3.4 架构设计

#### 3.4.1 系统架构

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558152)

#### 3.4.2 业务架构

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558213.(null))

#### 3.4.3 OpenTelemetry

##### 3.4.3.1 **链路追踪**

使用`jaeger`作为链路追踪，`jaeger`专门设计用于跟踪分布式应用程序中的请求，可以帮助理解多个服务之间的交互，并可视化请求在不同服务之间的流动

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558461.(null))

##### 3.4.3.2**日志搜集**

日志是故障排查中非常重要的数据来源。当系统出现故障时，管理员可以通过查看日志文件来定位问题所在。也可以帮助管理员对系统进行性能调优。通过分析日志数据，管理员可以了解系统的瓶颈所在。

![img](https://nzc62qt7x1.feishu.cn/space/api/box/stream/download/asynccode/?code=ZDZjNDM3NjE1Y2JkYzczY2E2YmJkMDYwZDEzN2RiMTlfeFc3ZEgzWlhjR05EZDJCdjZOdUh1Y2pvWVNCZHF1T3ZfVG9rZW46V2VFcmJXaHdzb1FCWkZ4RXFETWNWVjdLblliXzE2OTkyNTc1MTU6MTY5OTI2MTExNV9WNA)

`filebeat`收集业务日志，然后将日志输出到`kafka`中作为缓冲，`go-stash`获取`kafka`
中日志根据配置过滤字段，然后将过滤后的字段输出到`elasticsearch`中，最后由`kibana`负责呈现日志

![img](https://nzc62qt7x1.feishu.cn/space/api/box/stream/download/asynccode/?code=MTg0ZGVhZWI0MDRjYmI0ZDQ4NWY5YTdlODU0ODhjZTBfTEdHamxDVjdMYzd1YUhvQ1JqSjl1TGJJeXBYbHN0Y3VfVG9rZW46RFBVT2JGWjVsb05SSTB4M2RpSWNTMmZ4bmJlXzE2OTkyNTc1MTU6MTY5OTI2MTExNV9WNA)

##### 3.4.3.3 服务监控

Prometheus和Grafana的结合提供了实时多维度性能监控，使用简单且高度可扩展的工具，Prometheus收集并存储数据，支持警报和通知，而Grafana提供强大的数据可视化功能，这一组合使开发人员能够迅速发现问题、优化性能，同时保持系统的稳定性和可靠性。

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558453.(null))

`grafana`进行可视化呈现

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558734.(null))

#### 3.4.5 数据库设计

垂直分库，根据微服务的设计将不同的服务拆分到多个库中，创建了合适的索引，提高整体系统的安全性和性能

![img](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311061558034.(null))

### 3.5 项目亮点

#### 3.5.1 接入腾讯云SMS短信服务

支持了用户手机号一键注册/登录，提升了用户体验感

#### 3.5.2 接入七牛云Kodo和Doro对视频进行处理

将用户上传的视频存储到Kodo，再利用Doro对视频进行加水印和截帧等相关操作

#### 3.5.3 视频的最新/热度排序

利用redis的Zset结构，其中Score存储时间戳/热度值，就可以做到相应的排序规则

#### 3.5.4 评论嵌套

实现了用户评论嵌套的功能

#### 3.5.5 视频关键词搜索

将视频信息存储到ElasticSearch中，为用户提供一个搜索界面，以便他们可以执行关键字搜索，同时支持排序和分页

### 3.6 性能优化

#### 3.6.1 缓存

使用go-zero设计的缓存系统，使用redis对mysql中的信息做了缓存，减少了对后端存储的访问，减轻了数据库和服务器的压力

#### 3.6.2 并发操作

对于多个可以并行请求处理的逻辑，比如请求rpc，写入mysql、redis、es。我们均采用了协程并发调用，使串行变为并行，极大缩短了响应时间。

因此我们使用`go-zero`提供的并发组件[mr](https://github.com/zeromicro/go-zero/tree/master/core/mr)进行并发操作，将响应时间大大降低

#### 3.6.3 分页

对视频feed流、关注列表、点赞/收藏视频列表、评论均做了分页处理，显著降低页面加载时间

#### 3.6.4  使用Kafka进行异步处理和流量削峰

当用户上传视频到Kodo进行处理后，前端返回视频的相关信息给后端，此时使用kafka异步存储视频信息的方式，避免用户继续等待视频信息存入db的时间

在高并发场景下，当大量用户同时进行点赞操作时，Redis作为一个常用的内存数据库可能会面临性能瓶颈和承受能力限制。为了解决这个问题，我们采取了一种策略，首先将点赞流量写入Kafka中，从而实现流量削峰和分流的效果。

#### 3.6.5 Kafka的消息聚集和并发消费

##### 3.6.5.1 批量消息聚合提升`kafka`性能

![img](https://nzc62qt7x1.feishu.cn/space/api/box/stream/download/asynccode/?code=ZTI3MzE4YjI2YzM0ZTRjZGMzZDg4YTZmZjVmZmIwOWFfN0MwWGtWUEt5Q1lNZmp1dzFOSTUwOHRUYXpZdDNzTXRfVG9rZW46UkVHZ2JoRm43b3NxRWd4Nkx0OGNXekdNbkFjXzE2OTkyNTc1MTU6MTY5OTI2MTExNV9WNA)

之前每向`kafka`发送一条消息就会产生一次网络 IO 和一次磁盘 IO，做消息聚合后，比如聚合 100 条消息后再发送给` Kafka`，这个时候
100 条消息才会产生一次网络 IO 和磁盘 IO，这样大大提高 `Kafka`
的吞吐和性能。并且有聚合时间兜底，就算消息数量达不到聚合要求，超过聚合最大时间也会聚合当前所有消息发送给`Kafka`

##### 3.6.5.2 goruntie并发消费

针对于没有要求保障消息顺序的业务，在从kafka拉取消息后，会放入本地的channel中，然后会创建多协程并发消费，提高性能

![img](https://nzc62qt7x1.feishu.cn/space/api/box/stream/download/asynccode/?code=Zjk4Y2YwMzcwZWQzMzliNTQyZWU1MjRmZTA3MWU5NGNfZUhLZTNYMWxEU21VS1FCT2d3RzJUaEowQTBXT1VrN3FfVG9rZW46RjQ1SmI3ZkR3b2dLM3R4dExyamNDVHN1bmlnXzE2OTkyNTc1MTU6MTY5OTI2MTExNV9WNA)

### 3.7 CI/CD

使用Github Action作为CI/CD，操作流程如下

1. 使用github action运行test，进行自动化测试
2. 根据dockerfile构建镜像推送到dockerhub
3. 使用ssh登录远程服务器，执行脚本重新部署容器

```Go
jobs:
  # user-api
  build-and-push-user-api:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout the repository
        uses: actions/checkout@v2

      - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v1

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v1

      - name: Create and push user-api Docker image
        uses: docker/build-push-action@v2
        with:
          context: .
          file: ./server/user/api/Dockerfile
          push: true
          tags: ${{ secrets.DOCKERHUB_IMAGE }}user-api:latest
          platforms: linux/amd64,linux/arm64  # 构建多个架构的镜像
      - name: executing remote ssh commands using password
        uses: appleboy/ssh-action@v0.1.10
        with:
          host: ${{ secrets.HOST }}
          username: ${{ secrets.USERNAME }}
          password: ${{ secrets.PASSWORD }}
          port: ${{ secrets.PORT }}
          script: |
            cd /home/douniu
            docker-compose.yaml stop user-api
            docker-compose.yaml rm -f user-api
            docker image rm ${{ secrets.DOCKERHUB_IMAGE }}user-api:latest
            docker-compose -f docker-compose.yaml up -d user-api
```
