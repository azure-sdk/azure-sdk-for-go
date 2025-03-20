// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// AFDProfilesClient contains the methods for the AFDProfiles group.
// Don't use this type directly, use NewAFDProfilesClient() instead.
type AFDProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAFDProfilesClient creates a new instance of AFDProfilesClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAFDProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AFDProfilesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AFDProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEndpointNameAvailability - Check the availability of an afdx endpoint name, and return the globally unique endpoint
// host name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - checkEndpointNameAvailabilityInput - Input to check.
//   - options - AFDProfilesClientCheckEndpointNameAvailabilityOptions contains the optional parameters for the AFDProfilesClient.CheckEndpointNameAvailability
//     method.
func (client *AFDProfilesClient) CheckEndpointNameAvailability(ctx context.Context, resourceGroupName string, profileName string, checkEndpointNameAvailabilityInput CheckEndpointNameAvailabilityInput, options *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error) {
	var err error
	const operationName = "AFDProfilesClient.CheckEndpointNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkEndpointNameAvailabilityCreateRequest(ctx, resourceGroupName, profileName, checkEndpointNameAvailabilityInput, options)
	if err != nil {
		return AFDProfilesClientCheckEndpointNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AFDProfilesClientCheckEndpointNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AFDProfilesClientCheckEndpointNameAvailabilityResponse{}, err
	}
	resp, err := client.checkEndpointNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkEndpointNameAvailabilityCreateRequest creates the CheckEndpointNameAvailability request.
func (client *AFDProfilesClient) checkEndpointNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, profileName string, checkEndpointNameAvailabilityInput CheckEndpointNameAvailabilityInput, _ *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/checkEndpointNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkEndpointNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkEndpointNameAvailabilityHandleResponse handles the CheckEndpointNameAvailability response.
func (client *AFDProfilesClient) checkEndpointNameAvailabilityHandleResponse(resp *http.Response) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error) {
	result := AFDProfilesClientCheckEndpointNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckEndpointNameAvailabilityOutput); err != nil {
		return AFDProfilesClientCheckEndpointNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckHostNameAvailability - Validates the custom domain mapping to ensure it maps to the correct Azure Front Door endpoint
// in DNS.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - checkHostNameAvailabilityInput - Custom domain to be validated.
//   - options - AFDProfilesClientCheckHostNameAvailabilityOptions contains the optional parameters for the AFDProfilesClient.CheckHostNameAvailability
//     method.
func (client *AFDProfilesClient) CheckHostNameAvailability(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput CheckHostNameAvailabilityInput, options *AFDProfilesClientCheckHostNameAvailabilityOptions) (AFDProfilesClientCheckHostNameAvailabilityResponse, error) {
	var err error
	const operationName = "AFDProfilesClient.CheckHostNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkHostNameAvailabilityCreateRequest(ctx, resourceGroupName, profileName, checkHostNameAvailabilityInput, options)
	if err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	resp, err := client.checkHostNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkHostNameAvailabilityCreateRequest creates the CheckHostNameAvailability request.
func (client *AFDProfilesClient) checkHostNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput CheckHostNameAvailabilityInput, _ *AFDProfilesClientCheckHostNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/checkHostNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkHostNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkHostNameAvailabilityHandleResponse handles the CheckHostNameAvailability response.
func (client *AFDProfilesClient) checkHostNameAvailabilityHandleResponse(resp *http.Response) (AFDProfilesClientCheckHostNameAvailabilityResponse, error) {
	result := AFDProfilesClientCheckHostNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return AFDProfilesClientCheckHostNameAvailabilityResponse{}, err
	}
	return result, nil
}

// NewListResourceUsagePager - Checks the quota and actual usage of endpoints under the given Azure Front Door profile.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - options - AFDProfilesClientListResourceUsageOptions contains the optional parameters for the AFDProfilesClient.NewListResourceUsagePager
//     method.
func (client *AFDProfilesClient) NewListResourceUsagePager(resourceGroupName string, profileName string, options *AFDProfilesClientListResourceUsageOptions) *runtime.Pager[AFDProfilesClientListResourceUsageResponse] {
	return runtime.NewPager(runtime.PagingHandler[AFDProfilesClientListResourceUsageResponse]{
		More: func(page AFDProfilesClientListResourceUsageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AFDProfilesClientListResourceUsageResponse) (AFDProfilesClientListResourceUsageResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AFDProfilesClient.NewListResourceUsagePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listResourceUsageCreateRequest(ctx, resourceGroupName, profileName, options)
			}, nil)
			if err != nil {
				return AFDProfilesClientListResourceUsageResponse{}, err
			}
			return client.listResourceUsageHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listResourceUsageCreateRequest creates the ListResourceUsage request.
func (client *AFDProfilesClient) listResourceUsageCreateRequest(ctx context.Context, resourceGroupName string, profileName string, _ *AFDProfilesClientListResourceUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listResourceUsageHandleResponse handles the ListResourceUsage response.
func (client *AFDProfilesClient) listResourceUsageHandleResponse(resp *http.Response) (AFDProfilesClientListResourceUsageResponse, error) {
	result := AFDProfilesClientListResourceUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return AFDProfilesClientListResourceUsageResponse{}, err
	}
	return result, nil
}

// BeginUpgrade - Upgrade a profile from StandardAzureFrontDoor to PremiumAzureFrontDoor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - profileUpgradeParameters - Profile upgrade input parameter.
//   - options - AFDProfilesClientBeginUpgradeOptions contains the optional parameters for the AFDProfilesClient.BeginUpgrade
//     method.
func (client *AFDProfilesClient) BeginUpgrade(ctx context.Context, resourceGroupName string, profileName string, profileUpgradeParameters ProfileUpgradeParameters, options *AFDProfilesClientBeginUpgradeOptions) (*runtime.Poller[AFDProfilesClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, profileName, profileUpgradeParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDProfilesClientUpgradeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AFDProfilesClientUpgradeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Upgrade - Upgrade a profile from StandardAzureFrontDoor to PremiumAzureFrontDoor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
func (client *AFDProfilesClient) upgrade(ctx context.Context, resourceGroupName string, profileName string, profileUpgradeParameters ProfileUpgradeParameters, options *AFDProfilesClientBeginUpgradeOptions) (*http.Response, error) {
	var err error
	const operationName = "AFDProfilesClient.BeginUpgrade"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, profileName, profileUpgradeParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *AFDProfilesClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, profileName string, profileUpgradeParameters ProfileUpgradeParameters, _ *AFDProfilesClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/upgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, profileUpgradeParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// ValidateSecret - Validate a Secret in the profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium which is unique within the resource group.
//   - validateSecretInput - The Secret source.
//   - options - AFDProfilesClientValidateSecretOptions contains the optional parameters for the AFDProfilesClient.ValidateSecret
//     method.
func (client *AFDProfilesClient) ValidateSecret(ctx context.Context, resourceGroupName string, profileName string, validateSecretInput ValidateSecretInput, options *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error) {
	var err error
	const operationName = "AFDProfilesClient.ValidateSecret"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateSecretCreateRequest(ctx, resourceGroupName, profileName, validateSecretInput, options)
	if err != nil {
		return AFDProfilesClientValidateSecretResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AFDProfilesClientValidateSecretResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AFDProfilesClientValidateSecretResponse{}, err
	}
	resp, err := client.validateSecretHandleResponse(httpResp)
	return resp, err
}

// validateSecretCreateRequest creates the ValidateSecret request.
func (client *AFDProfilesClient) validateSecretCreateRequest(ctx context.Context, resourceGroupName string, profileName string, validateSecretInput ValidateSecretInput, _ *AFDProfilesClientValidateSecretOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/validateSecret"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, validateSecretInput); err != nil {
		return nil, err
	}
	return req, nil
}

// validateSecretHandleResponse handles the ValidateSecret response.
func (client *AFDProfilesClient) validateSecretHandleResponse(resp *http.Response) (AFDProfilesClientValidateSecretResponse, error) {
	result := AFDProfilesClientValidateSecretResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateSecretOutput); err != nil {
		return AFDProfilesClientValidateSecretResponse{}, err
	}
	return result, nil
}
