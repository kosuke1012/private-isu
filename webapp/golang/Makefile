include .env
all: app

app: *.go go.mod go.sum
	GOOS=linux GOARCH=amd64 go build -o app;

deploy-app:
	ssh private-isu-app "sudo rm -f /home/isucon/private_isu/webapp/golang/app" 
	scp app private-isu-app:/home/isucon/private_isu/webapp/golang
	ssh private-isu-app "sudo systemctl stop isu-go"
	ssh private-isu-app "sudo systemctl start isu-go"

download-slowlog:
	ssh private-isu-app "sudo cp /var/log/mysql/slow.log my_slow.log";	
	ssh private-isu-app "sudo chown isucon:isucon ~/my_slow.log";	
	scp private-isu-app:~/my_slow.log .
	ssh private-isu-app "sudo rm -f ~/my_slow.log";	

download-accesslog:
	ssh private-isu-app "sudo cp /var/log/nginx/access.log my_access.log";	
	ssh private-isu-app "sudo chown isucon:isucon ~/my_access.log"
	scp private-isu-app:~/my_access.log .
	ssh private-isu-app "sudo rm -f ~/my_access.log";	


bench-run:
	ssh private-isu-bench \
	"cd /home/isucon/private_isu.git/benchmarker; \
	./bin/benchmarker -u ./userdata -t http://192.168.0.42/"

mysql-rotate:
	ssh private-isu-app "sudo rm -f /var/log/mysql/slow.log"

mysql-restart:
	ssh private-isu-app "sudo systemctl restart mysql"

nginx-rotate:
	ssh private-isu-app "sudo rm -f /var/log/nginx/access.log"

nginx-restart:
	ssh private-isu-app "sudo systemctl restart nginx.service"

mysql-init: mysql-rotate mysql-restart

nginx-init: nginx-rotate nginx-restart

init: mysql-init nginx-init

show-pprof:
	$(eval latest := $(shell ls -rt ~/pprof/ | tail -n 1))
	go tool pprof -http=":8888" ~/pprof/$(latest);
pprof:
	go tool pprof -seconds 30 http://$(PRIVATE_ISU_APP)/debug/pprof/profile

alp:
	alp json --sort sum -r -m "/posts/[0-9]+,/@\w+,/image/[0-9]+" -o count,method,uri,min,avg,max,sum < my_access.log
ptq:
	pt-query-digest slow.log
