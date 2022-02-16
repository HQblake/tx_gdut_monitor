# 登录mysql并执行指定的sql文件
mysql -uroot -p$MYSQL_ROOT_PASSWORD <<EOF

# 选择 monitor数据库
use $MYSQL_DATABASE;
source $WORK_PATH/$DATABASE_INIT;
