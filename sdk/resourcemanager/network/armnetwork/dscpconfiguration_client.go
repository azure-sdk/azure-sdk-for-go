//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// DscpConfigurationClient contains the methods for the DscpConfiguration group.
// Don't use this type directly, use NewDscpConfigurationClient() instead.
type DscpConfigurationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDscpConfigurationClient creates a new instance of DscpConfigurationClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscpConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscpConfigurationClient, error) {
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
	client := &DscpConfigurationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - parameters - Parameters supplied to the create or update dscp configuration operation.
//   - options - DscpConfigurationClientBeginCreateOrUpdateOptions contains the optional parameters for the DscpConfigurationClient.BeginCreateOrUpdate
//     method.
func (client *DscpConfigurationClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*runtime.Poller[DscpConfigurationClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, dscpConfigurationName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DscpConfigurationClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DscpConfigurationClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *DscpConfigurationClient) createOrUpdate(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dscpConfigurationName, parameters, options)
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
func (client *DscpConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration, options *DscpConfigurationClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - options - DscpConfigurationClientBeginDeleteOptions contains the optional parameters for the DscpConfigurationClient.BeginDelete
//     method.
func (client *DscpConfigurationClient) BeginDelete(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*runtime.Poller[DscpConfigurationClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dscpConfigurationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DscpConfigurationClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DscpConfigurationClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *DscpConfigurationClient) deleteOperation(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dscpConfigurationName, options)
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
func (client *DscpConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a DSCP Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - dscpConfigurationName - The name of the resource.
//   - options - DscpConfigurationClientGetOptions contains the optional parameters for the DscpConfigurationClient.Get method.
func (client *DscpConfigurationClient) Get(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientGetOptions) (DscpConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dscpConfigurationName, options)
	if err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscpConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscpConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, dscpConfigurationName string, options *DscpConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dscpConfigurationName == "" {
		return nil, errors.New("parameter dscpConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dscpConfigurationName}", url.PathEscape(dscpConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscpConfigurationClient) getHandleResponse(resp *http.Response) (DscpConfigurationClientGetResponse, error) {
	result := DscpConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfiguration); err != nil {
		return DscpConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a DSCP Configuration.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - DscpConfigurationClientListOptions contains the optional parameters for the DscpConfigurationClient.NewListPager
//     method.
func (client *DscpConfigurationClient) NewListPager(resourceGroupName string, options *DscpConfigurationClientListOptions) *runtime.Pager[DscpConfigurationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscpConfigurationClientListResponse]{
		More: func(page DscpConfigurationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscpConfigurationClientListResponse) (DscpConfigurationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscpConfigurationClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DscpConfigurationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscpConfigurationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DscpConfigurationClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *DscpConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DscpConfigurationClient) listHandleResponse(resp *http.Response) (DscpConfigurationClientListResponse, error) {
	result := DscpConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfigurationListResult); err != nil {
		return DscpConfigurationClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all dscp configurations in a subscription.
//
// Generated from API version 2022-09-01
//   - options - DscpConfigurationClientListAllOptions contains the optional parameters for the DscpConfigurationClient.NewListAllPager
//     method.
func (client *DscpConfigurationClient) NewListAllPager(options *DscpConfigurationClientListAllOptions) *runtime.Pager[DscpConfigurationClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscpConfigurationClientListAllResponse]{
		More: func(page DscpConfigurationClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscpConfigurationClientListAllResponse) (DscpConfigurationClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscpConfigurationClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DscpConfigurationClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscpConfigurationClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *DscpConfigurationClient) listAllCreateRequest(ctx context.Context, options *DscpConfigurationClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dscpConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *DscpConfigurationClient) listAllHandleResponse(resp *http.Response) (DscpConfigurationClientListAllResponse, error) {
	result := DscpConfigurationClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscpConfigurationListResult); err != nil {
		return DscpConfigurationClientListAllResponse{}, err
	}
	return result, nil
}
