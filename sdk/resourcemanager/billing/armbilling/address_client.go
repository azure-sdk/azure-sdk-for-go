// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AddressClient contains the methods for the Address group.
// Don't use this type directly, use NewAddressClient() instead.
type AddressClient struct {
	internal *arm.Client
}

// NewAddressClient creates a new instance of AddressClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAddressClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AddressClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AddressClient{
		internal: cl,
	}
	return client, nil
}

// Validate - Validates an address. Use the operation to validate an address before using it as soldTo or a billTo address.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - parameters - Address details.
//   - options - AddressClientValidateOptions contains the optional parameters for the AddressClient.Validate method.
func (client *AddressClient) Validate(ctx context.Context, parameters AddressDetails, options *AddressClientValidateOptions) (AddressClientValidateResponse, error) {
	var err error
	const operationName = "AddressClient.Validate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateCreateRequest(ctx, parameters, options)
	if err != nil {
		return AddressClientValidateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AddressClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AddressClientValidateResponse{}, err
	}
	resp, err := client.validateHandleResponse(httpResp)
	return resp, err
}

// validateCreateRequest creates the Validate request.
func (client *AddressClient) validateCreateRequest(ctx context.Context, parameters AddressDetails, _ *AddressClientValidateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/validateAddress"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateHandleResponse handles the Validate response.
func (client *AddressClient) validateHandleResponse(resp *http.Response) (AddressClientValidateResponse, error) {
	result := AddressClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddressValidationResponse); err != nil {
		return AddressClientValidateResponse{}, err
	}
	return result, nil
}
