package producthandler

import (
	"encoding/json"
	"go/core/ports"
	"go/handlers/request"
	"go/handlers/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductUsercasePort ports.ProductUsercasePort
}

func NewProductHandler(productUsercasePort ports.ProductUsercasePort) ProductHandler {
	return ProductHandler{
		ProductUsercasePort: productUsercasePort,
	}
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request request.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	domain := request.ToDomain()
	err = p.ProductUsercasePort.Save(domain)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, nil)
}

func (p *ProductHandler) Read(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key, _ := strconv.Atoi(params["key"])
	product, err := p.ProductUsercasePort.FindById(key)
	if err != nil {
		panic(err)
	}
	response := response.NewProductResponse(product)
	responseJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, responseJson)
}

func (p *ProductHandler) ReadDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key, _ := strconv.Atoi(params["key"])
	product, err := p.ProductUsercasePort.FindByIdDetails(key)
	if err != nil {
		panic(err)
	}
	response := response.NewProductResponse(product)
	responseJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, responseJson)
}

func (p *ProductHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.ProductUsercasePort.FindAll()
	if err != nil {
		panic(err)
	}
	response := response.NewListProductResponse(products)
	responseJson, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, responseJson)
}

func (p *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request request.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	domain := request.ToDomain()
	err = p.ProductUsercasePort.Update(domain)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, nil)
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key, _ := strconv.Atoi(params["key"])
	product, err := p.ProductUsercasePort.FindById(key)
	if err != nil {
		panic(err)
	}
	if product == nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	p.ProductUsercasePort.Delete(key)
	responseSuccessfully(w, nil)
}

func responseSuccessfully(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	if body != nil {
		w.Write(body)
	}
	w.WriteHeader(http.StatusOK)
}
