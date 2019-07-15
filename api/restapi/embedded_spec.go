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
  "consumes": [
    "application/cloudevents+json"
  ],
  "produces": [
    "application/cloudevents+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "keptn api",
    "version": "0.1.0"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "openws"
        ],
        "operationId": "openWS",
        "responses": {
          "200": {
            "description": "Upgrading to WS"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/configure": {
      "post": {
        "tags": [
          "configure"
        ],
        "summary": "Forwards the received configure event to the eventbroker",
        "operationId": "configure",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ConfigureCE"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "configured",
            "schema": {
              "$ref": "#/definitions/ChannelInfo"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/event": {
      "post": {
        "tags": [
          "event"
        ],
        "summary": "Forwards the received event to the eventbroker",
        "operationId": "sendEvent",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/KeptnContextExtendedCE"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "forwarded",
            "schema": {
              "$ref": "#/definitions/ChannelInfo"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/project": {
      "post": {
        "tags": [
          "project"
        ],
        "summary": "Forwards the received project event to the eventbroker",
        "operationId": "project",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/CreateProjectCE"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "project created",
            "schema": {
              "$ref": "#/definitions/ChannelInfo"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ChannelInfo": {
      "type": "object",
      "required": [
        "token",
        "channelID"
      ],
      "properties": {
        "channelID": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "ConfigureCE": {
      "allOf": [
        {
          "$ref": "ce_v0_2_without_data.json#/definitions/event"
        },
        {
          "type": "object",
          "properties": {
            "data": {
              "required": [
                "org",
                "user",
                "token"
              ],
              "properties": {
                "org": {
                  "type": "string"
                },
                "token": {
                  "type": "string"
                },
                "user": {
                  "type": "string"
                }
              }
            }
          }
        }
      ]
    },
    "CreateProjectCE": {
      "allOf": [
        {
          "$ref": "ce_v0_2_without_data.json#/definitions/event"
        },
        {
          "type": "object",
          "properties": {
            "data": {
              "required": [
                "project",
                "stages"
              ],
              "properties": {
                "project": {
                  "type": "string"
                },
                "stages": {
                  "type": "array",
                  "items": {
                    "required": [
                      "name",
                      "deployment_strategy"
                    ],
                    "properties": {
                      "deployment_strategy": {
                        "type": "string"
                      },
                      "name": {
                        "type": "string"
                      },
                      "test_strategy": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      ]
    },
    "KeptnContextExtendedCE": {
      "allOf": [
        {
          "$ref": "https://raw.githubusercontent.com/cloudevents/spec/v0.2/spec.json#/definitions/event"
        },
        {
          "type": "object",
          "required": [
            "shkeptncontext"
          ],
          "properties": {
            "shkeptncontext": {
              "type": "string"
            }
          }
        }
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/cloudevents+json"
  ],
  "produces": [
    "application/cloudevents+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "keptn api",
    "version": "0.1.0"
  },
  "basePath": "/",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "openws"
        ],
        "operationId": "openWS",
        "responses": {
          "200": {
            "description": "Upgrading to WS"
          },
          "default": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int64"
                },
                "fields": {
                  "type": "string"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/configure": {
      "post": {
        "tags": [
          "configure"
        ],
        "summary": "Forwards the received configure event to the eventbroker",
        "operationId": "configure",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "allOf": [
                {
                  "type": "object",
                  "required": [
                    "specversion",
                    "id",
                    "type",
                    "source",
                    "shkeptncontext"
                  ],
                  "properties": {
                    "contenttype": {
                      "type": "string"
                    },
                    "extensions": {
                      "type": "object"
                    },
                    "id": {
                      "type": "string"
                    },
                    "shkeptncontext": {
                      "type": "string"
                    },
                    "source": {
                      "type": "string",
                      "format": "uri-reference"
                    },
                    "specversion": {
                      "type": "string"
                    },
                    "time": {
                      "type": "string",
                      "format": "date-time"
                    },
                    "type": {
                      "type": "string"
                    }
                  }
                },
                {
                  "type": "object",
                  "properties": {
                    "data": {
                      "required": [
                        "org",
                        "user",
                        "token"
                      ],
                      "properties": {
                        "org": {
                          "type": "string"
                        },
                        "token": {
                          "type": "string"
                        },
                        "user": {
                          "type": "string"
                        }
                      }
                    }
                  }
                }
              ]
            }
          }
        ],
        "responses": {
          "201": {
            "description": "configured",
            "schema": {
              "type": "object",
              "required": [
                "token",
                "channelID"
              ],
              "properties": {
                "channelID": {
                  "type": "string"
                },
                "token": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int64"
                },
                "fields": {
                  "type": "string"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/event": {
      "post": {
        "tags": [
          "event"
        ],
        "summary": "Forwards the received event to the eventbroker",
        "operationId": "sendEvent",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "allOf": [
                {
                  "type": "object",
                  "required": [
                    "specversion",
                    "id",
                    "type",
                    "source"
                  ],
                  "properties": {
                    "contenttype": {
                      "type": "string"
                    },
                    "data": {
                      "type": [
                        "object",
                        "string"
                      ]
                    },
                    "extensions": {
                      "type": "object"
                    },
                    "id": {
                      "type": "string"
                    },
                    "source": {
                      "type": "string",
                      "format": "uri-reference"
                    },
                    "specversion": {
                      "type": "string"
                    },
                    "time": {
                      "type": "string",
                      "format": "date-time"
                    },
                    "type": {
                      "type": "string"
                    }
                  }
                },
                {
                  "type": "object",
                  "required": [
                    "shkeptncontext"
                  ],
                  "properties": {
                    "shkeptncontext": {
                      "type": "string"
                    }
                  }
                }
              ]
            }
          }
        ],
        "responses": {
          "201": {
            "description": "forwarded",
            "schema": {
              "type": "object",
              "required": [
                "token",
                "channelID"
              ],
              "properties": {
                "channelID": {
                  "type": "string"
                },
                "token": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int64"
                },
                "fields": {
                  "type": "string"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/project": {
      "post": {
        "tags": [
          "project"
        ],
        "summary": "Forwards the received project event to the eventbroker",
        "operationId": "project",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "allOf": [
                {
                  "type": "object",
                  "required": [
                    "specversion",
                    "id",
                    "type",
                    "source",
                    "shkeptncontext"
                  ],
                  "properties": {
                    "contenttype": {
                      "type": "string"
                    },
                    "extensions": {
                      "type": "object"
                    },
                    "id": {
                      "type": "string"
                    },
                    "shkeptncontext": {
                      "type": "string"
                    },
                    "source": {
                      "type": "string",
                      "format": "uri-reference"
                    },
                    "specversion": {
                      "type": "string"
                    },
                    "time": {
                      "type": "string",
                      "format": "date-time"
                    },
                    "type": {
                      "type": "string"
                    }
                  }
                },
                {
                  "type": "object",
                  "properties": {
                    "data": {
                      "required": [
                        "project",
                        "stages"
                      ],
                      "properties": {
                        "project": {
                          "type": "string"
                        },
                        "stages": {
                          "type": "array",
                          "items": {
                            "required": [
                              "name",
                              "deployment_strategy"
                            ],
                            "properties": {
                              "deployment_strategy": {
                                "type": "string"
                              },
                              "name": {
                                "type": "string"
                              },
                              "test_strategy": {
                                "type": "string"
                              }
                            }
                          }
                        }
                      }
                    }
                  }
                }
              ]
            }
          }
        ],
        "responses": {
          "201": {
            "description": "project created",
            "schema": {
              "type": "object",
              "required": [
                "token",
                "channelID"
              ],
              "properties": {
                "channelID": {
                  "type": "string"
                },
                "token": {
                  "type": "string"
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "message"
              ],
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int64"
                },
                "fields": {
                  "type": "string"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ChannelInfo": {
      "type": "object",
      "required": [
        "token",
        "channelID"
      ],
      "properties": {
        "channelID": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "ConfigureCE": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "specversion",
            "id",
            "type",
            "source",
            "shkeptncontext"
          ],
          "properties": {
            "contenttype": {
              "type": "string"
            },
            "extensions": {
              "type": "object"
            },
            "id": {
              "type": "string"
            },
            "shkeptncontext": {
              "type": "string"
            },
            "source": {
              "type": "string",
              "format": "uri-reference"
            },
            "specversion": {
              "type": "string"
            },
            "time": {
              "type": "string",
              "format": "date-time"
            },
            "type": {
              "type": "string"
            }
          }
        },
        {
          "type": "object",
          "properties": {
            "data": {
              "required": [
                "org",
                "user",
                "token"
              ],
              "properties": {
                "org": {
                  "type": "string"
                },
                "token": {
                  "type": "string"
                },
                "user": {
                  "type": "string"
                }
              }
            }
          }
        }
      ]
    },
    "CreateProjectCE": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "specversion",
            "id",
            "type",
            "source",
            "shkeptncontext"
          ],
          "properties": {
            "contenttype": {
              "type": "string"
            },
            "extensions": {
              "type": "object"
            },
            "id": {
              "type": "string"
            },
            "shkeptncontext": {
              "type": "string"
            },
            "source": {
              "type": "string",
              "format": "uri-reference"
            },
            "specversion": {
              "type": "string"
            },
            "time": {
              "type": "string",
              "format": "date-time"
            },
            "type": {
              "type": "string"
            }
          }
        },
        {
          "type": "object",
          "properties": {
            "data": {
              "required": [
                "project",
                "stages"
              ],
              "properties": {
                "project": {
                  "type": "string"
                },
                "stages": {
                  "type": "array",
                  "items": {
                    "required": [
                      "name",
                      "deployment_strategy"
                    ],
                    "properties": {
                      "deployment_strategy": {
                        "type": "string"
                      },
                      "name": {
                        "type": "string"
                      },
                      "test_strategy": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      ]
    },
    "KeptnContextExtendedCE": {
      "allOf": [
        {
          "type": "object",
          "required": [
            "specversion",
            "id",
            "type",
            "source"
          ],
          "properties": {
            "contenttype": {
              "type": "string"
            },
            "data": {
              "type": [
                "object",
                "string"
              ]
            },
            "extensions": {
              "type": "object"
            },
            "id": {
              "type": "string"
            },
            "source": {
              "type": "string",
              "format": "uri-reference"
            },
            "specversion": {
              "type": "string"
            },
            "time": {
              "type": "string",
              "format": "date-time"
            },
            "type": {
              "type": "string"
            }
          }
        },
        {
          "type": "object",
          "required": [
            "shkeptncontext"
          ],
          "properties": {
            "shkeptncontext": {
              "type": "string"
            }
          }
        }
      ]
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
}`))
}
