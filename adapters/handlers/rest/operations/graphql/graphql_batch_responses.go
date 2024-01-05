//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// GraphqlBatchOKCode is the HTTP code returned for type GraphqlBatchOK
const GraphqlBatchOKCode int = 200

/*
GraphqlBatchOK Successful query (with select).

swagger:response graphqlBatchOK
*/
type GraphqlBatchOK struct {

	/*
	  In: Body
	*/
	Payload models.GraphQLResponses `json:"body,omitempty"`
}

// NewGraphqlBatchOK creates GraphqlBatchOK with default headers values
func NewGraphqlBatchOK() *GraphqlBatchOK {

	return &GraphqlBatchOK{}
}

// WithPayload adds the payload to the graphql batch o k response
func (o *GraphqlBatchOK) WithPayload(payload models.GraphQLResponses) *GraphqlBatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the graphql batch o k response
func (o *GraphqlBatchOK) SetPayload(payload models.GraphQLResponses) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GraphqlBatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.GraphQLResponses{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GraphqlBatchUnauthorizedCode is the HTTP code returned for type GraphqlBatchUnauthorized
const GraphqlBatchUnauthorizedCode int = 401

/*
GraphqlBatchUnauthorized Unauthorized or invalid credentials.

swagger:response graphqlBatchUnauthorized
*/
type GraphqlBatchUnauthorized struct {
}

// NewGraphqlBatchUnauthorized creates GraphqlBatchUnauthorized with default headers values
func NewGraphqlBatchUnauthorized() *GraphqlBatchUnauthorized {

	return &GraphqlBatchUnauthorized{}
}

// WriteResponse to the client
func (o *GraphqlBatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GraphqlBatchForbiddenCode is the HTTP code returned for type GraphqlBatchForbidden
const GraphqlBatchForbiddenCode int = 403

/*
GraphqlBatchForbidden Forbidden

swagger:response graphqlBatchForbidden
*/
type GraphqlBatchForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGraphqlBatchForbidden creates GraphqlBatchForbidden with default headers values
func NewGraphqlBatchForbidden() *GraphqlBatchForbidden {

	return &GraphqlBatchForbidden{}
}

// WithPayload adds the payload to the graphql batch forbidden response
func (o *GraphqlBatchForbidden) WithPayload(payload *models.ErrorResponse) *GraphqlBatchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the graphql batch forbidden response
func (o *GraphqlBatchForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GraphqlBatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GraphqlBatchUnprocessableEntityCode is the HTTP code returned for type GraphqlBatchUnprocessableEntity
const GraphqlBatchUnprocessableEntityCode int = 422

/*
GraphqlBatchUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response graphqlBatchUnprocessableEntity
*/
type GraphqlBatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGraphqlBatchUnprocessableEntity creates GraphqlBatchUnprocessableEntity with default headers values
func NewGraphqlBatchUnprocessableEntity() *GraphqlBatchUnprocessableEntity {

	return &GraphqlBatchUnprocessableEntity{}
}

// WithPayload adds the payload to the graphql batch unprocessable entity response
func (o *GraphqlBatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *GraphqlBatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the graphql batch unprocessable entity response
func (o *GraphqlBatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GraphqlBatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GraphqlBatchInternalServerErrorCode is the HTTP code returned for type GraphqlBatchInternalServerError
const GraphqlBatchInternalServerErrorCode int = 500

/*
GraphqlBatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response graphqlBatchInternalServerError
*/
type GraphqlBatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGraphqlBatchInternalServerError creates GraphqlBatchInternalServerError with default headers values
func NewGraphqlBatchInternalServerError() *GraphqlBatchInternalServerError {

	return &GraphqlBatchInternalServerError{}
}

// WithPayload adds the payload to the graphql batch internal server error response
func (o *GraphqlBatchInternalServerError) WithPayload(payload *models.ErrorResponse) *GraphqlBatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the graphql batch internal server error response
func (o *GraphqlBatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GraphqlBatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
