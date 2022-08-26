//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

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

// FleetsClient contains the methods for the Fleets group.
// Don't use this type directly, use NewFleetsClient() instead.
type FleetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFleetsClient creates a new instance of FleetsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFleetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FleetsClient, error) {
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
	client := &FleetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// parameters - The Fleet to create or update.
// options - FleetsClientBeginCreateOrUpdateOptions contains the optional parameters for the FleetsClient.BeginCreateOrUpdate
// method.
func (client *FleetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, parameters Fleet, options *FleetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[FleetsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[FleetsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
func (client *FleetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, parameters Fleet, options *FleetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FleetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, parameters Fleet, options *FleetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// options - FleetsClientBeginDeleteOptions contains the optional parameters for the FleetsClient.BeginDelete method.
func (client *FleetsClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientBeginDeleteOptions) (*runtime.Poller[FleetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FleetsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FleetsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
func (client *FleetsClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, options)
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
func (client *FleetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// options - FleetsClientGetOptions contains the optional parameters for the FleetsClient.Get method.
func (client *FleetsClient) Get(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientGetOptions) (FleetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, options)
	if err != nil {
		return FleetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FleetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FleetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FleetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FleetsClient) getHandleResponse(resp *http.Response) (FleetsClientGetResponse, error) {
	result := FleetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Fleet); err != nil {
		return FleetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists fleets in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// options - FleetsClientListOptions contains the optional parameters for the FleetsClient.List method.
func (client *FleetsClient) NewListPager(options *FleetsClientListOptions) *runtime.Pager[FleetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetsClientListResponse]{
		More: func(page FleetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetsClientListResponse) (FleetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FleetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FleetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FleetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FleetsClient) listCreateRequest(ctx context.Context, options *FleetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/fleets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FleetsClient) listHandleResponse(resp *http.Response) (FleetsClientListResponse, error) {
	result := FleetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetListResult); err != nil {
		return FleetsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists fleets in the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - FleetsClientListByResourceGroupOptions contains the optional parameters for the FleetsClient.ListByResourceGroup
// method.
func (client *FleetsClient) NewListByResourceGroupPager(resourceGroupName string, options *FleetsClientListByResourceGroupOptions) *runtime.Pager[FleetsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetsClientListByResourceGroupResponse]{
		More: func(page FleetsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetsClientListByResourceGroupResponse) (FleetsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FleetsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FleetsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FleetsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FleetsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FleetsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FleetsClient) listByResourceGroupHandleResponse(resp *http.Response) (FleetsClientListByResourceGroupResponse, error) {
	result := FleetsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetListResult); err != nil {
		return FleetsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListCredentials - Lists the user credentials of a Fleet.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// options - FleetsClientListCredentialsOptions contains the optional parameters for the FleetsClient.ListCredentials method.
func (client *FleetsClient) ListCredentials(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientListCredentialsOptions) (FleetsClientListCredentialsResponse, error) {
	req, err := client.listCredentialsCreateRequest(ctx, resourceGroupName, fleetName, options)
	if err != nil {
		return FleetsClientListCredentialsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FleetsClientListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FleetsClientListCredentialsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCredentialsHandleResponse(resp)
}

// listCredentialsCreateRequest creates the ListCredentials request.
func (client *FleetsClient) listCredentialsCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientListCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/listCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCredentialsHandleResponse handles the ListCredentials response.
func (client *FleetsClient) listCredentialsHandleResponse(resp *http.Response) (FleetsClientListCredentialsResponse, error) {
	result := FleetsClientListCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetCredentialResults); err != nil {
		return FleetsClientListCredentialsResponse{}, err
	}
	return result, nil
}

// Update - Patches a fleet resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-02-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// fleetName - The name of the Fleet resource.
// options - FleetsClientUpdateOptions contains the optional parameters for the FleetsClient.Update method.
func (client *FleetsClient) Update(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientUpdateOptions) (FleetsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, fleetName, options)
	if err != nil {
		return FleetsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FleetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FleetsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FleetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *FleetsClient) updateHandleResponse(resp *http.Response) (FleetsClientUpdateResponse, error) {
	result := FleetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Fleet); err != nil {
		return FleetsClientUpdateResponse{}, err
	}
	return result, nil
}
