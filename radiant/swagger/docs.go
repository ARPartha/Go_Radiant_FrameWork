package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
        "swagger": "2.0",
        "info": {
            "title": "Radiant Test",
            "description": "High performance, minimalist Go web framework with ease of customization in mind!",
            "version": "1.0.0",
            "termsOfService": "https://www.w3engineers.com",
            
            "license": {
                "name": "Apache 2.0",
                "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
            }
        },
        
        "paths": {
            "/object": {
                "get": {
                    "tags": [
                        "object"
                    ],
                    "description": "get all objects",
                    "operationId": "ObjectController.GetAll",
                    "responses": {
                        "200": {
                            "description": "",
                            "schema": {
                                "$ref": "#/definitions/models.Object"
                            }
                        },
                        "403": {
                            "description": ":objectId is empty"
                        }
                    }
                },
                "post": {
                    "tags": [
                        "object"
                    ],
                    "description": "create object",
                    "operationId": "ObjectController.Create",
                    "parameters": [
                        {
                            "in": "body",
                            "name": "body",
                            "description": "The object content",
                            "required": true,
                            "schema": {
                                "$ref": "#/definitions/models.Object"
                            }
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": "{string} models.Object"
                        },
                        "403": {
                            "description": "body is empty"
                        }
                    }
                }
            },
            "/users": {
                "get": {
                    "tags": [
                        "user"
                    ],
                    "description": "get all Users",
                    "operationId": "controllers.Getuser",
                    "responses": {
                        "200": {
                            "description": "",
                            "schema": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                },
                "post": {
                    "tags": [
                        "user"
                    ],
                    "description": "create users",
                    "operationId": "controllers.Createuser",
                    "parameters": [
                        {
                            "in": "body",
                            "name": "body",
                            "description": "body for user content",
                            "required": true,
                            "schema": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": "{int} models.User.Id"
                        },
                        "403": {
                            "description": "body is empty"
                        }
                    }
                }
            }
        },
        "definitions": {
            "models.Object": {
                "title": "Object",
                "type": "object",
                "properties": {
                    "FirstName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "LastName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "Phonenumber": {
                        "type": "string",
                        "example": "011555"
                    },
                    "Age": {
                        "type": "string",
                        "example": "0"
                    },
                    "Email": {
                        "type": "string",
                        "example": "x@gmail.com"
                    },
                    "Password": {
                        "type": "string",
                        "example": "****"
                    },
                    "DateOfBirth": {
                        "type": "string",
                        "example": "mm/dd/yyyy"
                    }
                }
            },
            "models.Profile": {
                "title": "User",
                "type": "object",
                "properties": {
                    "FirstName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "LastName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "Phonenumber": {
                        "type": "string",
                        "example": "011555"
                    },
                    "Age": {
                        "type": "string",
                        "example": "0"
                    },
                    "Email": {
                        "type": "string",
                        "example": "x@gmail.com"
                    },
                    "Password": {
                        "type": "string",
                        "example": "****"
                    },
                    "DateOfBirth": {
                        "type": "string",
                        "example": "mm/dd/yyyy"
                    }
                }
            },
            "models.User": {
                "title": "User",
                "type": "object",
                "properties": {
                    "FirstName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "LastName": {
                        "type": "string",
                        "example": "Name"
                    },
                    "Phonenumber": {
                        "type": "string",
                        "example": "011555"
                    },
                    "Age": {
                        "type": "string",
                        "example": "0"
                    },
                    "Email": {
                        "type": "string",
                        "example": "x@gmail.com"
                    },
                    "Password": {
                        "type": "string",
                        "example": "****"
                    },
                    "DateOfBirth": {
                        "type": "string",
                        "example": "mm/dd/yyyy"
                    }
                }
            }
        },
        "tags": [
            {
                "name": "object",
                "description": "Operations about object\n"
            },
            {
                "name": "user",
                "description": "Operations about Users\n"
            }
        ]
    
}`



type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "petstore.swagger.io",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {

	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
