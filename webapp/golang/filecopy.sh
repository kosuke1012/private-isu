#!/bin/bash
start_time=`date +%s`
while true
do
	for file in $( docker exec  webapp_app_1 ls /home/webapp/image/ ); do   
	echo ${file}
	docker cp webapp_app_1:/home/webapp/image/${file} .
	docker cp ${file} webapp_nginx_1:/public/image
	docker exec webapp_app_1 rm /home/webapp/image/${file}
	rm ${file}
	done
	now=`date +%s`
	diff=$((now-start_time))
	if [ ${diff} -gt 60 ]; then
	# 引数に指定された値を break コマンドに指定
	break
	fi
done