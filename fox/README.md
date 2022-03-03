### fox gumiho系统平台
```
# 编译生成mac、linux版
./build.sh
```

```
# 编译docker镜像
docker build -t ybluesky/gumiho-fox:v1 .

# 运行docker环境
docker run -d -p 8080:8080 --name gumiho-fox ybluesky/gumiho-fox:v1 /opt/fox server up

# 运行docker环境，自动重启
docker run -d -p 8081:8081 --restart=always --name gumiho-fox ybluesky/gumiho-fox:v1 /opt/fox server up

# 查看docker运行日志
docker logs -f gumiho-fox

# 推送docker镜像到docker hub
docker push ybluesky/gumiho-fox:v1

```
### 动态执行命令行
- conf.yaml 为配置规则
