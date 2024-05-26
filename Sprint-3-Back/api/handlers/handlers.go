package handlersrutas

import (
	"log"
	"net/http"
	httprequestt "utemtrading/api/httprequest"
)

func Handlerfuns() {
	//Obtenemos todos los precios de las criptos ejemplo: http://localhost:8080/cryptoprices
	http.HandleFunc("/cryptoprices", httprequestt.HandleCryptoPrices)

	//Obtenemos solo una moneda ejemplo http://localhost:8080/cryptoprice?symbol=BTCUSDT
	http.HandleFunc("/cryptoprice", httprequestt.HandleSingleCryptoPrice)

	puerto := ":8080"
	log.Printf("Servidor en ejecución en http://localhost%s\n", puerto)
	log.Fatal(http.ListenAndServe(puerto, nil))
}
