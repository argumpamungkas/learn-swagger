// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars": {
            "get": {
                "description": "Get details of all car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Cars"
                ],
                "summary": "Get All cars",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            },
            "post": {
                "description": "Post create a new car, NOTE : id auto increment, so in body car_id is deleted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Cars"
                ],
                "summary": "Post create car",
                "parameters": [
                    {
                        "description": "create car",
                        "name": "models.Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            }
        },
        "/cars/{carID}": {
            "get": {
                "description": "Get details of car corresponding to the input car_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Cars"
                ],
                "summary": "Get details for a given car_id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            },
            "put": {
                "description": "Update data car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Cars"
                ],
                "summary": "Updated car data with car_id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "car_id of the car to be updated",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated car",
                        "name": "models.Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleted data car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deleted Cars"
                ],
                "summary": "Deleted car data with car_id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "car_id of the car to be deleted",
                        "name": "carID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Car": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "car_id": {
                    "type": "integer"
                },
                "model": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Car API",
	Description:      "This is a sample service for managing cars",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}