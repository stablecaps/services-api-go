// Documentation of Services API for dashboard widget
//
//	 Schemes: http
//	 BasePath: /
//	 Version: 1.0.0
//	 Host: stablecaps.com
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//
// swagger:meta
package docs

import  "github.com/stablecaps/services-api-go"

// swagger:route POST /services/new service-tag idService
// NewService creates a new Service.
// responses:
//   200: dogSuccessResponse
// swagger:parameters idService
type dogParamsWrapper struct {
	// NewService creates a new Service.
	// in:body
	Body main.CreateServiceRequest
}
// Successful dog created response.
// swagger:response dogSuccessResponse
type dogResponseWrapper struct {
	// in:body
	Body main.WriteJson
}