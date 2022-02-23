# 启动编译环境
# 阶段一：用中间镜像进行编译
FROM node:14-alpine AS builder
# RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories

# 配置编译环境
RUN apk update
RUN apk upgrade
RUN apk add git

# 拉取并打包项目源码
RUN git clone -b develop https://gitee.com/zekeGitee_admin/tx_gdut_monitor.git /home/monitor --depth 1
WORKDIR /home/monitor/monitor-web/monitor-txui
RUN cd /home/monitor/monitor-web/monitor-txui
RUN npm install node-sass --sass_binary_site=https://npm.taobao.org/mirrors/node-sass/
RUN npm install --save dayjs
RUN npm run build

FROM nginx:latest
# 将dist文件中的内容复制到 /usr/share/nginx/html/ 这个目录下面
COPY --from=builder  /home/monitor/monitor-web/monitor-txui/dist/  /usr/share/nginx/html/
VOLUME ["/usr/share/nginx/html/static"]
VOLUME ["/etc/nginx/conf.d"]
EXPOSE 80
