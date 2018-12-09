package c

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var (
	NetClient      *http.Client
	Registry       *prometheus.Registry
	HttpServer     *http.Server
	MetricsHandler http.Handler

	ListenAddr   string
	SteamApiKeys []string
	RequestCount int64
	OnGoing      bool
)
