// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/admin/category-create": {
            "post": {
                "description": "创建分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "parentId",
                        "name": "parentId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-delete": {
            "delete": {
                "description": "删除分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-list": {
            "get": {
                "description": "获取分类列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "分类列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category-update": {
            "put": {
                "description": "修改分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "parentId",
                        "name": "parentId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/problem-create": {
            "post": {
                "description": "创建问题",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "创建问题",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_runtime",
                        "name": "max_runtime",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "max_mem",
                        "name": "max_mem",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "category_ids",
                        "name": "category_ids",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "test_cases",
                        "name": "test_cases",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/problem-update": {
            "put": {
                "description": "修改问题",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员私有方法"
                ],
                "summary": "修改问题",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "identity",
                        "name": "identity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "max_runtime",
                        "name": "max_runtime",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "max_mem",
                        "name": "max_mem",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "category_ids",
                        "name": "category_ids",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "test_cases",
                        "name": "test_cases",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-detail": {
            "get": {
                "description": "问题详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "problem_identity",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/problem-list": {
            "get": {
                "description": "获取问题列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "问题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category_identity",
                        "name": "category_identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rank-list": {
            "get": {
                "description": "用户排行榜",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户排行榜",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "用户注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_code",
                        "name": "userCode",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_phone",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "user_mail",
                        "name": "mail",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/send-code": {
            "post": {
                "description": "发送验证码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_mail",
                        "name": "mail",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/submit-list": {
            "get": {
                "description": "提交列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "提交列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "problem_identity",
                        "name": "problem_identity",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user_identity",
                        "name": "user_identity",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-detail": {
            "get": {
                "description": "用户详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user_identity",
                        "name": "identity",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}