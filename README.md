# 斗牛短视频

## 简介

斗牛短视频是由**拿下七牛offer队**开发，后端是基于go-zero框架开发的一款微服务Web端短视频应用，本项目使用七牛云对象存储(
Kodo)对视频，图片进行存储，智能多媒体服务(Dora)
对视频进行处理，Redis作为缓存，Kafka进行异步处理和流量削峰，Etcd作为服务发现与注册中心。同时项目还具有良好的可观测性，使用Jaeger进行全链路追踪，Promethues+Grafana对服务进行性能监控，filebeat+go-stash+elasticsearch+kibana采集日志并且进行可视化呈现，使用docker-compose、Github
Action进行CI/CD流程，最后使用traefik作为网关对服务进行负载均衡和转发

## 演示视频地址

xxxx

## 快速开始

### 前端

x x x

### 后端

> 前提条件：安装好docker, docker-compose

1. 设置**环境变量**：`DOCKERHUB_IMAGE=liuxian123/douniu-`

2. 启动 docker-compose-env.yaml里面的相关中间件容器**（不需要可观测性可以省略）**

```
docker-compose up -d -f docker-compose-env.yaml
```

3. 配置GitHub Action的密钥**（不需要CI/CD流程可省略）**

![image-20231107172809895](https://raw.githubusercontent.com/liuxianloveqiqi/Xian-imagehost/main/image/202311071728126.png)

4. 启动docker-compose.yaml里面的各个服务的容器**（必须）**

```
docker-compose up -d
```

注：**腾讯云密钥**请自己自行配置，相关中间件也可以自己配置，请在各个服务的etc下完成配置

