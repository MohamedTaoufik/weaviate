//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewObjectsHeadParams creates a new ObjectsHeadParams object
// with the default values initialized.
func NewObjectsHeadParams() *ObjectsHeadParams {
	var ()
	return &ObjectsHeadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsHeadParamsWithTimeout creates a new ObjectsHeadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObjectsHeadParamsWithTimeout(timeout time.Duration) *ObjectsHeadParams {
	var ()
	return &ObjectsHeadParams{

		timeout: timeout,
	}
}

// NewObjectsHeadParamsWithContext creates a new ObjectsHeadParams object
// with the default values initialized, and the ability to set a context for a request
func NewObjectsHeadParamsWithContext(ctx context.Context) *ObjectsHeadParams {
	var ()
	return &ObjectsHeadParams{

		Context: ctx,
	}
}

// NewObjectsHeadParamsWithHTTPClient creates a new ObjectsHeadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObjectsHeadParamsWithHTTPClient(client *http.Client) *ObjectsHeadParams {
	var ()
	return &ObjectsHeadParams{
		HTTPClient: client,
	}
}

/*ObjectsHeadParams contains all the parameters to send to the API endpoint
for the objects head operation typically these are written to a http.Request
*/
type ObjectsHeadParams struct {

	/*ID
	  Unique ID of the Object.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the objects head params
func (o *ObjectsHeadParams) WithTimeout(timeout time.Duration) *ObjectsHeadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects head params
func (o *ObjectsHeadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects head params
func (o *ObjectsHeadParams) WithContext(ctx context.Context) *ObjectsHeadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects head params
func (o *ObjectsHeadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects head params
func (o *ObjectsHeadParams) WithHTTPClient(client *http.Client) *ObjectsHeadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects head params
func (o *ObjectsHeadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the objects head params
func (o *ObjectsHeadParams) WithID(id strfmt.UUID) *ObjectsHeadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the objects head params
func (o *ObjectsHeadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsHeadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}