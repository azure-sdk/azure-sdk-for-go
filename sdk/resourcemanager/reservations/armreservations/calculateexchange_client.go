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

// CalculateExchangeClient contains the methods for the CalculateExchange group.
// Don't use this type directly, use NewCalculateExchangeClient() instead.
type CalculateExchangeClient struct {
	host string
	pl   runtime.Pipeline
}

// NewCalculateExchangeClient creates a new instance of CalculateExchangeClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCalculateExchangeClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CalculateExchangeClient, error) {
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
	client := &CalculateExchangeClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginPost - Calculates price for exchanging Reservations if there are no policy errors.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - body - Request containing purchases and refunds that need to be executed.
//   - options - CalculateExchangeClientBeginPostOptions contains the optional parameters for the CalculateExchangeClient.BeginPost
//     method.
func (client *CalculateExchangeClient) BeginPost(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeClientBeginPostOptions) (*runtime.Poller[CalculateExchangeClientPostResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post(ctx, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[CalculateExchangeClientPostResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[CalculateExchangeClientPostResponse](options.ResumeToken, client.pl, nil)
	}
}

// Post - Calculates price for exchanging Reservations if there are no policy errors.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *CalculateExchangeClient) post(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeClientBeginPostOptions) (*http.Response, error) {
	req, err := client.postCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// postCreateRequest creates the Post request.
func (client *CalculateExchangeClient) postCreateRequest(ctx context.Context, body CalculateExchangeRequest, options *CalculateExchangeClientBeginPostOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/calculateExchange"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
