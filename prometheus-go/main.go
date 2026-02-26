package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// creating the struct for devices
type Device struct {
	Id       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

// making the devices slice globally avaiable
var dvs []Device
var version string

// Creating our OWN METRIC
type metrics struct {
	devices prometheus.Gauge
	info    *prometheus.GaugeVec
}

func init() {
	// Initialize some mock devices
	dvs = []Device{
		{Id: 1, Mac: "5F-33-CC-1F-43-82", Firmware: "1.0.6"},
		{Id: 2, Mac: "6F-38-CD-8F-13-12", Firmware: "1.0.0"},
	}
	version = "1.0.0"
}

// for defining metrics
func newMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "devices_total",
			Help: "Total number of devices",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "device_info",
			Help: "Information about devices",
		}, []string{"version"}),
	}
	reg.MustRegister(m.devices)
	reg.MustRegister(m.info)
	return m
}

func main() {

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	m := newMetrics(reg)
	m.devices.Set(float64(len(dvs)))
	m.info.With(prometheus.Labels{"version": version}).Set(1)

	// multiple servers using go routines

	dMux := http.NewServeMux()
	dMux.HandleFunc("/devices", getDevices) // for serving the main content

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	select {}

	// Prometheus metrics at the /metrics endpoint.
	//
	// The metrics come from the default Prometheus registry,
	// which automatically includes:
	//   - Go runtime metrics (goroutines, GC, memory stats)
	//   - Process metrics (CPU, memory usage, file descriptors)
	//
	// The promhttp.Handler() function gathers all metrics
	// registered in the default registry and exposes them
	// in Prometheus text format.
}

// handler for getting list of devices
func getDevices(w http.ResponseWriter, r *http.Request) {
	// for converting go Struct to JSON
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, "error in conversion", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
