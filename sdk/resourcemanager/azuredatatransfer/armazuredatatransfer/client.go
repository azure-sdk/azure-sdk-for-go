//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredatatransfer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Client contains the methods for the AzureDataTransfer group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal *arm.Client
}

// NewClient creates a new instance of Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		internal: cl,
	}
	return client, nil
}

// ListApprovedSchemas - Lists approved schemas for Azure Data Transfer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-11-preview
//   - pipeline - The request to list approved schemas.
//   - options - ClientListApprovedSchemasOptions contains the optional parameters for the Client.ListApprovedSchemas method.
func (client *Client) ListApprovedSchemas(ctx context.Context, pipeline ListApprovedSchemasRequest, options *ClientListApprovedSchemasOptions) (ClientListApprovedSchemasResponse, error) {
	var err error
	const operationName = "Client.ListApprovedSchemas"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listApprovedSchemasCreateRequest(ctx, pipeline, options)
	if err != nil {
		return ClientListApprovedSchemasResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientListApprovedSchemasResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientListApprovedSchemasResponse{}, err
	}
	resp, err := client.listApprovedSchemasHandleResponse(httpResp)
	return resp, err
}

// listApprovedSchemasCreateRequest creates the ListApprovedSchemas request.
func (client *Client) listApprovedSchemasCreateRequest(ctx context.Context, pipeline ListApprovedSchemasRequest, options *ClientListApprovedSchemasOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AzureDataTransfer/listApprovedSchemas"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pipeline); err != nil {
		return nil, err
	}
	return req, nil
}

// listApprovedSchemasHandleResponse handles the ListApprovedSchemas response.
func (client *Client) listApprovedSchemasHandleResponse(resp *http.Response) (ClientListApprovedSchemasResponse, error) {
	result := ClientListApprovedSchemasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SchemasListResult); err != nil {
		return ClientListApprovedSchemasResponse{}, err
	}
	return result, nil
}

// ValidateSchema - Validates a schema for Azure Data Transfer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-11-preview
//   - schema - The schema to validate
//   - options - ClientValidateSchemaOptions contains the optional parameters for the Client.ValidateSchema method.
func (client *Client) ValidateSchema(ctx context.Context, schema Schema, options *ClientValidateSchemaOptions) (ClientValidateSchemaResponse, error) {
	var err error
	const operationName = "Client.ValidateSchema"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateSchemaCreateRequest(ctx, schema, options)
	if err != nil {
		return ClientValidateSchemaResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientValidateSchemaResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientValidateSchemaResponse{}, err
	}
	resp, err := client.validateSchemaHandleResponse(httpResp)
	return resp, err
}

// validateSchemaCreateRequest creates the ValidateSchema request.
func (client *Client) validateSchemaCreateRequest(ctx context.Context, schema Schema, options *ClientValidateSchemaOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AzureDataTransfer/validateSchema"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, schema); err != nil {
		return nil, err
	}
	return req, nil
}

// validateSchemaHandleResponse handles the ValidateSchema response.
func (client *Client) validateSchemaHandleResponse(resp *http.Response) (ClientValidateSchemaResponse, error) {
	result := ClientValidateSchemaResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateSchemaResult); err != nil {
		return ClientValidateSchemaResponse{}, err
	}
	return result, nil
}
