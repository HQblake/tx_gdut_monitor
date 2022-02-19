# 重命名镜像
# docker tag build_monitor-mysql:latest monitor-mysql:1.0
docker tag build_monitor-storage:latest monitor-storage:1.0
docker tag build_monitor-alert:latest monitor-alert:1.0
docker tag build_monitor-manage:latest monitor-manage:1.0
docker tag build_monitor-agent:latest monitor-agent:1.0
docker tag build_monitor-web:latest monitor-web:1.0

# 删除中间镜像
docker rmi -f $(docker images | grep "none" | awk '{print $3}')

# 删除旧镜像
# docker rmi -f build_monitor-mysql:latest
docker rmi -f build_monitor-storage:latest
docker rmi -f build_monitor-alert:latest
docker rmi -f build_monitor-manage:latest
docker rmi -f build_monitor-agent:latest
docker rmi -f build_monitor-web:latest