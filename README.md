## Punto 1

- en la carpeta infobinance: GetCryptoPrices y GetSingleCryptoPrice obtenemos los precios de todas las monedas y de una moneda.
- en la carpeta hhtprequest: HandleCryptoPrices Esta función maneja las solicitudes a la ruta /cryptoprices y los datos obtenidos de GetCryptoPrices los convierte en formato JSON y HandleCryptoPrices Esta función maneja las solicitudes a la ruta /cryptoprice (en singular) y los datos obtenidos de GetSingleCryptoPrice los convierte en formato JSON

## punto 2

- para iniciar el proyecto solo go run .
- una vez iniciado use alguna de las url de ejemplo para optener todas las monedas o una

- en la carpeta handlers:
  /\*
  func Handlerfuns() {
  //Obtenemos todos los precios de las criptos ejemplo: http://localhost:8080/cryptoprices
  http.HandleFunc("/cryptoprices", handleCryptoPrices)

      //Obtenemos solo una moneda ejemplo http://localhost:8080/cryptoprice?symbol=BTCUSDT
      http.HandleFunc("/cryptoprice", handleSingleCryptoPrice)

      puerto := ":8080"
      log.Printf("Servidor en ejecución en http://localhost%s\n", puerto)
      log.Fatal(http.ListenAndServe(puerto, nil))

}
\*/

## punto 3

- solo tiene 2 librerias
  -go get github.com/adshao/go-binance/v2
  -go get github.com/joho/godotenv
