### tail 钩子部分
```
# 编译生成mac、linux版
./build.sh
```

```
# 编译docker镜像
docker build -t ybluesky/golang:v1 .

# 运行docker环境
docker run -d -p 8081:8081 --name gumiho-tail ybluesky/golang:v1 /opt/tail server up

# 查看docker运行日志
docker logs -f gumiho-tail

```