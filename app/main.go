package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisições HTTP recebidas",
		},
		[]string{"path", "status"},
	)

	serviceUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_up",
			Help: "Disponibilidade do serviço (1 para UP, 0 para DOWN)",
		},
	)
)

func initi() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(serviceUp)

	serviceUp.Set(1)
}

type ResponseBody struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpRequestsTotal.WithLabelValues(r.URL.Path, "405").Inc()
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now().UTC().Format(time.RFC3339)

	response := ResponseBody{
		Nome:    "Projeto Korp",
		Horario: currentTime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	httpRequestsTotal.WithLabelValues("/projeto-korp", "200").Inc()
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	 
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Servidor HTTP rodando na porte 8080...")
	http.ListenAndServe(":8080", nil)
}