FROM mysql:latest

#定义工作目录变量
ENV WORK_PATH /usr/local/init

#定义会被容器自动执行的目录
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d

#定义要执行的sql文件名
ENV DATABASE_INIT monitor.sql

# 定义要执行的shell文件名
ENV CREATE_DATA_SHELL init.sh

#创建工作文件目录
RUN mkdir -p $WORK_PATH

#把sql文件复制到工作目录下
COPY ./$DATABASE_INIT $WORK_PATH/

#把要执行的shell文件放到WORK_PATH目录下，容器会自动执行这个shell
COPY ./$CREATE_DATA_SHELL $AUTO_RUN_DIR/

#给执行文件增加可执行权限
RUN chmod a+x $AUTO_RUN_DIR/$CREATE_DATA_SHELL