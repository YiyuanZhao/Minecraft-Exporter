FROM coreharbor.azurewaf.top/dockerhub/golang:1.23.2 AS builder  
  
# 设置工作目录  
WORKDIR /app  
  
# 复制 go.mod 和 go.sum 文件  
COPY go.mod go.sum ./  
  
# 下载依赖  
RUN go mod download  
  
# 复制源代码  
COPY . .  
  
# 构建应用  
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o minecraft_exporter .  

# 获取 mc-monitor 的可执行文件  
RUN git clone https://github.com/itzg/mc-monitor.git /mc-monitor && \  
    cd /mc-monitor && \  
    go build -buildvcs=false -o mc-monitor .
  
# 使用一个更小的基础镜像来运行应用  
FROM coreharbor.azurewaf.top/dockerhub/scratch
  
# 安装必要的包（如需要）  
# RUN apk --no-cache add ca-certificates  
  
# 复制构建好的二进制文件到新镜像  
COPY --from=builder /app/minecraft_exporter /usr/local/bin/  
COPY --from=builder /mc-monitor/mc-monitor /usr/local/bin/  
  
# 设置默认的命令  
ENTRYPOINT ["/usr/local/bin/minecraft_exporter"]  
  
# 暴露服务端口  
EXPOSE 8081  
