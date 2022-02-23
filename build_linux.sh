# 创建网桥
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 monitor

# 运行 monitor-storage
docker build -t monitor-storage:1.0 --force-rm=true ./build/monitor-storage
docker run -d -p 8080:8080 -v /home/monitor/data/storage:/data/storage --name storage --net monitor monitor-storage:1.0

# 运行 monitor-alert
docker build -t monitor-alert:1.0 --force-rm=true ./build/monitor-alert
docker run -d -p 8081:8081 -v /home/monitor/data/alert:/data/alert --name alert --net monitor monitor-alert:1.0

# 运行 monitor-manage
docker build -t monitor-manage:1.0 --force-rm=true ./build/monitor-manage
docker run -d -p 8082:8082 -p 8083:8083 -v /home/monitor/data/manage:/data/manage --name manage --net monitor monitor-manage:1.0

# 运行 5台 monitor-agent
docker build -t monitor-agent:1.0 --force-rm=true ./build/monitor-agent
docker run -d -p 8085:8084 -v /home/monitor/data/agent/agent-01:/data/agent --name agent-01 --net monitor monitor-agent:1.0
docker run -d -p 8086:8084 -v /home/monitor/data/agent/agent-02:/data/agent --name agent-02 --net monitor monitor-agent:1.0
docker run -d -p 8087:8084 -v /home/monitor/data/agent/agent-03:/data/agent --name agent-03 --net monitor monitor-agent:1.0
docker run -d -p 8088:8084 -v /home/monitor/data/agent/agent-04:/data/agent --name agent-04 --net monitor monitor-agent:1.0
docker run -d -p 8089:8084 -v /home/monitor/data/agent/agent-05:/data/agent --name agent-05 --net monitor monitor-agent:1.0

# 运行ui模块
docker build -t monitor-web:1.0 --force-rm=true ./build/monitor-web
docker run -d -p 8084:80 -v /home/monitor/data/web:/usr/share/nginx/html/static -v /home/monitor/data/web/nginx:/etc/nginx/conf.d --name web --net monitor monitor-web:1.0

# 删除中间镜像
docker rmi -f $(docker images | grep "none" | awk '{print $3}')
