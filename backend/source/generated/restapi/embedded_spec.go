// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Public API for MochiNote",
    "title": "MochiNote",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/login": {
      "post": {
        "description": "handle login request, username and password",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "LoginV1"
        ],
        "summary": "Sign up or log in",
        "operationId": "loginV1",
        "parameters": [
          {
            "description": "user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/userObject"
            }
          },
          {
            "type": "boolean",
            "default": false,
            "description": "Do sign up operation if set to true",
            "name": "signup",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/errResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "errResponse": {
      "type": "object",
      "properties": {
        "errMessage": {
          "description": "error message",
          "type": "string"
        },
        "status_code": {
          "description": "http error code",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "loginResponse": {
      "type": "object",
      "properties": {
        "username": {
          "description": "username of the user",
          "type": "string"
        }
      }
    },
    "userObject": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "description": "password of the user",
          "type": "string"
        },
        "username": {
          "description": "username of the user",
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the Public API for MochiNote",
    "title": "MochiNote",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/login": {
      "post": {
        "description": "handle login request, username and password",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "LoginV1"
        ],
        "summary": "Sign up or log in",
        "operationId": "loginV1",
        "parameters": [
          {
            "description": "user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/userObject"
            }
          },
          {
            "type": "boolean",
            "default": false,
            "description": "Do sign up operation if set to true",
            "name": "signup",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/errResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "errResponse": {
      "type": "object",
      "properties": {
        "errMessage": {
          "description": "error message",
          "type": "string"
        },
        "status_code": {
          "description": "http error code",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "loginResponse": {
      "type": "object",
      "properties": {
        "username": {
          "description": "username of the user",
          "type": "string"
        }
      }
    },
    "userObject": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "description": "password of the user",
          "type": "string"
        },
        "username": {
          "description": "username of the user",
          "type": "string"
        }
      }
    }
  }
}`))
}