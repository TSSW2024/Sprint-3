package binance

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

// GetCryptoPrices obtiene los precios de las criptomonedas desde Binance
func GetCryptoPrices() map[string]float64 {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		log.Fatal("API key and/or secret key are missing")
	}

	client := binance.NewClient(apiKey, secretKey)

	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		log.Fatalf("Failed to get crypto prices: %v", err)
	}

	cryptoPrices := make(map[string]float64)
	for _, p := range prices {
		priceFloat, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			log.Printf("Error al convertir el precio: %v", err)
			continue
		}
		cryptoPrices[p.Symbol] = priceFloat
	}

	return cryptoPrices
}

// GetSingleCryptoPrice obtiene el precio de una criptomoneda específica desde Binance
func GetSingleCryptoPrice(symbol string) (float64, error) {
	err := godotenv.Load()
	if err != nil {
		return 0, fmt.Errorf("error al cargar el archivo .env: %v", err)
	}

	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		return 0, fmt.Errorf("API key and/or secret key are missing")
	}

	client := binance.NewClient(apiKey, secretKey)

	price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to get crypto price: %v", err)
	}

	if len(price) == 0 {
		return 0, fmt.Errorf("no price found for symbol: %s", symbol)
	}

	priceFloat, err := strconv.ParseFloat(price[0].Price, 64)
	if err != nil {
		return 0, fmt.Errorf("error al convertir el precio: %v", err)
	}

	return priceFloat, nil
}
