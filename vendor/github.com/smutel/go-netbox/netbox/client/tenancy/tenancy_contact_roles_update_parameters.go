// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package tenancy

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
	"github.com/go-openapi/swag"

	"github.com/smutel/go-netbox/netbox/models"
)

// NewTenancyContactRolesUpdateParams creates a new TenancyContactRolesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactRolesUpdateParams() *TenancyContactRolesUpdateParams {
	return &TenancyContactRolesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactRolesUpdateParamsWithTimeout creates a new TenancyContactRolesUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactRolesUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactRolesUpdateParams {
	return &TenancyContactRolesUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactRolesUpdateParamsWithContext creates a new TenancyContactRolesUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactRolesUpdateParamsWithContext(ctx context.Context) *TenancyContactRolesUpdateParams {
	return &TenancyContactRolesUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactRolesUpdateParamsWithHTTPClient creates a new TenancyContactRolesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactRolesUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactRolesUpdateParams {
	return &TenancyContactRolesUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyContactRolesUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy contact roles update operation.

   Typically these are written to a http.Request.
*/
type TenancyContactRolesUpdateParams struct {

	// Data.
	Data *models.ContactRole

	/* ID.

	   A unique integer value identifying this contact role.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactRolesUpdateParams) WithDefaults() *TenancyContactRolesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact roles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactRolesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) WithContext(ctx context.Context) *TenancyContactRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) WithData(data *models.ContactRole) *TenancyContactRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) SetData(data *models.ContactRole) {
	o.Data = data
}

// WithID adds the id to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) WithID(id int64) *TenancyContactRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy contact roles update params
func (o *TenancyContactRolesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
