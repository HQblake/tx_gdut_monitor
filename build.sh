# 运行 monitor-storage
docker build -t monitor-storage:1.0 --force-rm=true ./build/monitor-storage
docker run -d -p 8080:8080 -v /mnt/f/Docker/monitor/data/storage:/data --name storage monitor-storage:1.0

# 运行 monitor-alert
docker build -t monitor-alert:1.0 --force-rm=true ./build/monitor-alert
docker run -d -p 8081:8081 -v /mnt/f/Docker/monitor/data/alert:/data --name alert monitor-alert:1.0

# 运行 monitor-manage
docker build -t monitor-manage:1.0 --force-rm=true ./build/monitor-manage
docker run -d -p 8082:8082 -p 8083:8083 -v /mnt/f/Docker/monitor/data/manage:/data --name manage monitor-manage:1.0

# 运行 monitor-agent
docker build -t monitor-agent:1.0 --force-rm=true ./build/monitor-agent
docker run -d -p 8084:8084 -v /mnt/f/Docker/monitor/data/agent:/data --name agent monitor-agent:1.0
