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

// AFDCustomDomainsClient contains the methods for the AFDCustomDomains group.
// Don't use this type directly, use NewAFDCustomDomainsClient() instead.
type AFDCustomDomainsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAFDCustomDomainsClient creates a new instance of AFDCustomDomainsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAFDCustomDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AFDCustomDomainsClient, error) {
	cl, err := arm.NewClient(moduleName+".AFDCustomDomainsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AFDCustomDomainsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new domain within the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - customDomainName - Name of the domain under the profile which is unique globally
//   - customDomain - Domain properties
//   - options - AFDCustomDomainsClientBeginCreateOptions contains the optional parameters for the AFDCustomDomainsClient.BeginCreate
//     method.
func (client *AFDCustomDomainsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsClientBeginCreateOptions) (*runtime.Poller[AFDCustomDomainsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, customDomainName, customDomain, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDCustomDomainsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDCustomDomainsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a new domain within the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *AFDCustomDomainsClient) create(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, customDomainName, customDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *AFDCustomDomainsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomain AFDDomain, options *AFDCustomDomainsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, customDomain)
}

// BeginDelete - Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription,
// resource group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - customDomainName - Name of the domain under the profile which is unique globally.
//   - options - AFDCustomDomainsClientBeginDeleteOptions contains the optional parameters for the AFDCustomDomainsClient.BeginDelete
//     method.
func (client *AFDCustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginDeleteOptions) (*runtime.Poller[AFDCustomDomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, customDomainName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDCustomDomainsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDCustomDomainsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource
// group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *AFDCustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AFDCustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing AzureFrontDoor domain with the specified domain name under the specified subscription, resource
// group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - customDomainName - Name of the domain under the profile which is unique globally.
//   - options - AFDCustomDomainsClientGetOptions contains the optional parameters for the AFDCustomDomainsClient.Get method.
func (client *AFDCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientGetOptions) (AFDCustomDomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return AFDCustomDomainsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AFDCustomDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDCustomDomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AFDCustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AFDCustomDomainsClient) getHandleResponse(resp *http.Response) (AFDCustomDomainsClientGetResponse, error) {
	result := AFDCustomDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDDomain); err != nil {
		return AFDCustomDomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProfilePager - Lists existing AzureFrontDoor domains.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - options - AFDCustomDomainsClientListByProfileOptions contains the optional parameters for the AFDCustomDomainsClient.NewListByProfilePager
//     method.
func (client *AFDCustomDomainsClient) NewListByProfilePager(resourceGroupName string, profileName string, options *AFDCustomDomainsClientListByProfileOptions) *runtime.Pager[AFDCustomDomainsClientListByProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[AFDCustomDomainsClientListByProfileResponse]{
		More: func(page AFDCustomDomainsClientListByProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AFDCustomDomainsClientListByProfileResponse) (AFDCustomDomainsClientListByProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AFDCustomDomainsClientListByProfileResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AFDCustomDomainsClientListByProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AFDCustomDomainsClientListByProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProfileHandleResponse(resp)
		},
	})
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *AFDCustomDomainsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *AFDCustomDomainsClientListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *AFDCustomDomainsClient) listByProfileHandleResponse(resp *http.Response) (AFDCustomDomainsClientListByProfileResponse, error) {
	result := AFDCustomDomainsClientListByProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDDomainListResult); err != nil {
		return AFDCustomDomainsClientListByProfileResponse{}, err
	}
	return result, nil
}

// BeginRefreshValidationToken - Updates the domain validation token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - customDomainName - Name of the domain under the profile which is unique globally.
//   - options - AFDCustomDomainsClientBeginRefreshValidationTokenOptions contains the optional parameters for the AFDCustomDomainsClient.BeginRefreshValidationToken
//     method.
func (client *AFDCustomDomainsClient) BeginRefreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginRefreshValidationTokenOptions) (*runtime.Poller[AFDCustomDomainsClientRefreshValidationTokenResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshValidationToken(ctx, resourceGroupName, profileName, customDomainName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDCustomDomainsClientRefreshValidationTokenResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDCustomDomainsClientRefreshValidationTokenResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RefreshValidationToken - Updates the domain validation token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *AFDCustomDomainsClient) refreshValidationToken(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginRefreshValidationTokenOptions) (*http.Response, error) {
	req, err := client.refreshValidationTokenCreateRequest(ctx, resourceGroupName, profileName, customDomainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshValidationTokenCreateRequest creates the RefreshValidationToken request.
func (client *AFDCustomDomainsClient) refreshValidationTokenCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, options *AFDCustomDomainsClientBeginRefreshValidationTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}/refreshValidationToken"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Updates an existing domain within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group.
//   - customDomainName - Name of the domain under the profile which is unique globally
//   - customDomainUpdateProperties - Domain properties
//   - options - AFDCustomDomainsClientBeginUpdateOptions contains the optional parameters for the AFDCustomDomainsClient.BeginUpdate
//     method.
func (client *AFDCustomDomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsClientBeginUpdateOptions) (*runtime.Poller[AFDCustomDomainsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, customDomainName, customDomainUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AFDCustomDomainsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDCustomDomainsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing domain within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *AFDCustomDomainsClient) update(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, customDomainName, customDomainUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AFDCustomDomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, customDomainName string, customDomainUpdateProperties AFDDomainUpdateParameters, options *AFDCustomDomainsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, customDomainUpdateProperties)
}
