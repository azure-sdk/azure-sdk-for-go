//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn

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

// ManagementClient contains the methods for the CdnManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagementClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEndpointNameAvailability - Check the availability of a resource name. This is needed for resources where name is globally
// unique, such as a afdx endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - checkEndpointNameAvailabilityInput - Input to check.
//   - options - ManagementClientCheckEndpointNameAvailabilityOptions contains the optional parameters for the ManagementClient.CheckEndpointNameAvailability
//     method.
func (client *ManagementClient) CheckEndpointNameAvailability(ctx context.Context, resourceGroupName string, checkEndpointNameAvailabilityInput CheckEndpointNameAvailabilityInput, options *ManagementClientCheckEndpointNameAvailabilityOptions) (ManagementClientCheckEndpointNameAvailabilityResponse, error) {
	req, err := client.checkEndpointNameAvailabilityCreateRequest(ctx, resourceGroupName, checkEndpointNameAvailabilityInput, options)
	if err != nil {
		return ManagementClientCheckEndpointNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckEndpointNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientCheckEndpointNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkEndpointNameAvailabilityHandleResponse(resp)
}

// checkEndpointNameAvailabilityCreateRequest creates the CheckEndpointNameAvailability request.
func (client *ManagementClient) checkEndpointNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, checkEndpointNameAvailabilityInput CheckEndpointNameAvailabilityInput, options *ManagementClientCheckEndpointNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/checkEndpointNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkEndpointNameAvailabilityInput)
}

// checkEndpointNameAvailabilityHandleResponse handles the CheckEndpointNameAvailability response.
func (client *ManagementClient) checkEndpointNameAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckEndpointNameAvailabilityResponse, error) {
	result := ManagementClientCheckEndpointNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckEndpointNameAvailabilityOutput); err != nil {
		return ManagementClientCheckEndpointNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - Check the availability of a resource name. This is needed for resources where name is globally
// unique, such as a CDN endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - checkNameAvailabilityInput - Input to check.
//   - options - ManagementClientCheckNameAvailabilityOptions contains the optional parameters for the ManagementClient.CheckNameAvailability
//     method.
func (client *ManagementClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *ManagementClientCheckNameAvailabilityOptions) (ManagementClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityInput, options)
	if err != nil {
		return ManagementClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ManagementClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *ManagementClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Cdn/checkNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityInput)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ManagementClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckNameAvailabilityResponse, error) {
	result := ManagementClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return ManagementClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckNameAvailabilityWithSubscription - Check the availability of a resource name. This is needed for resources where name
// is globally unique, such as a CDN endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - checkNameAvailabilityInput - Input to check.
//   - options - ManagementClientCheckNameAvailabilityWithSubscriptionOptions contains the optional parameters for the ManagementClient.CheckNameAvailabilityWithSubscription
//     method.
func (client *ManagementClient) CheckNameAvailabilityWithSubscription(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *ManagementClientCheckNameAvailabilityWithSubscriptionOptions) (ManagementClientCheckNameAvailabilityWithSubscriptionResponse, error) {
	req, err := client.checkNameAvailabilityWithSubscriptionCreateRequest(ctx, checkNameAvailabilityInput, options)
	if err != nil {
		return ManagementClientCheckNameAvailabilityWithSubscriptionResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckNameAvailabilityWithSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientCheckNameAvailabilityWithSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityWithSubscriptionHandleResponse(resp)
}

// checkNameAvailabilityWithSubscriptionCreateRequest creates the CheckNameAvailabilityWithSubscription request.
func (client *ManagementClient) checkNameAvailabilityWithSubscriptionCreateRequest(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *ManagementClientCheckNameAvailabilityWithSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityInput)
}

// checkNameAvailabilityWithSubscriptionHandleResponse handles the CheckNameAvailabilityWithSubscription response.
func (client *ManagementClient) checkNameAvailabilityWithSubscriptionHandleResponse(resp *http.Response) (ManagementClientCheckNameAvailabilityWithSubscriptionResponse, error) {
	result := ManagementClientCheckNameAvailabilityWithSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return ManagementClientCheckNameAvailabilityWithSubscriptionResponse{}, err
	}
	return result, nil
}

// ValidateProbe - Check if the probe path is a valid path and the file can be accessed. Probe path is the path to a file
// hosted on the origin server to help accelerate the delivery of dynamic content via the CDN
// endpoint. This path is relative to the origin path specified in the endpoint configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - validateProbeInput - Input to check.
//   - options - ManagementClientValidateProbeOptions contains the optional parameters for the ManagementClient.ValidateProbe
//     method.
func (client *ManagementClient) ValidateProbe(ctx context.Context, validateProbeInput ValidateProbeInput, options *ManagementClientValidateProbeOptions) (ManagementClientValidateProbeResponse, error) {
	req, err := client.validateProbeCreateRequest(ctx, validateProbeInput, options)
	if err != nil {
		return ManagementClientValidateProbeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientValidateProbeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientValidateProbeResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateProbeHandleResponse(resp)
}

// validateProbeCreateRequest creates the ValidateProbe request.
func (client *ManagementClient) validateProbeCreateRequest(ctx context.Context, validateProbeInput ValidateProbeInput, options *ManagementClientValidateProbeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/validateProbe"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, validateProbeInput)
}

// validateProbeHandleResponse handles the ValidateProbe response.
func (client *ManagementClient) validateProbeHandleResponse(resp *http.Response) (ManagementClientValidateProbeResponse, error) {
	result := ManagementClientValidateProbeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateProbeOutput); err != nil {
		return ManagementClientValidateProbeResponse{}, err
	}
	return result, nil
}
