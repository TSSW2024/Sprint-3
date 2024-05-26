package httprequestt

import (
	"encoding/json"
	"fmt"
	"net/http"
	info "utemtrading/api/infobinance"
)

// Esta función maneja las solicitudes HTTP a la ruta /cryptoprices.
func HandleCryptoPrices(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cryptoprices" {
		http.NotFound(w, r)
		return
	}

	prices := info.GetCryptoPrices()

	// Crear un mapa anidado con el nombre "precios"
	response := map[string]interface{}{
		"precios": prices,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al convertir a JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Esta función maneja las solicitudes HTTP a la ruta /cryptoprice
func HandleSingleCryptoPrice(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "Symbol parameter is missing", http.StatusBadRequest)
		return
	}

	price, err := info.GetSingleCryptoPrice(symbol)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el precio: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"symbol": symbol,
		"price":  price,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al convertir a JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
