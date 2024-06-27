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
        "/area": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all areas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.AreaResponseDTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new area",
                "parameters": [
                    {
                        "description": "Area data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AreaCreateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AreaResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/area/{area_id}/point": {
            "post": {
                "description": "Creates a new point associated with an area.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Point"
                ],
                "summary": "Create Point",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Area ID",
                        "name": "area_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Point data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PointDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PointResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/area/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get an area by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Area ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AreaResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an area",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Area ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Area data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AreaUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AreaResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete an area",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Area ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/area/{id}/point": {
            "get": {
                "description": "Retrieves all points associated with an area by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Point"
                ],
                "summary": "Get Points by Area ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Area ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.PointResponseDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/dashboard": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns the user dashboard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User dashboard",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/point/{id}": {
            "put": {
                "description": "Updates a point associated with an area by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Point"
                ],
                "summary": "Update Point",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Point ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated point data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PointUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.PointResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a point associated with an area by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Point"
                ],
                "summary": "Delete Point",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Point ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/report": {
            "post": {
                "description": "creates report in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "Create Report",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReportCreatedDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    }
                }
            }
        },
        "/report/": {
            "put": {
                "description": "updates report in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "Update Report",
                "parameters": [
                    {
                        "description": "Report Update Params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReportUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    }
                }
            }
        },
        "/report/user/{id}": {
            "get": {
                "description": "Get all reports from a user by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "Get Reports by User ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Report"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    }
                }
            }
        },
        "/report/{id}": {
            "get": {
                "description": "gets report from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "Get Report",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "report ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ReportResponseDTO"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "returns all users in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get All User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserResponseDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    }
                }
            }
        },
        "/user/edit": {
            "put": {
                "description": "edits user information in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Edit User",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "returns token for login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "registers user in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserCreatedDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Returns a user from the database by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AreaCreateDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.AreaResponseDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "points": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PointResponseDTO"
                    }
                }
            }
        },
        "dto.AreaUpdateDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.PhotoDTO": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.PointDTO": {
            "type": "object",
            "required": [
                "latitude",
                "longitude"
            ],
            "properties": {
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                }
            }
        },
        "dto.PointResponseDTO": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.PointUpdateDTO": {
            "type": "object",
            "properties": {
                "area_id": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                }
            }
        },
        "dto.ReportCreatedDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "report_status_id": {
                    "type": "integer"
                },
                "report_type_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.ReportResponseDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "report": {
                    "$ref": "#/definitions/model.Report"
                }
            }
        },
        "dto.ReportUpdateDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PhotoDTO"
                    }
                },
                "report_id": {
                    "type": "integer"
                },
                "report_status_id": {
                    "type": "integer"
                },
                "report_type_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.UserCreatedDTO": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponseDTO": {
            "type": "object",
            "required": [
                "createdAt",
                "email",
                "firstname",
                "id",
                "lastname",
                "role",
                "username"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserUpdateDTO": {
            "type": "object",
            "required": [
                "userid"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "userid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "report_id": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Photo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "report": {
                    "$ref": "#/definitions/model.Report"
                },
                "reportID": {
                    "type": "integer"
                }
            }
        },
        "model.Report": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Comment"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_update": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Photo"
                    }
                },
                "report_date": {
                    "type": "string"
                },
                "report_status": {
                    "$ref": "#/definitions/model.ReportStatus"
                },
                "report_status_id": {
                    "type": "integer"
                },
                "report_type": {
                    "$ref": "#/definitions/model.ReportType"
                },
                "report_type_id": {
                    "type": "integer"
                },
                "updates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ReportUpdate"
                    }
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.ReportStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.ReportType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.ReportUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "report": {
                    "$ref": "#/definitions/model.Report"
                },
                "report_id": {
                    "type": "integer"
                },
                "report_status": {
                    "$ref": "#/definitions/model.ReportStatus"
                },
                "report_status_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "creationDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/model.Role"
                },
                "role_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
