# Minecraft Exporter  
   
## Overview  
   
Minecraft Exporter is a Go application that monitors the status of a Minecraft server and exposes metrics about the number of online players and maximum players allowed. These metrics can be scraped by Prometheus for monitoring and alerting purposes.  
   
## Features  
   
- Monitors the current number of online players on a Minecraft server.  
- Displays the maximum number of players allowed on the server.  
- Exposes metrics in a format compatible with Prometheus.  
- Built using Docker for easy deployment.  
   
## Requirements  
   
- Go 1.23.2 or later  
- Docker  
- A running instance of the Minecraft server.  
   
## Getting Started  
   
### Clone the Repository  
   
```bash  
git clone https://github.com/yourusername/minecraft-exporter.git  
cd minecraft-exporter  
```  
   
### Build and Run with Docker  
   
You can build and run the application using Docker. The provided `Dockerfile` will create a container image for the Minecraft Exporter.  
   
```bash  
# Build the Docker image  
docker build -t minecraft_exporter .  
   
# Run the Docker container  
docker run -p 8082:8082 minecraft_exporter  
```  
   
### Access Metrics  
   
Once the application is running, you can access the metrics at:  
   
```  
http://localhost:8082/metrics  
```  
   
## mc-monitor Command  
   
The `mc-monitor` command is used to retrieve the status of the Minecraft server. This tool is implemented in the [itzg/mc-monitor](https://github.com/itzg/mc-monitor.git) GitHub repository. 
   
### Dockerfile Explanation  
   
The Dockerfile is structured in two stages:  
   
1. **Builder Stage**:  
    - Uses `golang:1.23.2` as the base image.  
    - Sets the working directory.  
    - Copies the Go module files and downloads dependencies.  
    - Copies the source code and builds the application.  
   
2. **Final Stage**:  
    - Uses a smaller base image (`itzg/mc-monitor`).  
    - Copies the built binary from the builder stage.  
    - Sets the default command to run the exporter.  
    - Exposes port `8082`.  
   
### Jenkins CI/CD Pipeline  
   
A Jenkinsfile is provided for automating the build and deployment process using Kaniko. It defines a pipeline that builds the Docker image and pushes it to a specified destination.   
  
### Usage  
   
- The application runs a loop that executes the `/mc-monitor status` command every 10 seconds to fetch the current status of the Minecraft server.  
- It uses regex to parse the output and update the Prometheus metrics accordingly.  
   
### Contributing  
   
Feel free to contribute by submitting issues and pull requests. Your contributions will help improve the project!  
   
### License  
   
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.  
   
---  
   
For any questions or issues, please open an issue in the GitHub repository. Happy gaming!