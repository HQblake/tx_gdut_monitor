# dockerfile demo - zeke
# 启动编译环境
# 阶段一：用中间镜像进行编译
FROM golang:1.17-alpine AS builder
# 配置编译环境
# RUN command
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中，这一步后续考虑换成拉取git项目
# 假设当前目录已经在宿主机的tx_monitor目录下了,路径的明确很重要

COPY . /go/src/tx_gdut_monitor
# 编译
# 工作目录
WORKDIR /go/src/tx_gdut_monitor/monitor-manage
RUN go install ./monitor-manage/...

# 阶段二：生成最终镜像
FROM alpine:latest
COPY --from=builder /go/bin/monitor-manage /bin/monitor-manage
# 也可以考虑拷贝一个默认的配置文件
VOLUME ["/data/monitor-manage"]
# 服务端口暴露
# 管理服务这里一个是http端口，一个是rpc端口
EXPOSE 8080 8083
# 设置服务入口
ENTRYPOINT ["/bin/monitor-manage"]
CMD ["/bin/monitor-manage", "-config", "/data/monitor-manage/config.yml"]

# 有配置文件的话记得声明数据卷，并运行的时候指定数据卷位置,并提前放好配置文件
