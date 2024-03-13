// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://github.com/MrDweller/work-handler"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/assign-worker": {
            "post": {
                "description": "Assign a worker a specified work task.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Worker"
                ],
                "summary": "Assign a worker a specified work task.",
                "parameters": [
                    {
                        "description": "AssignWorkerDTO JSON",
                        "name": "AssignWorkerDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/workhandler.AssignWorkerDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/workhandler.WorkDTO"
                        }
                    }
                }
            }
        },
        "/work": {
            "post": {
                "description": "Create a new work task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new work task",
                "parameters": [
                    {
                        "description": "CreateWorkDTO JSON",
                        "name": "CreateWorkDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/workhandler.CreateWorkDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/workhandler.WorkDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "workhandler.AssignWorkerDTO": {
            "type": "object",
            "properties": {
                "workId": {
                    "type": "string"
                },
                "workerId": {
                    "type": "string"
                }
            }
        },
        "workhandler.CreateWorkDTO": {
            "type": "object",
            "properties": {
                "eventType": {
                    "type": "string"
                },
                "productId": {
                    "type": "string"
                }
            }
        },
        "workhandler.WorkDTO": {
            "type": "object",
            "properties": {
                "eventType": {
                    "type": "string"
                },
                "productId": {
                    "type": "string"
                },
                "workId": {
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
	Title:            "Work Handler",
	Description:      "This page shows the REST interfaces offered by the Work Handler.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}