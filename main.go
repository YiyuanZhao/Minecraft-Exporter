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
		// Execute the mc-monitor status command
		output, err := exec.Command("/mc-monitor", "status").Output()
		// output, err := exec.Command("echo", "localhost:25565 : version=1.21.1 online=0 max=20 motd='A Minecraft Server'").Output()
		if err != nil {
			fmt.Println("Error executing command:", err)
			continue
		}

		// Parse the command output
		parseOutput(string(output))

		time.Sleep(10 * time.Second) // Update every 10 seconds 
	}
}

func parseOutput(output string) {
	mu.Lock()
	defer mu.Unlock()

	// Regular expression to extract online and max
	re := regexp.MustCompile(`online=(\d+) max=(\d+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) == 3 {
		online, max := matches[1], matches[2]

		// Convert string to float64
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
	fmt.Println("Starting exporter on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
