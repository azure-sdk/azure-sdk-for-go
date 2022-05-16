//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AFDOriginGroupsClient contains the methods for the AFDOriginGroups group.
// Don't use this type directly, use NewAFDOriginGroupsClient() instead.
type AFDOriginGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAFDOriginGroupsClient creates a new instance of AFDOriginGroupsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAFDOriginGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AFDOriginGroupsClient, error) {
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
	client := &AFDOriginGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a new origin group within the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// originGroupName - Name of the origin group which is unique within the endpoint.
// originGroup - Origin group properties
// options - AFDOriginGroupsClientBeginCreateOptions contains the optional parameters for the AFDOriginGroupsClient.BeginCreate
// method.
func (client *AFDOriginGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup AFDOriginGroup, options *AFDOriginGroupsClientBeginCreateOptions) (*runtime.Poller[AFDOriginGroupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, originGroupName, originGroup, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AFDOriginGroupsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginGroupsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a new origin group within the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *AFDOriginGroupsClient) create(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup AFDOriginGroup, options *AFDOriginGroupsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originGroup, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *AFDOriginGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroup AFDOriginGroup, options *AFDOriginGroupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, originGroup)
}

// BeginDelete - Deletes an existing origin group within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// originGroupName - Name of the origin group which is unique within the profile.
// options - AFDOriginGroupsClientBeginDeleteOptions contains the optional parameters for the AFDOriginGroupsClient.BeginDelete
// method.
func (client *AFDOriginGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientBeginDeleteOptions) (*runtime.Poller[AFDOriginGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, originGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AFDOriginGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginGroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an existing origin group within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *AFDOriginGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, originGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AFDOriginGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing origin group within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// originGroupName - Name of the origin group which is unique within the endpoint.
// options - AFDOriginGroupsClientGetOptions contains the optional parameters for the AFDOriginGroupsClient.Get method.
func (client *AFDOriginGroupsClient) Get(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientGetOptions) (AFDOriginGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, originGroupName, options)
	if err != nil {
		return AFDOriginGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AFDOriginGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AFDOriginGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AFDOriginGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AFDOriginGroupsClient) getHandleResponse(resp *http.Response) (AFDOriginGroupsClientGetResponse, error) {
	result := AFDOriginGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDOriginGroup); err != nil {
		return AFDOriginGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProfilePager - Lists all of the existing origin groups within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// options - AFDOriginGroupsClientListByProfileOptions contains the optional parameters for the AFDOriginGroupsClient.ListByProfile
// method.
func (client *AFDOriginGroupsClient) NewListByProfilePager(resourceGroupName string, profileName string, options *AFDOriginGroupsClientListByProfileOptions) *runtime.Pager[AFDOriginGroupsClientListByProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[AFDOriginGroupsClientListByProfileResponse]{
		More: func(page AFDOriginGroupsClientListByProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AFDOriginGroupsClientListByProfileResponse) (AFDOriginGroupsClientListByProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AFDOriginGroupsClientListByProfileResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AFDOriginGroupsClientListByProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AFDOriginGroupsClientListByProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProfileHandleResponse(resp)
		},
	})
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *AFDOriginGroupsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *AFDOriginGroupsClientListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *AFDOriginGroupsClient) listByProfileHandleResponse(resp *http.Response) (AFDOriginGroupsClientListByProfileResponse, error) {
	result := AFDOriginGroupsClientListByProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AFDOriginGroupListResult); err != nil {
		return AFDOriginGroupsClientListByProfileResponse{}, err
	}
	return result, nil
}

// NewListResourceUsagePager - Checks the quota and actual usage of the given AzureFrontDoor origin group under the given
// CDN profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// originGroupName - Name of the origin group which is unique within the endpoint.
// options - AFDOriginGroupsClientListResourceUsageOptions contains the optional parameters for the AFDOriginGroupsClient.ListResourceUsage
// method.
func (client *AFDOriginGroupsClient) NewListResourceUsagePager(resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientListResourceUsageOptions) *runtime.Pager[AFDOriginGroupsClientListResourceUsageResponse] {
	return runtime.NewPager(runtime.PagingHandler[AFDOriginGroupsClientListResourceUsageResponse]{
		More: func(page AFDOriginGroupsClientListResourceUsageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AFDOriginGroupsClientListResourceUsageResponse) (AFDOriginGroupsClientListResourceUsageResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listResourceUsageCreateRequest(ctx, resourceGroupName, profileName, originGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AFDOriginGroupsClientListResourceUsageResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AFDOriginGroupsClientListResourceUsageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AFDOriginGroupsClientListResourceUsageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listResourceUsageHandleResponse(resp)
		},
	})
}

// listResourceUsageCreateRequest creates the ListResourceUsage request.
func (client *AFDOriginGroupsClient) listResourceUsageCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, options *AFDOriginGroupsClientListResourceUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listResourceUsageHandleResponse handles the ListResourceUsage response.
func (client *AFDOriginGroupsClient) listResourceUsageHandleResponse(resp *http.Response) (AFDOriginGroupsClientListResourceUsageResponse, error) {
	result := AFDOriginGroupsClientListResourceUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return AFDOriginGroupsClientListResourceUsageResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing origin group within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// originGroupName - Name of the origin group which is unique within the profile.
// originGroupUpdateProperties - Origin group properties
// options - AFDOriginGroupsClientBeginUpdateOptions contains the optional parameters for the AFDOriginGroupsClient.BeginUpdate
// method.
func (client *AFDOriginGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties AFDOriginGroupUpdateParameters, options *AFDOriginGroupsClientBeginUpdateOptions) (*runtime.Poller[AFDOriginGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, originGroupName, originGroupUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AFDOriginGroupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AFDOriginGroupsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an existing origin group within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *AFDOriginGroupsClient) update(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties AFDOriginGroupUpdateParameters, options *AFDOriginGroupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, originGroupName, originGroupUpdateProperties, options)
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

// updateCreateRequest creates the Update request.
func (client *AFDOriginGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, originGroupName string, originGroupUpdateProperties AFDOriginGroupUpdateParameters, options *AFDOriginGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, originGroupUpdateProperties)
}
