package app

import (
	"log"
	"net"
	"net/http"
	"sync"
)

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		log.Fatal(a.runAuthGRPC())
	}()

	go func() {
		defer wg.Done()
		log.Fatal(a.runHTTP())
	}()

	wg.Wait()
	return nil
}

func (a *App) runAuthGRPC() error {
	listener, err := net.Listen("tcp", a.appConfig.App.PortAuthGRPC)
	if err != nil {
		return err
	}

	log.Println("Shop GRPC server running on port:", a.appConfig.App.PortAuthGRPC)

	return a.grpcShopServer.Serve(listener)
}

func (a *App) runHTTP() error {
	log.Println("Shop HTTP server is running on port:", a.appConfig.App.PortAuthHTTP)

	return http.ListenAndServe(a.appConfig.App.PortAuthHTTP, a.muxAuth)
}
