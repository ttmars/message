kill -9 `ps -ef|grep message | grep -v grep|awk '{print $2}'`
cd /root/goProject/message
git pull origin master
go build -o message main.go
nohup ./message >./log 2>&1 &
