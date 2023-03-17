FROM golang:latest

# 创建目录
RUN mkdir /HttpServer
# 工作目录/HttpServer
WORKDIR /HttpServer

# 拷贝当前目录下的所有文件到工作目录
COPY . .
#COPY main.go WORKDIR
#COPY readme.md WORKDIR
#COPY Dockerfile WORKDIR
#COPY go.mod WORKDIR

# 下载依赖
RUN go mod tidy
RUN go build -o MyServer ./

# 暴露端口
EXPOSE 80

# 执行命令
CMD ["./MyServer"]