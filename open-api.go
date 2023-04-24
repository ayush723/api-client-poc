// Package rest the API used to use the Map Authority map management server
//
// This is the public REST interface to the map management server
//
//	    Schemes: http
//	    Host: localhost:9000
//	    BasePath: /v1
//	    Version: 1.0.0
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//		   - application/x-zserio-object
//
//	    Security:
//	    - ApiKeyAuth: []
//
//	    SecurityDefinitions:
//	    ApiKeyAuth:
//	         type: apiKey
//	         name: x-api-key
//	         in: header
//
// swagger:meta
package main

import "embed"

// openAPIContent contains openapi.json file generated via go-swagger from spec
//
// //go:generate swagger-codegen generate -i openapi.json -l python -o python-client
//
//go:generate swagger generate spec -m -w . -o ./openapi.json
//go:generate swagger-codegen generate -i openapi.json -l go -o swagger
//go:embed openapi.json
var openAPIContent embed.FS
