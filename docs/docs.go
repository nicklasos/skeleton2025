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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/v1/cities": {
            "get": {
                "description": "Get a list of all cities",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cities"
                ],
                "summary": "List all cities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_cities.CitiesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_cities.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "big.Int": {
            "type": "object"
        },
        "internal_cities.CitiesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal_cities.City"
                    }
                }
            }
        },
        "internal_cities.City": {
            "type": "object",
            "properties": {
                "available_modules": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "created_at": {
                    "$ref": "#/definitions/pgtype.Timestamp"
                },
                "id": {
                    "type": "integer"
                },
                "lat": {
                    "$ref": "#/definitions/pgtype.Numeric"
                },
                "lon": {
                    "$ref": "#/definitions/pgtype.Numeric"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "$ref": "#/definitions/pgtype.Timestamp"
                }
            }
        },
        "internal_cities.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "pgtype.InfinityModifier": {
            "type": "integer",
            "enum": [
                1,
                0,
                -1
            ],
            "x-enum-varnames": [
                "Infinity",
                "Finite",
                "NegativeInfinity"
            ]
        },
        "pgtype.Numeric": {
            "type": "object",
            "properties": {
                "exp": {
                    "type": "integer"
                },
                "infinityModifier": {
                    "$ref": "#/definitions/pgtype.InfinityModifier"
                },
                "int": {
                    "$ref": "#/definitions/big.Int"
                },
                "naN": {
                    "type": "boolean"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "pgtype.Timestamp": {
            "type": "object",
            "properties": {
                "infinityModifier": {
                    "$ref": "#/definitions/pgtype.InfinityModifier"
                },
                "time": {
                    "description": "Time zone will be ignored when encoding to PostgreSQL.",
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyApp API",
	Description:      "Api for SmartCity project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
