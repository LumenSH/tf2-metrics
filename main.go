package main

import (
	"flag"
	"git.lumen.sh/xNevo/tf2-metrics/c"
	"git.lumen.sh/xNevo/tf2-metrics/tasks"
	"git.lumen.sh/xNevo/tf2-metrics/vectors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	c.SteamApiKeys = []string{os.Getenv("STEAM_KEY1"), os.Getenv("STEAM_KEY2")}

	flag.StringVar(&c.ListenAddr, "listen", "0.0.0.0:5000", "")
	flag.Parse()

	c.NetClient = &http.Client{
		Timeout: time.Second * 30,
	}
	log.Println("Net client created")

	c.Registry = prometheus.NewRegistry()
	log.Println("Registry created")

	vectors.Init()
	log.Println("Metric templates loaded")

	c.MetricsHandler = promhttp.InstrumentMetricHandler(
		c.Registry, promhttp.HandlerFor(c.Registry, promhttp.HandlerOpts{}),
	)

	log.Println("Add metrics to registry")
	vectors.Load()

	c.HttpServer = &http.Server{
		ReadTimeout:  45 * time.Second,
		WriteTimeout: 45 * time.Second,
	}
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s - %s", r.RemoteAddr, r.RequestURI)
		c.RequestCount++
		tasks.Run()
		c.MetricsHandler.ServeHTTP(w, r)
	})
}

func main() {
	log.Fatal(http.ListenAndServe(c.ListenAddr, nil))
}
