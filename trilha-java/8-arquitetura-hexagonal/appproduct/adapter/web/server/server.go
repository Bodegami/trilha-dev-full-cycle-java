package server

import (
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/adapter/web/handler"
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{service}
}

func (w *WebServer) Serve() {

	router := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, n, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.LstdFlags),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
