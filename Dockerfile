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

  
# 使用一个更小的基础镜像来运行应用  
FROM coreharbor.azurewaf.top/dockerhub/itzg/mc-monitor:latest
  
# 安装必要的包（如需要）  
# RUN apk add --no-cache bash 
  
# 复制构建好的二进制文件到新镜像  
COPY --from=builder /app/minecraft_exporter /usr/local/bin/
  
# 设置默认的命令  
ENTRYPOINT ["/usr/local/bin/minecraft_exporter"]  
  
# 暴露服务端口  
EXPOSE 8081  
