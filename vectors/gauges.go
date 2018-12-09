package vectors

import (
	"github.com/prometheus/client_golang/prometheus"
)

func Init() {

	// Server
	GaugeVecs["serverMaxPlayers"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "server_maxplayers",
		Help:      "Count of maxplayers per server",
	}, []string{"maxplayers"})

	GaugeVecs["serverVersion"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "server_version",
		Help:      "Count of servers running a specific version",
	}, []string{"version"})

	GaugeVecs["serverOS"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "server_os",
		Help:      "Server operating system",
	}, []string{"os"})

	GaugeVecs["serverPort"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "server_port",
		Help:      "Server listening port",
	}, []string{"port"})

	GaugeVecs["secureServers"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "server_secure",
		Help:      "Count secure/insecure servers",
	}, []string{"secure"})

	// Map
	GaugeVecs["map"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "map",
		Help:      "General player/bot data to a single map",
	}, []string{"map", "playertype", "owner"})

	GaugeVecs["mapCount"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "map_server_count",
		Help:      "Map server count",
	}, []string{"map", "provider"})

	// Misc
	GaugeVecs["svTags"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "sv_tags",
		Help:      "sv_tags list",
	}, []string{"tag"}) // Unused!

	Gauges["masterServerStatus"] = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "tf2",
		Subsystem: "general",
		Name:      "masterserver_status",
		Help:      "Masterserver status",
	})

}
