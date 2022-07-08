all: app

app: *.go go.mod go.sum
	go build -o app

bench-run:
	docker exec -d webapp_mysql_1 rm /var/lib/mysql/my_mysql_slow_query.log
	docker exec -d webapp_nginx_1 rm /var/log/nginx/access.log
	docker restart webapp_mysql_1
	docker restart webapp_nginx_1
	sleep 10
	go tool pprof -seconds 15 http://localhost/debug/pprof/profile &
	./filecopy.sh &
	docker run --network host -i private-isu-benchmarker \
	 /opt/go/bin/benchmarker -t http://host.docker.internal -u /opt/go/userdata
	docker cp webapp_mysql_1:/var/lib/mysql/my_mysql_slow_query.log .
	docker cp webapp_nginx_1:/var/log/nginx/access.log .

pprof:
	$(eval latest := $(shell ls -rt ~/pprof/ | tail -n 1))
	go tool pprof -http=":8888" ~/pprof/$(latest);

alp:
	alp json --sort sum -r -m "/posts/[0-9]+,/@\w+,/image/[0-9]+" -o count,method,uri,min,avg,max,sum < access.log