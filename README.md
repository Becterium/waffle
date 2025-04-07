# waffle
-_-

基于微服务框架Kratos开发，Kratos封装官方的net/http包和GRPC以处理http和rpc的请求和响应，整个后端分为一个Geteway网关服务和两个rpc后端服务，分别对应app包下的waffle,user,media

#### 如何启动

前置条件：要装好MySQL，Redis，Consul，Elasticsearch，Kafka ，Minio ，Jaeger

使用Docker镜像的话可以看deploy文件夹下内容

配置好conf文件，可查看Kratos[https://go-kratos.dev/docs/]

###### 直接Golang编译运行

```sh
Dir1 = app/waffle/interface
cd $(Dir1)
make run

Dir2 = app/user/service
cd $(Dir2)
make run

Dir3 = app/media/service
cd $(Dir3)
make run
```

###### Docker-compose编译镜像运行

```sh
docker build --build-arg APP_RELATIVE_PATH=user/service -t 192.168.37.130:8009/library/waffle/user-kratos:latest .

docker build --build-arg APP_RELATIVE_PATH=waffle/interface -t 192.168.37.130:8009/library/waffle/waffle-kratos:latest .

docker build --build-arg APP_RELATIVE_PATH=media/service -t 192.168.37.130:8009/library/waffle/media-kratos:latest .

docker run 192.168.37.130:8009/library/waffle/waffle-kratos:latest conf文件夹地址
docker run 192.168.37.130:8009/library/waffle/user-kratos:latest conf文件夹地址
docker run 192.168.37.130:8009/library/waffle/media-kratos:latest conf文件夹地址
```

###### Jenkins + Kubernetes

```sh
有Jenkinsfile文件，直接创建Pipeline，JenkinsFile模式为SCM,路径为GitHub仓库路径
```

