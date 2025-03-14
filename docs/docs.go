// Code generated by swaggo/swag. DO NOT EDIT
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
        "/checkout": {
            "post": {
                "description": "Buy tickets for a specific event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Buy tickets for an event",
                "parameters": [
                    {
                        "description": "Input data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.BuyTicketsInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.BuyTicketsOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "Get all events with their details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "List all events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.ListEventsOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new event with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Create a new event",
                "parameters": [
                    {
                        "description": "Input data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateEventInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateEventOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/events/{eventID}": {
            "get": {
                "description": "Get details of an event by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Get event details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.GetEventOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/events/{eventID}/spots": {
            "get": {
                "description": "List all spots for a specific event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "List spots for an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.ListSpotsOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a specified number of spots for an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Events"
                ],
                "summary": "Create spots for an event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Input data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CreateSpotsRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateSpotsOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.CreateSpotsRequest": {
            "type": "object",
            "properties": {
                "number_of_spots": {
                    "type": "integer"
                }
            }
        },
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "usecases.BuyTicketsInput": {
            "type": "object",
            "properties": {
                "card_hash": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "event_id": {
                    "type": "string"
                },
                "spots": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ticket_type": {
                    "type": "string"
                }
            }
        },
        "usecases.BuyTicketsOutput": {
            "type": "object",
            "properties": {
                "tickets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecases.TicketOutput"
                    }
                }
            }
        },
        "usecases.CreateEventInput": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateEventOutput": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateSpotsOutput": {
            "type": "object",
            "properties": {
                "spots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecases.SpotOutput"
                    }
                }
            }
        },
        "usecases.EventOutput": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecases.GetEventOutput": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "organization": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "string"
                }
            }
        },
        "usecases.ListEventsOutput": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecases.EventOutput"
                    }
                }
            }
        },
        "usecases.ListSpotsOutput": {
            "type": "object",
            "properties": {
                "event": {
                    "$ref": "#/definitions/usecases.EventOutput"
                },
                "spots": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecases.SpotOutput"
                    }
                }
            }
        },
        "usecases.SpotOutput": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "reserved": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string"
                },
                "ticket_id": {
                    "type": "string"
                }
            }
        },
        "usecases.TicketOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "spot_id": {
                    "type": "string"
                },
                "ticket_type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Events API",
	Description:      "This is a server for managing events. Imersão Full Cycle",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
