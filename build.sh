# 运行 monitor-storage
docker build -t monitor-storage:1.0 -f ./build/monitor-storage/Dockerfile --rm .
docker run -d -p 8080:8080 -v /mnt/f/Docker/monitor/data:/data --name storage monitor-storage:1.0

docker build -t monitor-alert:1.0 -f ./build/monitor-alert/Dockerfile --rm .
docker run -d -p 8081:8081 -v /mnt/f/Docker/monitor/data:/data --name alert monitor-alert:1.0

docker build -t monitor-manage:1.0 -f ./build/monitor-manage/Dockerfile --rm .
docker build -t monitor-agent:1.0 -f ./build/monitor-agent/Dockerfile --rm .
