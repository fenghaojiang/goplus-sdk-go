// Code generated by go-swagger; DO NOT EDIT.

package approve_controller_v_1

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

// NewApprovalContractUsingGETParams creates a new ApprovalContractUsingGETParams object
// with the default values initialized.
func NewApprovalContractUsingGETParams() *ApprovalContractUsingGETParams {
	var ()
	return &ApprovalContractUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApprovalContractUsingGETParamsWithTimeout creates a new ApprovalContractUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApprovalContractUsingGETParamsWithTimeout(timeout time.Duration) *ApprovalContractUsingGETParams {
	var ()
	return &ApprovalContractUsingGETParams{

		timeout: timeout,
	}
}

// NewApprovalContractUsingGETParamsWithContext creates a new ApprovalContractUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewApprovalContractUsingGETParamsWithContext(ctx context.Context) *ApprovalContractUsingGETParams {
	var ()
	return &ApprovalContractUsingGETParams{

		Context: ctx,
	}
}

// NewApprovalContractUsingGETParamsWithHTTPClient creates a new ApprovalContractUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApprovalContractUsingGETParamsWithHTTPClient(client *http.Client) *ApprovalContractUsingGETParams {
	var ()
	return &ApprovalContractUsingGETParams{
		HTTPClient: client,
	}
}

/*
ApprovalContractUsingGETParams contains all the parameters to send to the API endpoint
for the approval contract using g e t operation typically these are written to a http.Request
*/
type ApprovalContractUsingGETParams struct {

	/*Authorization
	  Authorization (test：Bearer 81|9ihH8JzEuFu4MQ9DjWmH5WrNCPW1zQ9cCv8WrbB1)

	*/
	Authorization *string
	/*ChainID
	  Chain id, (ETH: 1,  BSC: 56, OKC: 66, Heco: 128, Polygon: 137, Fantom:250, Arbitrum: 42161, Avalanche: 43114)

	*/
	ChainID string
	/*ContractAddresses
	  Contract needs to be detected

	*/
	ContractAddresses *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithTimeout(timeout time.Duration) *ApprovalContractUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithContext(ctx context.Context) *ApprovalContractUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithHTTPClient(client *http.Client) *ApprovalContractUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithAuthorization(authorization *string) *ApprovalContractUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithChainID adds the chainID to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithChainID(chainID string) *ApprovalContractUsingGETParams {
	o.SetChainID(chainID)
	return o
}

// SetChainID adds the chainId to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetChainID(chainID string) {
	o.ChainID = chainID
}

// WithContractAddresses adds the contractAddresses to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) WithContractAddresses(contractAddresses *string) *ApprovalContractUsingGETParams {
	o.SetContractAddresses(contractAddresses)
	return o
}

// SetContractAddresses adds the contractAddresses to the approval contract using g e t params
func (o *ApprovalContractUsingGETParams) SetContractAddresses(contractAddresses *string) {
	o.ContractAddresses = contractAddresses
}

// WriteToRequest writes these params to a swagger request
func (o *ApprovalContractUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// path param chain_id
	if err := r.SetPathParam("chain_id", o.ChainID); err != nil {
		return err
	}

	if o.ContractAddresses != nil {

		// query param contract_addresses
		var qrContractAddresses string
		if o.ContractAddresses != nil {
			qrContractAddresses = *o.ContractAddresses
		}
		qContractAddresses := qrContractAddresses
		if qContractAddresses != "" {
			if err := r.SetQueryParam("contract_addresses", qContractAddresses); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
