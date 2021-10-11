// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// AddGroupUsersV1OKCode is the HTTP code returned for type AddGroupUsersV1OK
const AddGroupUsersV1OKCode int = 200

/*AddGroupUsersV1OK Success

swagger:response addGroupUsersV1OK
*/
type AddGroupUsersV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.GroupResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1OK creates AddGroupUsersV1OK with default headers values
func NewAddGroupUsersV1OK() *AddGroupUsersV1OK {

	return &AddGroupUsersV1OK{}
}

// WithPayload adds the payload to the add group users v1 o k response
func (o *AddGroupUsersV1OK) WithPayload(payload *models.GroupResponse) *AddGroupUsersV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 o k response
func (o *AddGroupUsersV1OK) SetPayload(payload *models.GroupResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1BadRequestCode is the HTTP code returned for type AddGroupUsersV1BadRequest
const AddGroupUsersV1BadRequestCode int = 400

/*AddGroupUsersV1BadRequest Bad Request

swagger:response addGroupUsersV1BadRequest
*/
type AddGroupUsersV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1BadRequest creates AddGroupUsersV1BadRequest with default headers values
func NewAddGroupUsersV1BadRequest() *AddGroupUsersV1BadRequest {

	return &AddGroupUsersV1BadRequest{}
}

// WithPayload adds the payload to the add group users v1 bad request response
func (o *AddGroupUsersV1BadRequest) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 bad request response
func (o *AddGroupUsersV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1UnauthorizedCode is the HTTP code returned for type AddGroupUsersV1Unauthorized
const AddGroupUsersV1UnauthorizedCode int = 401

/*AddGroupUsersV1Unauthorized Unauthorized

swagger:response addGroupUsersV1Unauthorized
*/
type AddGroupUsersV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1Unauthorized creates AddGroupUsersV1Unauthorized with default headers values
func NewAddGroupUsersV1Unauthorized() *AddGroupUsersV1Unauthorized {

	return &AddGroupUsersV1Unauthorized{}
}

// WithPayload adds the payload to the add group users v1 unauthorized response
func (o *AddGroupUsersV1Unauthorized) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 unauthorized response
func (o *AddGroupUsersV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1ForbiddenCode is the HTTP code returned for type AddGroupUsersV1Forbidden
const AddGroupUsersV1ForbiddenCode int = 403

/*AddGroupUsersV1Forbidden Forbidden

swagger:response addGroupUsersV1Forbidden
*/
type AddGroupUsersV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1Forbidden creates AddGroupUsersV1Forbidden with default headers values
func NewAddGroupUsersV1Forbidden() *AddGroupUsersV1Forbidden {

	return &AddGroupUsersV1Forbidden{}
}

// WithPayload adds the payload to the add group users v1 forbidden response
func (o *AddGroupUsersV1Forbidden) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 forbidden response
func (o *AddGroupUsersV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1NotFoundCode is the HTTP code returned for type AddGroupUsersV1NotFound
const AddGroupUsersV1NotFoundCode int = 404

/*AddGroupUsersV1NotFound Not Found

swagger:response addGroupUsersV1NotFound
*/
type AddGroupUsersV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1NotFound creates AddGroupUsersV1NotFound with default headers values
func NewAddGroupUsersV1NotFound() *AddGroupUsersV1NotFound {

	return &AddGroupUsersV1NotFound{}
}

// WithPayload adds the payload to the add group users v1 not found response
func (o *AddGroupUsersV1NotFound) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 not found response
func (o *AddGroupUsersV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1ConflictCode is the HTTP code returned for type AddGroupUsersV1Conflict
const AddGroupUsersV1ConflictCode int = 409

/*AddGroupUsersV1Conflict Conflict

swagger:response addGroupUsersV1Conflict
*/
type AddGroupUsersV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1Conflict creates AddGroupUsersV1Conflict with default headers values
func NewAddGroupUsersV1Conflict() *AddGroupUsersV1Conflict {

	return &AddGroupUsersV1Conflict{}
}

// WithPayload adds the payload to the add group users v1 conflict response
func (o *AddGroupUsersV1Conflict) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 conflict response
func (o *AddGroupUsersV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddGroupUsersV1InternalServerErrorCode is the HTTP code returned for type AddGroupUsersV1InternalServerError
const AddGroupUsersV1InternalServerErrorCode int = 500

/*AddGroupUsersV1InternalServerError Internal Server Error

swagger:response addGroupUsersV1InternalServerError
*/
type AddGroupUsersV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewAddGroupUsersV1InternalServerError creates AddGroupUsersV1InternalServerError with default headers values
func NewAddGroupUsersV1InternalServerError() *AddGroupUsersV1InternalServerError {

	return &AddGroupUsersV1InternalServerError{}
}

// WithPayload adds the payload to the add group users v1 internal server error response
func (o *AddGroupUsersV1InternalServerError) WithPayload(payload *models.ErrResponse) *AddGroupUsersV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add group users v1 internal server error response
func (o *AddGroupUsersV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddGroupUsersV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}