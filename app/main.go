package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseBody struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	fmt.Println("Servidor HTTP rodando na porte 8080...")
	http.ListenAndServe(":8080", nil)
}