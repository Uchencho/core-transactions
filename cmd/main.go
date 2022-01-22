package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Uchencho/core-transactions/internal"
	"github.com/Uchencho/core-transactions/internal/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

func getServerAddress() string {
	const defaultServerAddress = "127.0.0.1:8000"
	serverAddress, present := os.LookupEnv("PORT")
	if present {
		return serverAddress
	}
	return defaultServerAddress
}

func getGRPCAddress() string {
	const defaultServerAddress = "127.0.0.1:4444"
	serverAddress, present := os.LookupEnv("PORT")
	if present {
		return serverAddress
	}
	return defaultServerAddress
}

func main() {
	port := getServerAddress()

	opts := grpc.WithInsecure()
	connOpts := grpc.WithConnectParams(grpc.ConnectParams{
		Backoff:           backoff.DefaultConfig,
		MinConnectTimeout: 5 * time.Second,
	})
	addr := getGRPCAddress()

	clientConn, err := grpc.Dial(addr, opts, connOpts)
	if err != nil {
		log.Fatal(err)
	}

	appOpt := func(o *app.Option) {
		o.GetTransaction = func(accountId string) (internal.Transaction, error) {
			return internal.Transaction{}, nil
		}
	}

	a := app.New(clientConn, appOpt)
	log.Println(fmt.Sprintf("Starting server on address: %s", port))
	log.Println(http.ListenAndServe(port, a.Handler()))
}
