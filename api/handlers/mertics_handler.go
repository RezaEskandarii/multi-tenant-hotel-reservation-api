package handlers

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"reservation-api/internal/global_variables"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricHandler struct {
	Config *global_variables.Config
}

func (handler *MetricHandler) Register(config *global_variables.Config) {
	handler.Config = config
	go handler.listenToMetrics()
}
func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func (handler *MetricHandler) listenToMetrics() {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", handler.Config.Application.MetricEndPointPort), nil)
}
