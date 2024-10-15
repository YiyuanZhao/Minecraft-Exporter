FROM coreharbor.azurewaf.top/dockerhub/golang:1.23.2 AS builder  
  
#  Set working directory 
WORKDIR /app  
  
# Copy go.mod and go.sum files 
COPY go.mod go.sum ./  
  
# Download dependencies 
RUN go mod download  
  
# Copy source code 
COPY . .  
  
# Build application 
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o minecraft_exporter .  

  
# Use a smaller base image to run the application
FROM coreharbor.azurewaf.top/dockerhub/itzg/mc-monitor:latest
  
# Install necessary packages (if needed)  
# RUN apk add --no-cache bash 
  
# Copy the built binary to the new image
COPY --from=builder /app/minecraft_exporter /usr/local/bin/
  
# Set default command  
ENTRYPOINT ["/usr/local/bin/minecraft_exporter"]  
  
# Expose service port  
EXPOSE 8082
