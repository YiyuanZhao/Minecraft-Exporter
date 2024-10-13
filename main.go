package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	onlinePlayers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "minecraft_online_players",
		Help: "Number of players currently online on the Minecraft server",
	})
	maxPlayers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "minecraft_max_players",
		Help: "Maximum number of players allowed on the Minecraft server",
	})
	mu sync.Mutex
)

func init() {
	// Register the gauges with Prometheus
	prometheus.MustRegister(onlinePlayers)
	prometheus.MustRegister(maxPlayers)
}

func updateMetrics() {
	for {
		// 执行 mc-monitor status 命令
		output, err := exec.Command("/usr/local/bin/mc-monitor", "status").Output()
		// output, err := exec.Command("echo", "localhost:25565 : version=1.21.1 online=0 max=20 motd='A Minecraft Server'").Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}

		// 解析命令输出
		parseOutput(string(output))

		time.Sleep(10 * time.Second) // 每10秒更新一次
	}
}

func parseOutput(output string) {
	mu.Lock()
	defer mu.Unlock()

	// 正则表达式提取 online 和 max
	re := regexp.MustCompile(`online=(\d+) max=(\d+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) == 3 {
		online, max := matches[1], matches[2]

		// 将字符串转换为 float64
		onlinePlayers.Set(float64(atoi(online)))
		maxPlayers.Set(float64(atoi(max)))
	} else {
		fmt.Println("No matches found in output")
	}
}

func atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

func main() {
	go updateMetrics()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting exporter on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
