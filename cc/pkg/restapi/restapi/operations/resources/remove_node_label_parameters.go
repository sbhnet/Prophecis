// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveNodeLabelParams creates a new RemoveNodeLabelParams object
// no default values defined in spec.
func NewRemoveNodeLabelParams() RemoveNodeLabelParams {

	return RemoveNodeLabelParams{}
}

// RemoveNodeLabelParams contains all the bound params for the remove node label operation
// typically these are obtained from a http.Request
//
// swagger:parameters RemoveNodeLabel
type RemoveNodeLabelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*label.
	  Required: true
	  In: path
	*/
	Label string
	/*node name.
	  Required: true
	  In: path
	*/
	NodeName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRemoveNodeLabelParams() beforehand.
func (o *RemoveNodeLabelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rLabel, rhkLabel, _ := route.Params.GetOK("label")
	if err := o.bindLabel(rLabel, rhkLabel, route.Formats); err != nil {
		res = append(res, err)
	}

	rNodeName, rhkNodeName, _ := route.Params.GetOK("nodeName")
	if err := o.bindNodeName(rNodeName, rhkNodeName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLabel binds and validates parameter Label from path.
func (o *RemoveNodeLabelParams) bindLabel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Label = raw

	return nil
}

// bindNodeName binds and validates parameter NodeName from path.
func (o *RemoveNodeLabelParams) bindNodeName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.NodeName = raw

	return nil
}
