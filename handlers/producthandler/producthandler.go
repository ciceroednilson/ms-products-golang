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

// Create a product
// @Summary      Product
// @Description  create register of product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request body request.ProductRequest true "query params"
// @Success      200
// @Router       /products [post]
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

// Search a product
// @Summary      Product
// @Description  find a product
// @Tags         product
// @Produce      json
// @Param 		 key query string  false  "product search by key"  Format(int)
// @Success      200 {object} response.ProductResponse
// @Router       /products [get]
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

// Search a product with details
// @Summary      Product
// @Description  find a product details
// @Tags         product
// @Produce      json
// @Param 		 key query string  false  "product search by key"  Format(int)
// @Success      200 {object} response.ProductResponse
// @Router       /products/{key}/details [get]
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

// Search array of products
// @Summary      Product
// @Description  find array of products
// @Tags         product
// @Produce      json
// @Success      200 {array} response.ProductResponse
// @Router       /products [get]
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

// Update a product
// @Summary      Product
// @Description  update register of product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request body request.ProductRequest true "query params"
// @Success      200
// @Router       /products [put]
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

// Delete a product with details
// @Summary      Product
// @Description  delete a product details
// @Tags         product
// @Produce      json
// @Param 		 key query string  false  "delete product by key"  Format(int)
// @Success      200
// @Router       /products/{key}/ [delete]
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
