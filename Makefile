SERVER := ./server/bin/smsgate-server
CONSOLE := ./console/bin/smsgate-console
REPORT := ./report/bin/smsgate-report

all:
	go build -o ${SERVER} ./server/main.go
	go build -o ${CONSOLE} ./console/main.go
	go build -o ${REPORT} ./report/main.go

test:
	go test -cover ./...

testv:
	go test -cover -v ./...

pack:
	cd console/ && tar -zcvf ../smsgate-console.`date +"%Y%m%d.%H%M%S"`.tar.gz bin/smsgate-console bin/start.sh bin/stop.sh static/ views/
	cd server/ && tar -zcvf ../smsgate-server.`date +"%Y%m%d.%H%M%S"`.tar.gz bin/smsgate-server bin/start.sh bin/stop.sh
