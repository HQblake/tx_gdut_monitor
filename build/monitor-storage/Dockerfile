# 阶段一：用中间镜像进行编译
FROM golang:1.17-alpine AS builder
# 配置编译环境
# RUN command
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 从 gitee 上拉取项目源码
# 目前先从develop分支下拉取
git clone -b develop https://gitee.com/zekeGitee_admin/tx_gdut_monitor.git

COPY . /go/src/tx_gdut_monitor
# 编译
# 工作目录
WORKDIR /go/src/tx_gdut_monitor/monitor-storage
RUN go install ./monitor-storage/...

# 阶段二：生成最终镜像
FROM alpine:latest
COPY --from=builder /go/bin/monitor-storage /bin/monitor-storage
# 也可以考虑拷贝一个默认的配置文件
VOLUME ["/data/monitor-storage"]
# 服务端口暴露
# 管理服务这里一个是http端口，一个是rpc端口
EXPOSE 8080 8083
# 设置服务入口
ENTRYPOINT ["/bin/monitor-manage"]
CMD ["/bin/monitor-manage", "-config", "/data/monitor-manage/config.yml"]