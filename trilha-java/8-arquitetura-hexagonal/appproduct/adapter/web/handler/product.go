package handler

import (
	"encoding/json"
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/adapter/dto"
	"github.com/bodegami/trilha-dev-full-cycle-java/arquitetura-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			createErrorResponse(http.StatusNotFound, err.Error(), w)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			createErrorResponse(http.StatusBadRequest, err.Error(), w)
			return
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			createErrorResponse(http.StatusNotFound, err.Error(), w)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			createErrorResponse(http.StatusNotFound, err.Error(), w)
			return
		}
		result, err := service.Disable(product)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			createErrorResponse(http.StatusInternalServerError, err.Error(), w)
			return
		}
	})
}

func createErrorResponse(statusError int, messageError string, w http.ResponseWriter) {
	w.WriteHeader(statusError)
	w.Write(jsonError(messageError))
}
