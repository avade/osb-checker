// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
)

// ServiceBindingUnbindingOKCode is the HTTP code returned for type ServiceBindingUnbindingOK
const ServiceBindingUnbindingOKCode int = 200

/*ServiceBindingUnbindingOK OK

swagger:response serviceBindingUnbindingOK
*/
type ServiceBindingUnbindingOK struct {

	/*
	  In: Body
	*/
	Payload models.Object `json:"body,omitempty"`
}

// NewServiceBindingUnbindingOK creates ServiceBindingUnbindingOK with default headers values
func NewServiceBindingUnbindingOK() *ServiceBindingUnbindingOK {

	return &ServiceBindingUnbindingOK{}
}

// WithPayload adds the payload to the service binding unbinding o k response
func (o *ServiceBindingUnbindingOK) WithPayload(payload models.Object) *ServiceBindingUnbindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding o k response
func (o *ServiceBindingUnbindingOK) SetPayload(payload models.Object) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ServiceBindingUnbindingAcceptedCode is the HTTP code returned for type ServiceBindingUnbindingAccepted
const ServiceBindingUnbindingAcceptedCode int = 202

/*ServiceBindingUnbindingAccepted Accepted

swagger:response serviceBindingUnbindingAccepted
*/
type ServiceBindingUnbindingAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.AsyncOperation `json:"body,omitempty"`
}

// NewServiceBindingUnbindingAccepted creates ServiceBindingUnbindingAccepted with default headers values
func NewServiceBindingUnbindingAccepted() *ServiceBindingUnbindingAccepted {

	return &ServiceBindingUnbindingAccepted{}
}

// WithPayload adds the payload to the service binding unbinding accepted response
func (o *ServiceBindingUnbindingAccepted) WithPayload(payload *models.AsyncOperation) *ServiceBindingUnbindingAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding accepted response
func (o *ServiceBindingUnbindingAccepted) SetPayload(payload *models.AsyncOperation) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingBadRequestCode is the HTTP code returned for type ServiceBindingUnbindingBadRequest
const ServiceBindingUnbindingBadRequestCode int = 400

/*ServiceBindingUnbindingBadRequest Bad Request

swagger:response serviceBindingUnbindingBadRequest
*/
type ServiceBindingUnbindingBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingBadRequest creates ServiceBindingUnbindingBadRequest with default headers values
func NewServiceBindingUnbindingBadRequest() *ServiceBindingUnbindingBadRequest {

	return &ServiceBindingUnbindingBadRequest{}
}

// WithPayload adds the payload to the service binding unbinding bad request response
func (o *ServiceBindingUnbindingBadRequest) WithPayload(payload *models.Error) *ServiceBindingUnbindingBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding bad request response
func (o *ServiceBindingUnbindingBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingUnbindingGoneCode is the HTTP code returned for type ServiceBindingUnbindingGone
const ServiceBindingUnbindingGoneCode int = 410

/*ServiceBindingUnbindingGone Gone

swagger:response serviceBindingUnbindingGone
*/
type ServiceBindingUnbindingGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingUnbindingGone creates ServiceBindingUnbindingGone with default headers values
func NewServiceBindingUnbindingGone() *ServiceBindingUnbindingGone {

	return &ServiceBindingUnbindingGone{}
}

// WithPayload adds the payload to the service binding unbinding gone response
func (o *ServiceBindingUnbindingGone) WithPayload(payload *models.Error) *ServiceBindingUnbindingGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding unbinding gone response
func (o *ServiceBindingUnbindingGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingUnbindingGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}