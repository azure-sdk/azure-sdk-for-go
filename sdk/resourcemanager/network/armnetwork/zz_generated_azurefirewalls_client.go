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

// AzureFirewallsClient contains the methods for the AzureFirewalls group.
// Don't use this type directly, use NewAzureFirewallsClient() instead.
type AzureFirewallsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAzureFirewallsClient creates a new instance of AzureFirewallsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAzureFirewallsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureFirewallsClient, error) {
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
	client := &AzureFirewallsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the specified Azure Firewall.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// azureFirewallName - The name of the Azure Firewall.
// parameters - Parameters supplied to the create or update Azure Firewall operation.
// options - AzureFirewallsClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureFirewallsClient.BeginCreateOrUpdate
// method.
func (client *AzureFirewallsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureFirewallsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, azureFirewallName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AzureFirewallsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureFirewallsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Azure Firewall.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *AzureFirewallsClient) createOrUpdate(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
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
func (client *AzureFirewallsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters AzureFirewall, options *AzureFirewallsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Azure Firewall.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// azureFirewallName - The name of the Azure Firewall.
// options - AzureFirewallsClientBeginDeleteOptions contains the optional parameters for the AzureFirewallsClient.BeginDelete
// method.
func (client *AzureFirewallsClient) BeginDelete(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsClientBeginDeleteOptions) (*runtime.Poller[AzureFirewallsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, azureFirewallName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AzureFirewallsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureFirewallsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified Azure Firewall.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *AzureFirewallsClient) deleteOperation(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
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
func (client *AzureFirewallsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Azure Firewall.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// azureFirewallName - The name of the Azure Firewall.
// options - AzureFirewallsClientGetOptions contains the optional parameters for the AzureFirewallsClient.Get method.
func (client *AzureFirewallsClient) Get(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsClientGetOptions) (AzureFirewallsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureFirewallName, options)
	if err != nil {
		return AzureFirewallsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureFirewallsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureFirewallsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureFirewallsClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, options *AzureFirewallsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureFirewallsClient) getHandleResponse(resp *http.Response) (AzureFirewallsClientGetResponse, error) {
	result := AzureFirewallsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureFirewall); err != nil {
		return AzureFirewallsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Azure Firewalls in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// options - AzureFirewallsClientListOptions contains the optional parameters for the AzureFirewallsClient.List method.
func (client *AzureFirewallsClient) NewListPager(resourceGroupName string, options *AzureFirewallsClientListOptions) *runtime.Pager[AzureFirewallsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureFirewallsClientListResponse]{
		More: func(page AzureFirewallsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureFirewallsClientListResponse) (AzureFirewallsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureFirewallsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureFirewallsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureFirewallsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AzureFirewallsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AzureFirewallsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls"
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureFirewallsClient) listHandleResponse(resp *http.Response) (AzureFirewallsClientListResponse, error) {
	result := AzureFirewallsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureFirewallListResult); err != nil {
		return AzureFirewallsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the Azure Firewalls in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// options - AzureFirewallsClientListAllOptions contains the optional parameters for the AzureFirewallsClient.ListAll method.
func (client *AzureFirewallsClient) NewListAllPager(options *AzureFirewallsClientListAllOptions) *runtime.Pager[AzureFirewallsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureFirewallsClientListAllResponse]{
		More: func(page AzureFirewallsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureFirewallsClientListAllResponse) (AzureFirewallsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureFirewallsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureFirewallsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureFirewallsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *AzureFirewallsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls"
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *AzureFirewallsClient) listAllHandleResponse(resp *http.Response) (AzureFirewallsClientListAllResponse, error) {
	result := AzureFirewallsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureFirewallListResult); err != nil {
		return AzureFirewallsClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates tags of an Azure Firewall resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// azureFirewallName - The name of the Azure Firewall.
// parameters - Parameters supplied to update azure firewall tags.
// options - AzureFirewallsClientBeginUpdateTagsOptions contains the optional parameters for the AzureFirewallsClient.BeginUpdateTags
// method.
func (client *AzureFirewallsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsClientBeginUpdateTagsOptions) (*runtime.Poller[AzureFirewallsClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, azureFirewallName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AzureFirewallsClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AzureFirewallsClientUpdateTagsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateTags - Updates tags of an Azure Firewall resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *AzureFirewallsClient) updateTags(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, azureFirewallName, parameters, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *AzureFirewallsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters TagsObject, options *AzureFirewallsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureFirewallName == "" {
		return nil, errors.New("parameter azureFirewallName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureFirewallName}", url.PathEscape(azureFirewallName))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
