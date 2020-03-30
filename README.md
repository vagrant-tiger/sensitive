# sensitive
基于importcjj/sensitive+echo封装的敏感词过滤服务

测试替换接口：5w敏感词,请求800，并发800的平均响应时间为0.0538s,R/S为12875；

编译：
在mac系统中将go编译打包成可在linux上直接执行的方法
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 可执行文件目录/serve serve.go 

docker部署：

docker build -t sensitive .

docker run --name sensitive -d -p 8090:80 -v 可执行文件的目录:/go/bin/ -v 敏感词文件目录:/go/dic/  sensitive