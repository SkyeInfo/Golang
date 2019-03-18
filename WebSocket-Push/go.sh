#!/bin/sh
# /etc/rc.d/init.d/go.sh
# Runs the Gogs
# chkconfig:   - 85 15 
. /etc/init.d/functions

cur_dir=$(cd `dirname $0`; pwd) #获取当前路径
NAME=pintuan-go   #运行文件名
GOGS_HOME="$cur_dir" #文件路径
GOGS_PATH=${GOGS_HOME}/$NAME #文件地址
SERVICENAME="pintuan-go"
LOCKFILE="$cur_dir/lockfile"
LOGPATH=${GOGS_HOME}/lockfile #日志路径
LOGFILE=${LOGPATH}/gogs.log #完成的日志地址
RETVAL=0

# 没有执行文件退出
#[ -x ${GOGS_PATH} ] || exit 0
# 如果不创建日志目录退出。
#[ -x ${LOGPATH} ] || exit 0
DAEMON_OPTS="--check $NAME"

start() {
  cd ${GOGS_HOME}
  now_time=`date +'%s'`
  echo -n "${now_time} - Starting ${SERVICENAME}:"
  daemon $DAEMON_OPTS "${GOGS_PATH} > ${LOGFILE} 2>&1 &"  # 执行文件  记录日志
  RETVAL=$? #如果是0执行成功   否则执行失败
  echo
  [ $RETVAL = 0 ] && touch ${LOCKFILE} #touch 当前已经存在文件，文件的访问、修改时间进行改变不改变文件的内容
  return $RETVAL
}

stop() {
  	cd ${GOGS_HOME}
  	now_time=`date +'%s'`
        echo -n " ${now_time} - Shutting down ${SERVICENAME}: "
        pid=$(ps -ef | grep "pintuan-go" | grep -v grep | awk '{print $2}')
        echo $pid
        kill $pid
        RETVAL=$?
        echo
        [ $RETVAL = 0 ] && rm -f "${LOGPATH}/gogs.log"
}

case "$1" in
    start)
        #status ${NAME} > /dev/null 2>&1 && exit 0
        status ${NAME}
	start
        ;;
    stop)
        stop
        ;;
    status)
        status ${NAME}
        ;;
    restart)
        stop
        start
        ;;
    reload)
        stop
        start
        ;;
    *)
        echo "Usage: ${NAME} {start|stop|status|restart}"
        exit 1
        ;;
esac
exit $RETVAL