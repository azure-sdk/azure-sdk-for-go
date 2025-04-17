// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagementClient contains the methods for the PeeringManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckServiceProviderAvailability - Checks if the peering service provider is present within 1000 miles of customer's location
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - checkServiceProviderAvailabilityInput - The CheckServiceProviderAvailabilityInput indicating customer location and service
//     provider.
//   - options - ManagementClientCheckServiceProviderAvailabilityOptions contains the optional parameters for the ManagementClient.CheckServiceProviderAvailability
//     method.
func (client *ManagementClient) CheckServiceProviderAvailability(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput, options *ManagementClientCheckServiceProviderAvailabilityOptions) (ManagementClientCheckServiceProviderAvailabilityResponse, error) {
	var err error
	const operationName = "ManagementClient.CheckServiceProviderAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkServiceProviderAvailabilityCreateRequest(ctx, checkServiceProviderAvailabilityInput, options)
	if err != nil {
		return ManagementClientCheckServiceProviderAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckServiceProviderAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementClientCheckServiceProviderAvailabilityResponse{}, err
	}
	resp, err := client.checkServiceProviderAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkServiceProviderAvailabilityCreateRequest creates the CheckServiceProviderAvailability request.
func (client *ManagementClient) checkServiceProviderAvailabilityCreateRequest(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput, _ *ManagementClientCheckServiceProviderAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/checkServiceProviderAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkServiceProviderAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkServiceProviderAvailabilityHandleResponse handles the CheckServiceProviderAvailability response.
func (client *ManagementClient) checkServiceProviderAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckServiceProviderAvailabilityResponse, error) {
	result := ManagementClientCheckServiceProviderAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return ManagementClientCheckServiceProviderAvailabilityResponse{}, err
	}
	return result, nil
}
