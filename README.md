# waffle
Kratos ,  just demo now  -_-

这Java开发的中间件镜像真占内存，难绷 -_-



###### 用二进制启动命令

```sh
configs是配置项文件夹

./bin/server -conf configs
```

###### 用DockerFile构建镜像并启动镜像

```sh
示例构建user服务的镜像
cd */waffle
docker build --build-arg APP_RELATIVE_PATH=user/service -t user-kartos .

查看镜像
[root@localhost service]# docker images | grep user-kartos
REPOSITORY         TAG             IMAGE ID       CREATED        SIZE

user-kartos        latest          320cc304a853   7 minutes ago  111MB

启动镜像(由于使用了MySQL,Redis和Consul,需要启动这三个服务，并在configs配置好)
docker run -p 8000:8000 -p 9000:9000 -v /mnt/hgfs/DevopsShareDir/waffle/app/user/service/configs:/data/conf user-kartos:latest
```

