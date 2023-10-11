package core

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"log"
	"log/slog"
	"net/http"
)

var Reg = prometheus.NewRegistry()
var Met = NewMetrics(Reg)

type Metrics struct {
	VisitCount prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		VisitCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "visit_count",
			Help: "访问量",
		}),
	}
	reg.MustRegister(m.VisitCount)
	return m
}

func ExportMetric() {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic occurred", "error", err, "server", "ExportMetric")
		}
	}()

	go func() {
		c := cron.New()
		c.AddFunc("@daily", func() { Met.VisitCount.Set(0) })
		c.Start()
	}()

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	slog.Info(fmt.Sprintf("metrics expose at http://localhost:%v/metrics", 9000))
	http.Handle("/metrics", promhttp.HandlerFor(Reg, promhttp.HandlerOpts{Registry: Reg}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 9000), nil))
}
