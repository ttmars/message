kill -9 `ps -ef|grep message|grep -v grep|awk '{print $2}'`
go build -o message ./main.go
nohup ./message >./log 2>&1 &
