# 阶段一：用中间镜像进行编译
FROM golang:1.17-alpine AS builder
# 配置编译环境
# RUN command
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 安装Git
RUN apk update
RUN apk upgrade
RUN apk add git

# 从Git中将项目源码加载至镜像中
# COPY ./monitor-agent /go/src/monitor-agent
RUN git clone -b develop https://gitee.com/zekeGitee_admin/tx_gdut_monitor.git /git
RUN mv /git/monitor-agent /go/src/

# 编译
# 工作目录
WORKDIR /go/src/monitor-agent
RUN go install ./...
RUN mv /go/bin/cmd /go/bin/agent

# # 阶段二：生成最终镜像
FROM alpine:latest

# 系统默认配置文件加载到指定目录中
COPY --from=builder /go/src/monitor-agent/configs /data/agent/configs

COPY --from=builder /go/bin/agent /bin/agent
# 也可以考虑拷贝一个默认的配置文件
VOLUME ["/data/agent"]
# 服务端口暴露
EXPOSE 8084
# 设置服务入口
CMD ["/bin/agent", "-config", "/data/agent/configs/config.yaml"]