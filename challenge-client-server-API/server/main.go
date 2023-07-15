package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type currencyQuote struct {
	USDBRL struct {
		Code       string `json:"code"`        // Código
		Codein     string `json:"codein"`      // Código Entrada
		Name       string `json:"name"`        // Nome
		High       string `json:"high"`        // Alta
		Low        string `json:"low"`         // Baixa
		VarBid     string `json:"varBid"`      // Variação Bid
		PctChange  string `json:"pctChange"`   // Percentual de Mudança
		Bid        string `json:"bid"`         // Oferta
		Ask        string `json:"ask"`         // Perguntar
		Timestamp  string `json:"timestamp"`   // Carimbo de data/hora
		CreateDate string `json:"create_date"` // Data de criação
	} `json:"USDBRL"`
}

func main() {
	// Criação do router utilizando Gorilla Mux.
	// Create the router using Gorilla Mux.
	router := mux.NewRouter()

	// Rota para manipular as requisições de cotação com um parâmetro coin.
	// Route to handle quote requests with a coin parameter.
	router.HandleFunc("/quote/{coin}", handleQuoteRequest)
	// Configuração do servidor HTTP.
	// HTTP server configuration.
	server := &http.Server{
		Handler: router, // Handler é o manipulador das requisições HTTP. Ele especifica a lógica para lidar com as requisições recebidas.
		// Handler is the HTTP request handler. It specifies the logic for handling incoming requests.
		Addr: ":8082", // Addr é o endereço TCP em que o servidor HTTP será executado.
		// Addr is the TCP address on which the HTTP server will run.
		WriteTimeout: 15 * time.Second, // Tempo limite para gravação da resposta HTTP. // Write timeout for HTTP response.
		ReadTimeout:  15 * time.Second, // Tempo limite para leitura da requisição HTTP. // Read timeout for HTTP request.
	}

	log.Fatal(server.ListenAndServe())
}

func handleQuoteRequest(w http.ResponseWriter, r *http.Request) {
	// Obtenção dos parâmetros da rota.
	// Get path parameters from the route.
	pathParams := mux.Vars(r)
	coin := pathParams["coin"]
	coin = strings.ToLower(coin)
	log.Println("Requested Coin:", coin)

	// Criação do contexto com tempo limite de 2 segundos.
	// Create a context with a timeout of 2 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	coinMappings := map[string]string{
		"usd":               "USD",
		"dolar":             "USD",
		"dólar":             "USD",
		"dollar":            "USD",
		"dóllar":            "USD",
		"dollares":          "USD",
		"dóllares":          "USD",
		"dollars":           "USD",
		"dóllars":           "USD",
		"dolares":           "USD",
		"dólares":           "USD",
		"americandollar":    "USD",
		"american_dollar":   "USD",
		"american_dóllar":   "USD",
		"american_dólar":    "USD",
		"american_dollares": "USD",
		"american_dóllares": "USD",
		"american_dollars":  "USD",
		"american_dóllars":  "USD",
		"american_dolares":  "USD",
		"american_dólares":  "USD",
	}

	// Verificação se a moeda solicitada é suportada.
	// Check if the requested coin is supported.
	if coin, ok := coinMappings[coin]; !ok {
		log.Println("Unsupported Coin:", coin)
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	log.Println("Supported Coin:", coin)
	coin = coinMappings[coin]

	// Chamada da função para obter a cotação da moeda.
	// Call the function to get the currency quote.
	quote, err := getQuoteOfCoin(ctx, coin)
	if err != nil {
		if err == context.DeadlineExceeded {
			http.Error(w, "408 Request Timeout", http.StatusRequestTimeout)
			log.Println("Timeout exceeded. Execution time was insufficient.")
			return
		}

		log.Println(err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Configuração do cabeçalho da resposta com o tipo JSON.
	// Set the response header to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Codificação da resposta em formato JSON e envio ao cliente.
	// Encode the response to JSON format and send it to the client.
	json.NewEncoder(w).Encode(quote)
}

func getQuoteOfCoin(ctx context.Context, coin string) (string, error) {
	var quote currencyQuote

	// Construção da URL da API de cotação com base na moeda solicitada.
	// Construct the URL for the currency quote API based on the requested coin.
	requestURL := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s-BRL", coin)

	// Criação da requisição HTTP com o contexto e a URL configurados.
	// Create an HTTP request with the configured context and URL.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return quote.USDBRL.Bid, err
	}

	// Criação do cliente HTTP.
	// Create an HTTP client.
	// client := &http.Client{}

	// Envio da requisição e obtenção da resposta.
	// Send the request and get the response.
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return quote.USDBRL.Bid, err
	}
	defer response.Body.Close()

	// Decodificação da resposta JSON para a estrutura currencyQuote.
	// Decode the JSON response into the currencyQuote structure.
	if err := json.NewDecoder(response.Body).Decode(&quote); err != nil {
		return quote.USDBRL.Bid, err
	}

	return quote.USDBRL.Bid, nil
}
