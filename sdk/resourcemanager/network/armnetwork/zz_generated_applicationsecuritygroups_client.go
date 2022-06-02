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

// ApplicationSecurityGroupsClient contains the methods for the ApplicationSecurityGroups group.
// Don't use this type directly, use NewApplicationSecurityGroupsClient() instead.
type ApplicationSecurityGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationSecurityGroupsClient creates a new instance of ApplicationSecurityGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationSecurityGroupsClient, error) {
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
	client := &ApplicationSecurityGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// applicationSecurityGroupName - The name of the application security group.
// parameters - Parameters supplied to the create or update ApplicationSecurityGroup operation.
// options - ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationSecurityGroupsClient.BeginCreateOrUpdate
// method.
func (client *ApplicationSecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationSecurityGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ApplicationSecurityGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationSecurityGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *ApplicationSecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
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
func (client *ApplicationSecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// BeginDelete - Deletes the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// applicationSecurityGroupName - The name of the application security group.
// options - ApplicationSecurityGroupsClientBeginDeleteOptions contains the optional parameters for the ApplicationSecurityGroupsClient.BeginDelete
// method.
func (client *ApplicationSecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*runtime.Poller[ApplicationSecurityGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationSecurityGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ApplicationSecurityGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationSecurityGroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
func (client *ApplicationSecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
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
func (client *ApplicationSecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// Get - Gets information about the specified application security group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// applicationSecurityGroupName - The name of the application security group.
// options - ApplicationSecurityGroupsClientGetOptions contains the optional parameters for the ApplicationSecurityGroupsClient.Get
// method.
func (client *ApplicationSecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientGetOptions) (ApplicationSecurityGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationSecurityGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationSecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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
func (client *ApplicationSecurityGroupsClient) getHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientGetResponse, error) {
	result := ApplicationSecurityGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroup); err != nil {
		return ApplicationSecurityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the application security groups in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// options - ApplicationSecurityGroupsClientListOptions contains the optional parameters for the ApplicationSecurityGroupsClient.List
// method.
func (client *ApplicationSecurityGroupsClient) NewListPager(resourceGroupName string, options *ApplicationSecurityGroupsClientListOptions) *runtime.Pager[ApplicationSecurityGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationSecurityGroupsClientListResponse]{
		More: func(page ApplicationSecurityGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationSecurityGroupsClientListResponse) (ApplicationSecurityGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationSecurityGroupsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationSecurityGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationSecurityGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationSecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationSecurityGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups"
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
func (client *ApplicationSecurityGroupsClient) listHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientListResponse, error) {
	result := ApplicationSecurityGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroupListResult); err != nil {
		return ApplicationSecurityGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all application security groups in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// options - ApplicationSecurityGroupsClientListAllOptions contains the optional parameters for the ApplicationSecurityGroupsClient.ListAll
// method.
func (client *ApplicationSecurityGroupsClient) NewListAllPager(options *ApplicationSecurityGroupsClientListAllOptions) *runtime.Pager[ApplicationSecurityGroupsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationSecurityGroupsClientListAllResponse]{
		More: func(page ApplicationSecurityGroupsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationSecurityGroupsClientListAllResponse) (ApplicationSecurityGroupsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationSecurityGroupsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationSecurityGroupsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationSecurityGroupsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *ApplicationSecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *ApplicationSecurityGroupsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups"
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
func (client *ApplicationSecurityGroupsClient) listAllHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientListAllResponse, error) {
	result := ApplicationSecurityGroupsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroupListResult); err != nil {
		return ApplicationSecurityGroupsClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an application security group's tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// applicationSecurityGroupName - The name of the application security group.
// parameters - Parameters supplied to update application security group tags.
// options - ApplicationSecurityGroupsClientUpdateTagsOptions contains the optional parameters for the ApplicationSecurityGroupsClient.UpdateTags
// method.
func (client *ApplicationSecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsClientUpdateTagsOptions) (ApplicationSecurityGroupsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ApplicationSecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ApplicationSecurityGroupsClient) updateTagsHandleResponse(resp *http.Response) (ApplicationSecurityGroupsClientUpdateTagsResponse, error) {
	result := ApplicationSecurityGroupsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationSecurityGroup); err != nil {
		return ApplicationSecurityGroupsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
