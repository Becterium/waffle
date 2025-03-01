FROM 192.168.37.130:8009/library/golang:1.22.7 AS builder

ARG APP_RELATIVE_PATH

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

FROM 192.168.37.130:8009/library/debian:stable-20250113-slim-aptUpdated

# 确保 /etc/apt 目录存在
RUN mkdir -p /etc/apt

# 创建 sources.list 文件并写入内容
#RUN echo "deb http://mirrors.ustc.edu.cn/debian/ stable main" > /etc/apt/sources.list && \
#    echo "deb-src http://mirrors.ustc.edu.cn/debian/ stable main" >> /etc/apt/sources.list

ARG APP_RELATIVE_PATH

#RUN apt-get update && apt-get install -y --no-install-recommends \
#		ca-certificates  \
#        netbase \
#        && rm -rf /var/lib/apt/lists/ \
#        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]