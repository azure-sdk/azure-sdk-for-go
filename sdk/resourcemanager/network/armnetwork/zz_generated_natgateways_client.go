//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// NatGatewaysClient contains the methods for the NatGateways group.
// Don't use this type directly, use NewNatGatewaysClient() instead.
type NatGatewaysClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNatGatewaysClient creates a new instance of NatGatewaysClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNatGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NatGatewaysClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &NatGatewaysClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// natGatewayName - The name of the nat gateway.
// parameters - Parameters supplied to the create or update nat gateway operation.
// options - NatGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the NatGatewaysClient.BeginCreateOrUpdate
// method.
func (client *NatGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*armruntime.Poller[NatGatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, natGatewayName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[NatGatewaysClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[NatGatewaysClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NatGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NatGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// natGatewayName - The name of the nat gateway.
// options - NatGatewaysClientBeginDeleteOptions contains the optional parameters for the NatGatewaysClient.BeginDelete method.
func (client *NatGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*armruntime.Poller[NatGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, natGatewayName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[NatGatewaysClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[NatGatewaysClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *NatGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, natGatewayName, options)
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
func (client *NatGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified nat gateway in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// natGatewayName - The name of the nat gateway.
// options - NatGatewaysClientGetOptions contains the optional parameters for the NatGatewaysClient.Get method.
func (client *NatGatewaysClient) Get(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientGetOptions) (NatGatewaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NatGatewaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NatGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NatGatewaysClient) getHandleResponse(resp *http.Response) (NatGatewaysClientGetResponse, error) {
	result := NatGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGateway); err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all nat gateways in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - NatGatewaysClientListOptions contains the optional parameters for the NatGatewaysClient.List method.
func (client *NatGatewaysClient) NewListPager(resourceGroupName string, options *NatGatewaysClientListOptions) *runtime.Pager[NatGatewaysClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[NatGatewaysClientListResponse]{
		More: func(page NatGatewaysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NatGatewaysClientListResponse) (NatGatewaysClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NatGatewaysClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NatGatewaysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NatGatewaysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *NatGatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *NatGatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways"
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
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NatGatewaysClient) listHandleResponse(resp *http.Response) (NatGatewaysClientListResponse, error) {
	result := NatGatewaysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGatewayListResult); err != nil {
		return NatGatewaysClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the Nat Gateways in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - NatGatewaysClientListAllOptions contains the optional parameters for the NatGatewaysClient.ListAll method.
func (client *NatGatewaysClient) NewListAllPager(options *NatGatewaysClientListAllOptions) *runtime.Pager[NatGatewaysClientListAllResponse] {
	return runtime.NewPager(runtime.PageProcessor[NatGatewaysClientListAllResponse]{
		More: func(page NatGatewaysClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NatGatewaysClientListAllResponse) (NatGatewaysClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NatGatewaysClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return NatGatewaysClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NatGatewaysClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *NatGatewaysClient) listAllCreateRequest(ctx context.Context, options *NatGatewaysClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *NatGatewaysClient) listAllHandleResponse(resp *http.Response) (NatGatewaysClientListAllResponse, error) {
	result := NatGatewaysClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGatewayListResult); err != nil {
		return NatGatewaysClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates nat gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// natGatewayName - The name of the nat gateway.
// parameters - Parameters supplied to update nat gateway tags.
// options - NatGatewaysClientUpdateTagsOptions contains the optional parameters for the NatGatewaysClient.UpdateTags method.
func (client *NatGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysClientUpdateTagsOptions) (NatGatewaysClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NatGatewaysClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NatGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *NatGatewaysClient) updateTagsHandleResponse(resp *http.Response) (NatGatewaysClientUpdateTagsResponse, error) {
	result := NatGatewaysClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGateway); err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	return result, nil
}
