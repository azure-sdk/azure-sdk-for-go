//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armreservations

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ReturnClient contains the methods for the Return group.
// Don't use this type directly, use NewReturnClient() instead.
type ReturnClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReturnClient creates a new instance of ReturnClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReturnClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReturnClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReturnClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Post - Return a reservation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// body - Information needed for returning reservation.
// options - ReturnClientPostOptions contains the optional parameters for the ReturnClient.Post method.
func (client *ReturnClient) Post(ctx context.Context, body RefundRequest, options *ReturnClientPostOptions) (ReturnClientPostResponse, error) {
	req, err := client.postCreateRequest(ctx, body, options)
	if err != nil {
		return ReturnClientPostResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReturnClientPostResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ReturnClientPostResponse{}, runtime.NewResponseError(resp)
	}
	return client.postHandleResponse(resp)
}

// postCreateRequest creates the Post request.
func (client *ReturnClient) postCreateRequest(ctx context.Context, body RefundRequest, options *ReturnClientPostOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/return"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// postHandleResponse handles the Post response.
func (client *ReturnClient) postHandleResponse(resp *http.Response) (ReturnClientPostResponse, error) {
	result := ReturnClientPostResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.RefundResponse); err != nil {
		return ReturnClientPostResponse{}, err
	}
	return result, nil
}
