docker run -id --name=consul \
  -p 8300:8300 \
  -p 8301:8301 \
  -p 8302:8302 \
  -p 8500:8500 \
  -p 8600:8600 \
  -v /data/consul/data:/bitnami \
  192.168.37.130:8009/library/bitnami/consul:1.20.1-debian-12-r0

docker run -id --name=myredis \
  -p 6379:6379 \
  -v /data/redis/conf/redis.conf:/etc/redis/redis.conf \
  -v /data/redis/data:/data \
  192.168.37.130:8009/library/redis:7.4.0