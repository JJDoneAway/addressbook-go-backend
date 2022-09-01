// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Johannes Höhne",
            "email": "Johannes.Hoehne1@mail.schwarz"
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
        "/addresses": {
            "get": {
                "description": "Provide a list of all currently known addresses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "List all addresses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Address"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Will add a new addresses entity to the storage. The new created addresses will be returned. Don't add the Id to the addresses parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Add a new addresses",
                "parameters": [
                    {
                        "description": "The new addresses without ID",
                        "name": "addresses",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "ID must be zero, Unparsable JSON body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Will delete all addresses. an empty list will be returned",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Delete all addresses",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/addresses/{id}": {
            "get": {
                "description": "Get a address with the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Get one address",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the user",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    },
                    "400": {
                        "description": "Unknown ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Will update an existing address which is identified via its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Update an existing address",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the address",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The new address without ID",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Address"
                        }
                    },
                    "400": {
                        "description": "Unknown ID, ID miss match, Unparsable JSON body",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a address with the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Delete one address",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the address",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Unknown ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "Provide a list of all currently known Prometheus metrics",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "metrics"
                ],
                "summary": "List all Prometheus metrics",
                "responses": {
                    "200": {
                        "description": "Prometheus metrics line by line",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Address": {
            "type": "object",
            "required": [
                "email",
                "first-name",
                "last-name",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first-name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last-name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GO Example Addressbook",
	Description:      "This is a simple GO application with some basic REST CRUD operations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
