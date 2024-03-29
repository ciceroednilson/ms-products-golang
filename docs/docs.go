// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Cícero Ednilson",
            "url": "https://www.linkedin.com/in/c%C3%ADcero-ednilson-de-barros-machado-35153888/",
            "email": "ciceroednilson@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "find array of products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.ProductResponse"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "update register of product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Product",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "create register of product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Product",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products/{key}/": {
            "delete": {
                "description": "delete a product details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Product",
                "parameters": [
                    {
                        "type": "string",
                        "format": "int",
                        "description": "delete product by key",
                        "name": "key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/products/{key}/details": {
            "get": {
                "description": "find a product details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Product",
                "parameters": [
                    {
                        "type": "string",
                        "format": "int",
                        "description": "product search by key",
                        "name": "key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ProductResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ProductRequest": {
            "description": "ProductRequest Product fields",
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.ProductResponse": {
            "description": "ProductResponse Product fields",
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    },
    "externalDocs": {
        "description": "Create Swagger docs with swaggo",
        "url": "https://github.com/swaggo/http-swagger"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Product API",
	Description:      "This is a api to manager the products .",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
