# Easy_Douyin Project
本项目为模仿抖音APP的Easy抖音项目，使用RPC框架实现微服务。
## 项目模块介绍
三种服务模块分别是 用户服务、视频服务、评论服务
## 技术框架介绍：
- 语言：GO语言
- 底层存储: MySQL
- 缓存：Redis
- 服务注册: Etcd
- RPC框架: Kitex
- ORM框架: GORM
- HTTP框架: Gin
- 链路追踪: Jarger, opentracing
# 运行项目
需要提前准备：
- 安装docker
- 注：目前仅实现了用户服务，项目仅能运行用户服务，后续将逐步添加

1. 首先进入项目目录/easy_douyin，打开终端
2. 运行命令：docker-compose up
3. 新开终端，进入 /easy_douyin/cmd/user目录
    - 先运行命令：sh build.sh 
    - 再运行命令：sh ./output/bootstrap.sh
4. 运行 /easy_douyin/cmd/api/main.go 
5. 可以使用Postman进行测试了
