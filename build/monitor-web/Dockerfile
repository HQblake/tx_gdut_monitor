# 启动编译环境
# 阶段一：用中间镜像进行编译
FROM alpine:latest AS builder
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories
# 配置编译环境
# RUN command
RUN apk update
RUN apk upgrade
RUN apk add git
# RUN apk add nodejs
# RUN apk add npm

RUN git clone -b develop https://gitee.com/zekeGitee_admin/tx_gdut_monitor.git /home/monitor
# WORKDIR /home/monitor/monitor-web/monitor-txui
# RUN cd /home/monitor/monitor-web/monitor-txui
# RUN npm install
# RUN npm run build

FROM nginx:latest
# 将dist文件中的内容复制到 /usr/share/nginx/html/ 这个目录下面
COPY --from=builder  /home/monitor/monitor-web/monitor-txui/dist/  /usr/share/nginx/html/
EXPOSE 80
