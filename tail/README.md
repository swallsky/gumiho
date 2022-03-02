### tail 钩子部分
```
# 编译生成mac、linux版
./build.sh
```

```
# 编译docker镜像
docker build -t ybluesky/gumiho-tail:v1 .

# 运行docker环境
docker run -d -p 8081:8081 --name gumiho-tail ybluesky/gumiho-tail:v1 /opt/tail server up

# 运行docker环境，自动重启
docker run -d -p 8081:8081 --restart=always --name gumiho-tail ybluesky/golang:v1 /opt/tail server up

# 查看docker运行日志
docker logs -f gumiho-tail

```
### 动态执行命令行
- conf.yaml router 为配置规则
- script/*.sh 为可执行脚本
- http://localhost:8081/script/:key :key为配置router的key
