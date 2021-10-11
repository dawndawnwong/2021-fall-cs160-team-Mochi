// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateGroupV1HandlerFunc turns a function with the right signature into a create group v1 handler
type CreateGroupV1HandlerFunc func(CreateGroupV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateGroupV1HandlerFunc) Handle(params CreateGroupV1Params) middleware.Responder {
	return fn(params)
}

// CreateGroupV1Handler interface for that can handle valid create group v1 params
type CreateGroupV1Handler interface {
	Handle(CreateGroupV1Params) middleware.Responder
}

// NewCreateGroupV1 creates a new http.Handler for the create group v1 operation
func NewCreateGroupV1(ctx *middleware.Context, handler CreateGroupV1Handler) *CreateGroupV1 {
	return &CreateGroupV1{Context: ctx, Handler: handler}
}

/* CreateGroupV1 swagger:route POST /v1/groups groupsV1 createGroupV1

create a group

*/
type CreateGroupV1 struct {
	Context *middleware.Context
	Handler CreateGroupV1Handler
}

func (o *CreateGroupV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateGroupV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}