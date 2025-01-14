## Docker启动Consul
docker run -id --name=consul \
  -p 8300:8300 \
  -p 8301:8301 \
  -p 8302:8302 \
  -p 8500:8500 \
  -p 8600:8600 \
  -v /data/consul/data:/bitnami \
  192.168.37.130:8009/library/bitnami/consul:1.20.1-debian-12-r0

## Docker启动Redis
docker run -id --name=myredis \
  -p 6379:6379 \
  -v /data/redis/conf/redis.conf:/etc/redis/redis.conf \
  -v /data/redis/data:/data \
  192.168.37.130:8009/library/redis:7.4.0

## Docker启动Kafka
docker network create app-tier --driver bridge
## 开启kafka的服务实例
docker run -d --name kafka-server --hostname kafka-server \
    --network app-tier \
    --network-alias kafka-server \
    -p 9094:9094 \
    -e KAFKA_CFG_NODE_ID=0 \
    -e KAFKA_CFG_PROCESS_ROLES=controller,broker \
    -e KAFKA_CFG_LISTENERS=INTERNAL://:9092,CONTROLLER://:9093,EXTERNAL://:9094 \
    -e KAFKA_CFG_ADVERTISED_LISTENERS=INTERNAL://kafka-server:9092,EXTERNAL://192.168.37.134:9094 \
    -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT,INTERNAL:PLAINTEXT \
    -e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka-server:9093 \
    -e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER \
    -e KAFKA_CFG_INTER_BROKER_LISTENER_NAME=INTERNAL \
    -e KAFKA_CFG_METADATA_LOG_DIRS=/bitnami/kafka/data/meta \
    -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true \
    -v /data/kafka/data:/bitnami/kafka \
    192.168.37.130:8009/library/bitnami/kafka:3.8.0-debian-12-r5
## 开启kafka客户端实例
docker run -it --rm \
    --network app-tier \
        192.168.37.130:8009/library/bitnami/kafka:3.8.0-debian-12-r5
 kafka-topics.sh --list  --bootstrap-server kafka-server:9092
## 开启kafka-ui实例
docker run -d -p 9090:8080 --network app-tier \
    -e DYNAMIC_CONFIG_ENABLED=true 192.168.37.130:8009/library/provectuslabs/kafka-ui:v0.7.2
