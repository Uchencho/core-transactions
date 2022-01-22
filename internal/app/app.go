package app

import (
	"fmt"
	"net/http"

	"github.com/Uchencho/core-proto/generated/accounts"
	"github.com/Uchencho/core-transactions/internal/db"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
)

// App is a representation of the set of functionalities for the service
type App struct {
	GetTransaction http.HandlerFunc
}

// Option is a representation of an optional arguement
type Option struct {
	AccountClient  accounts.ClientClient
	GetTransaction db.GetTransactionFunc
}

// Options is variadic function for configuring an optional arguement
type Options func(oa *Option)

// New creates a new app
func New(grpcClient grpc.ClientConnInterface, opts ...Options) App {

	o := Option{
		AccountClient: accounts.NewClientClient(grpcClient),
	}

	for _, option := range opts {
		option(&o)
	}

	getTransaction := GetTransactionHandler(o.AccountClient, o.GetTransaction)

	return App{
		GetTransaction: getTransaction,
	}
}

// Handler returns the main Handler for the application
func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, fmt.Sprintf("/transactions/accounts/:%s", accountIdParam), a.GetTransaction)

	h := http.HandlerFunc(router.ServeHTTP)
	return h
}
