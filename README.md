ms-products-golang




go install github.com/swaggo/swag/cmd/swag@latest

vi ~/.profile 

export PATH=$(go env GOPATH)/bin:$PATH

source ~/.profile 

swag init


//create swagger docs

swag init -g cmd/main.go 


Docs

https://github.com/swaggo/http-swagger

swag init -g cmd/main.go handlers/producthandler/producthandler.go 
