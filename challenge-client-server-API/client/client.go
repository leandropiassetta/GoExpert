package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

// O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.

// Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

func main() {
	timeout := 300 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8082", nil)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout exceeded. Execution time was insufficient.")
		}
		panic(err)
	}

	// make the request
	result, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer result.Body.Close()

	// print the result to the terminal window (stdout)

	io.Copy(os.Stdout, result.Body)
}
