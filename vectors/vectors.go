package vectors

import (
	"git.lumen.sh/xNevo/tf2-metrics/c"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GaugeVecs map[string]*prometheus.GaugeVec
	Gauges    map[string]prometheus.Gauge
)

func init() {
	GaugeVecs = make(map[string]*prometheus.GaugeVec, 30)
	Gauges = make(map[string]prometheus.Gauge, 15)
}

func Load() {
	for _, gaugeVec := range GaugeVecs {
		c.Registry.MustRegister(gaugeVec)
	}

	for _, gauge := range Gauges {
		c.Registry.MustRegister(gauge)
	}
}

func Reset() {
	for _, gaugeVec := range GaugeVecs {
		gaugeVec.Reset()
	}

	for _, gauge := range Gauges {
		gauge.Set(0)
	}
}
