
# 1创建Dockerfile  
在主目录下创建一个Dockerfile  
```dockerfile
# 创建目录
# RUN mkdir /app /config /logs 可以同时创建多个目录
RUN mkdir /HttpServer
# 工作目录/HttpServer
WORKDIR /HttpServer

# 拷贝当前目录下的所有文件到工作目录
COPY . .

# 下载依赖
RUN go mod tidy
RUN go build -o MyServer ./

# 暴露端口
EXPOSE 80

# 执行命令
CMD ["./MyServer"]
```
# 2 构建镜像
命令如下：
```shell
docker build -t my-server .
```

# 3 运行容器
命令如下：
```shell
 docker run -d -p 8088:80 -v E:\aida:/HttpServer/aida my-server
```
```shell
-d : 后台运行
-p 8088:80 : 映射端口 8088 外部端口 80 容器内部端口
-v E:\aida:/HttpServer/aida : 映射外部的存储路径 E:\aida 外部路径 /HttpServer/aida 容器内部路径
```
# 4 容器操作
```shell
docker ps #获取当前运行的容器

docker stop cae5fc81916e # 停止容器id为 cae5fc81916e 的容器
docker kill cae5fc81916e # 强行停止容器id为 cae5fc81916e 的容器

docker exec -it cae5fc81916e /bin/bash #进入容器内部的控制台 可以在里面kill 进程 等操作

docker rm cae5fc81916e # 删除容器id为cae5fc81916e的容器
docker rm container1 container2 container3 # 同时删除容器id为container1 container2 container3的容器
docker container prune #删除所有已停止的容器

docker rmi my-server #删除刚才创建的镜像
docker rmi image1 image2 image3 # 同时删除镜像 image1 image2 image3
docker image prune # 删除所有未被标记、未被使用的镜像

```